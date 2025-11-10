package deepspec

import (
	"github.com/charmbracelet/lipgloss"
)

// HelpComponent represents the help text display component
type HelpComponent struct {
	width     int
	connected bool
}

// NewHelpComponent creates a new help component
func NewHelpComponent() *HelpComponent {
	return &HelpComponent{}
}

// SetWidth updates the help component width
func (h *HelpComponent) SetWidth(width int) {
	h.width = width
}

// SetConnectionStatus updates the server connection status
func (h *HelpComponent) SetConnectionStatus(isConnected bool) {
	h.connected = isConnected
}

// View renders the help component
func (h *HelpComponent) View() string {
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	// Define server status styles
	statusStyle := lipgloss.NewStyle().Bold(true)
	statusText := ""
	if h.connected {
		statusText = statusStyle.Foreground(lipgloss.Color(ctGreen)).Render("● online")
	} else {
		statusText = statusStyle.Foreground(lipgloss.Color(errorRed)).Render("● offline")
	}

	helpText := helpStyle.Render("Press Ctrl+C or Esc to quit • Enter to send message")

	// Calculate remaining width for server status
	helpTextWidth := lipgloss.Width(helpText)
	remainingWidth := h.width - helpTextWidth

	// Combine help text and server status with proper spacing
	return helpText + lipgloss.PlaceHorizontal(remainingWidth, lipgloss.Right, statusText)
}
