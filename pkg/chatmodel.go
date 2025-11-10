package deepspec

import (
	"context"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"google.golang.org/genai"
)

// chatResponseMsg is a message containing the AI response
type chatResponseMsg struct {
	response string
	err      error
}

// healthCheckMsg is a message containing the health check result
type healthCheckMsg struct {
	healthy bool
}

// ChatModel represents the chat view with input and message history
type ChatModel struct {
	input        *InputComponent
	viewport     *ViewportComponent
	help         *HelpComponent
	width        int
	height       int
	vertexClient *VertexClient
	chatSession  *genai.Chat
	waiting      bool
	serverOnline bool
	mcpClient    *mcpClient
}

// NewChatModel creates a new chat model instance
func NewChatModel() *ChatModel {
	ctx := context.Background()
	vertexClient, err := NewVertexClient(ctx)

	var chatSession *genai.Chat
	if err == nil && vertexClient != nil {
		chatSession, err = vertexClient.StartChat()
	}

	mcpClient := NewMCPClient()
	if startErr := mcpClient.Start(ctx); startErr != nil {
		LogToFile("Failed to start MCP client: %v", startErr)
	}

	model := &ChatModel{
		input:        NewInputComponent(),
		viewport:     NewViewportComponent(),
		help:         NewHelpComponent(),
		vertexClient: vertexClient,
		chatSession:  chatSession,
		waiting:      false,
		mcpClient:    mcpClient,
	}
	if err != nil {
		// Clean up the error message for better UX
		errMsg := err.Error()
		// Check for common Vertex AI errors and simplify the message
		if strings.Contains(errMsg, "project/location or API key must be set when using Vertex AI backend") {
			errMsg = "Vertex AI configuration missing. Please set GOOGLE_CLOUD_PROJECT and GCP_LOCATION environment variables."
		}
		model.viewport.AddMessage(ErrorStyle.Render("Vertex AI Error: " + errMsg))
	}
	return model
}

// Init initializes the chat model
func (c *ChatModel) Init() tea.Cmd {
	return tea.Batch(
		c.performHealthCheck(),
	)
}

// performHealthCheck checks the MCP server health
func (c *ChatModel) performHealthCheck() tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		healthy, _ := c.mcpClient.HealthCheck(ctx)
		return healthCheckMsg{healthy: healthy}
	}
}

// scheduleHealthCheck schedules periodic health checks
func (c *ChatModel) scheduleHealthCheck() tea.Cmd {
	return tea.Tick(time.Second*3, func(time.Time) tea.Msg {
		ctx := context.Background()
		healthy, _ := c.mcpClient.HealthCheck(ctx)
		return healthCheckMsg{healthy: healthy}
	})
}

// Update handles messages for the chat model
func (c *ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case healthCheckMsg:
		c.serverOnline = msg.healthy
		c.help.SetConnectionStatus(c.serverOnline)
		return c, c.scheduleHealthCheck()

	case chatResponseMsg:
		c.waiting = false
		if msg.err != nil {
			c.viewport.ErrorLoader(msg.err.Error())
		} else {
			c.viewport.CompleteLoader(msg.response)
		}
		return c, nil

	case tea.KeyMsg:
		// Don't accept input while waiting for response
		if c.waiting {
			return c, nil
		}

		switch msg.Type {
		case tea.KeyEnter:
			value := c.input.Value()
			if value != "" {
				if c.input.IsCommandMode() {
					c.handleCommand(value)
					c.input.Clear()
					c.input.SetCommandMode(false)
					return c, nil
				} else {
					return c, c.handleChatMessage(value)
				}
			}
			return c, nil
		}

	default:
		// Allow viewport to handle spinner ticks & mouse
		if cmd = c.viewport.Update(msg); cmd != nil {
			return c, cmd
		}
	}

	// Update the input component
	cmd = c.input.Update(msg)

	return c, cmd
}

// handleChatMessage processes regular chat messages
func (c *ChatModel) handleChatMessage(message string) tea.Cmd {
	// Display user message via viewport helper
	c.viewport.AddUserMessage(message)
	c.input.Clear()
	c.input.SetCommandMode(false)

	if c.vertexClient == nil || c.chatSession == nil {
		c.viewport.ErrorLoader("Vertex AI client not initialized")
		return nil
	}

	c.waiting = true
	loaderCmd := c.viewport.StartLoader()

	// Async send
	messageCmd := func() tea.Msg {
		response, err := c.vertexClient.SendMessage(c.chatSession, message)
		if err != nil && strings.Contains(err.Error(), "session") {
			c.chatSession, _ = c.vertexClient.StartChat()
			response, err = c.vertexClient.SendMessage(c.chatSession, message)
		}
		return chatResponseMsg{response: response, err: err}
	}

	return tea.Batch(loaderCmd, messageCmd)
}

// handleCommand processes user commands
func (c *ChatModel) handleCommand(input string) {
	// Remove leading /
	input = strings.TrimPrefix(input, "/")
	input = strings.TrimSpace(input)

	switch input {
	case "?":
		c.showHelp()
	default:
		c.viewport.AddMessage("> /" + input)
		c.viewport.AddMessage("Unknown command. Type '/?' for help.")
	}
}

// showHelp displays available commands
func (c *ChatModel) showHelp() {
	helpText := []string{
		"Available Commands:",
		"  ? - Show this help message",
		"",
	}

	for _, line := range helpText {
		c.viewport.AddMessage(line)
	}
}

// SetSize updates the chat model dimensions
func (c *ChatModel) SetSize(width, height int) {
	c.width = width
	c.height = height

	// Calculate viewport height (total height minus input and help lines)
	viewportHeight := height - 4

	c.viewport.SetSize(width, viewportHeight)
	c.input.SetWidth(width)
	c.help.SetWidth(width)
}

// View renders the chat model
func (c *ChatModel) View() string {
	// Render viewport, input, and help
	viewportArea := c.viewport.View()
	inputArea := c.input.View()
	helpArea := c.help.View()

	// Combine viewport, input, and help
	return lipgloss.JoinVertical(
		lipgloss.Left,
		viewportArea,
		inputArea,
		helpArea,
	)
}
