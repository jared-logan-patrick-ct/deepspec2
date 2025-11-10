package deepspec

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model is a local interface for models that support sizing
type Model interface {
	tea.Model
	SetSize(width, height int)
}

// TUI represents the terminal user interface
type TUI struct {
	models []Model
	model  Model
	width  int
	height int
}

// NewTUI creates a new TUI instance
func NewTUI() *TUI {
	// Initialize with a chat model
	models := []Model{NewChatModel()}

	return &TUI{
		models: models,
		model:  models[0],
	}
}

// Init initializes the TUI
func (t *TUI) Init() tea.Cmd {
	return t.model.Init()
}

// Update handles messages and updates the TUI state
func (t *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.width = msg.Width
		t.height = msg.Height

		// Update size for all models
		for i := range t.models {
			t.models[i].SetSize(msg.Width, msg.Height)
		}
		return t, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return t, tea.Quit
		}
	}

	updatedModel, cmd := t.model.Update(msg)
	t.model = updatedModel.(Model)
	return t, cmd
}

// View renders the TUI
func (t *TUI) View() string {
	// Render the active model
	return t.model.View()
}

// StartTUI starts the terminal user interface
func StartTUI() error {
	tui := NewTUI()
	p := tea.NewProgram(tui, tea.WithAltScreen(), tea.WithMouseAllMotion())
	_, err := p.Run()
	return err
}
