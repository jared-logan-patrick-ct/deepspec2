package deepspec

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ViewportComponent represents the viewport for displaying messages
type ViewportComponent struct {
	viewport    viewport.Model
	content     []string
	ready       bool
	spinner     spinner.Model
	loaderIndex int
	showLoader  bool
}

// NewViewportComponent creates a new viewport component
func NewViewportComponent() *ViewportComponent {
	sp := spinner.New()
	sp.Style = LoaderStyle
	sp.Spinner = spinner.Line

	v := &ViewportComponent{
		content:     []string{},
		spinner:     sp,
		loaderIndex: -1,
		showLoader:  false,
	}

	// Add introductory banner and text
	v.addIntro()

	return v
}

// addIntro adds the welcome banner and intro text
func (v *ViewportComponent) addIntro() {
	// Split banner into lines for gradient coloring
	lines := []string{
		"     _",
		"  __| | ___  ___ _ __  ___ _ __   ___  ___",
		" / _  |/ _ \\/ _ \\ '_ \\/ __| '_ \\ / _ \\/ __|",
		"| (_| |  __/  __/ |_) \\__ \\ |_) |  __/ (__",
		" \\__,_|\\___|\\___| .__/|___/ .__/ \\___|\\___| " + Version,
		"                |_|       |_|               ",
	}

	// Apply gradient colors: yellow -> green -> purple
	colors := []string{ctYellow, ctYellow, ctGreen, ctGreen, ctPurple, ctPurple}
	var coloredLines []string
	for i, line := range lines {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[i]))
		coloredLines = append(coloredLines, style.Render(line))
	}

	banner := strings.Join(coloredLines, "\n")
	intro := "Welcome to deepspec. Enter a message to get started or execute commands using `/`.\n"

	v.content = append(v.content, banner)
	v.content = append(v.content, intro)
}

// SetSize updates the viewport dimensions
func (v *ViewportComponent) SetSize(width, height int) {
	marginTop := 1 // Account for the input component's top margin
	if !v.ready {
		v.viewport = viewport.New(width, height-marginTop)
		v.viewport.YPosition = 0
		v.viewport.MouseWheelEnabled = true
		v.ready = true
	} else {
		v.viewport.Width = width
		v.viewport.Height = height - marginTop
	}
	v.updateContent()
}

// AddMessage adds a new message to the viewport with text wrapping
func (v *ViewportComponent) AddMessage(message string) {
	v.content = append(v.content, message)
	v.updateContent()
}

// AddUserMessage convenience helper for chat user messages
func (v *ViewportComponent) AddUserMessage(text string) {
	v.AddMessage(UserMessageLabelStyle.Render("➤ ") + UserMessageTextStyle.Render(text))
}

// StartLoader inserts a loader line and begins animation
func (v *ViewportComponent) StartLoader() tea.Cmd {
	// Reset spinner
	v.spinner = spinner.New()
	v.spinner.Style = LoaderStyle
	v.spinner.Spinner = spinner.Points
	v.showLoader = true
	v.AddMessage(LoaderStyle.Render(v.spinner.View()))
	v.loaderIndex = v.Count() - 1
	return v.spinner.Tick
}

// UpdateLoaderTick handles spinner tick updates
func (v *ViewportComponent) UpdateLoaderTick(msg tea.Msg) tea.Cmd {
	if !v.showLoader {
		return nil
	}
	if tick, ok := msg.(spinner.TickMsg); ok {
		var cmd tea.Cmd
		v.spinner, cmd = v.spinner.Update(tick)
		if v.loaderIndex >= 0 {
			v.ReplaceMessage(v.loaderIndex, LoaderStyle.Render(v.spinner.View()))
		}
		return cmd
	}
	return nil
}

// CompleteLoader replaces loader with final assistant response
func (v *ViewportComponent) CompleteLoader(response string) {
	if v.loaderIndex >= 0 {
		v.ReplaceMessage(v.loaderIndex, "↳ "+AssistantStyle.Render(response)+"\n")
	} else {
		v.AddMessage(AssistantStyle.Render(response) + "\n")
	}
	v.loaderIndex = -1
	v.showLoader = false
}

// ErrorLoader replaces loader with an error line
func (v *ViewportComponent) ErrorLoader(errText string) {
	if v.loaderIndex >= 0 {
		v.ReplaceMessage(v.loaderIndex, ErrorStyle.Render("Error: "+errText))
	} else {
		v.AddMessage(ErrorStyle.Render("Error: " + errText))
	}
	v.loaderIndex = -1
	v.showLoader = false
}

// ReplaceMessage replaces the message at a given index
func (v *ViewportComponent) ReplaceMessage(index int, message string) {
	if index < 0 || index >= len(v.content) {
		return
	}
	v.content[index] = message
	v.updateContent()
}

// Count returns number of messages
func (v *ViewportComponent) Count() int { return len(v.content) }

// Clear clears all messages from the viewport
func (v *ViewportComponent) Clear() {
	v.content = []string{}
	v.updateContent()
}

// updateContent updates the viewport content
func (v *ViewportComponent) updateContent() {
	if v.ready {
		// Wrap all content to viewport width
		wrappedContent := make([]string, len(v.content))
		for i, msg := range v.content {
			wrappedContent[i] = lipgloss.NewStyle().Width(v.viewport.Width).Render(msg)
		}
		v.viewport.SetContent(strings.Join(wrappedContent, "\n"))
		v.viewport.GotoBottom()
	}
}

// Implement Init method for ViewportComponent
func (v *ViewportComponent) Init() tea.Cmd {
	return nil
}

// Update handles messages for the viewport and ensures re-rendering
func (v *ViewportComponent) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// Handle spinner updates
	if spinnerCmd := v.UpdateLoaderTick(msg); spinnerCmd != nil {
		cmds = append(cmds, spinnerCmd)
	}

	v.viewport, cmd = v.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

// View renders the viewport
func (v *ViewportComponent) View() string {
	if !v.ready {
		return ""
	}
	style := lipgloss.NewStyle().Height(v.viewport.Height)
	return style.Render(v.viewport.View())
}
