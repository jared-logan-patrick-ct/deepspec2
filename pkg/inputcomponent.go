package deepspec

import (
	"regexp"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var commandModeRegex = regexp.MustCompile(`^/[a-zA-Z\?]*$`)

// InputComponent represents the input box component
type InputComponent struct {
	input       textinput.Model
	width       int
	borderColor string
	commandMode bool
}

// NewInputComponent creates a new input component
func NewInputComponent() *InputComponent {
	ti := textinput.New()
	ti.Placeholder = "Enter message"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	return &InputComponent{
		input:       ti,
		borderColor: ctPurple,
	}
}

// Update handles messages for the input component
func (i *InputComponent) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	i.input, cmd = i.input.Update(msg)

	// Check if we should enter or exit command mode
	i.detectCommandMode()

	return cmd
}

// detectCommandMode checks if the input matches command pattern
func (i *InputComponent) detectCommandMode() {
	value := i.input.Value()
	wasCommandMode := i.commandMode
	i.commandMode = commandModeRegex.MatchString(value)

	if wasCommandMode != i.commandMode {
		i.updateInputMode()
	}
}

// updateInputMode updates the input component based on the current mode
func (i *InputComponent) updateInputMode() {
	if i.commandMode {
		i.input.Placeholder = "Enter command"
		i.borderColor = ctYellow
	} else {
		i.input.Placeholder = "Enter message"
		i.borderColor = ctPurple
	}
}

// SetWidth updates the input width
func (i *InputComponent) SetWidth(width int) {
	i.width = width
	i.input.Width = width - 2 // Minimal padding
}

// SetPlaceholder updates the input placeholder text
func (i *InputComponent) SetPlaceholder(text string) {
	i.input.Placeholder = text
}

// SetBorderColor updates the border color
func (i *InputComponent) SetBorderColor(color string) {
	i.borderColor = color
}

// Clear clears the input value
func (i *InputComponent) Clear() {
	i.input.SetValue("")
}

// Value returns the current input value
func (i *InputComponent) Value() string {
	return i.input.Value()
}

// IsCommandMode returns whether the input is in command mode
func (i *InputComponent) IsCommandMode() bool {
	return i.commandMode
}

// SetCommandMode allows external setting of command mode
func (i *InputComponent) SetCommandMode(mode bool) {
	i.commandMode = mode
	i.updateInputMode()
}

// View renders the input component
func (ic *InputComponent) View() string {
	borderStyle := lipgloss.NewStyle().
		MarginTop(1). // Add a top margin
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Color(ic.borderColor)).
		Width(ic.width)

	return borderStyle.Render(ic.input.View())
}
