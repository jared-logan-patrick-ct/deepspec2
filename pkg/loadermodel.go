package deepspec

import (
	"github.com/charmbracelet/bubbles/spinner"
)

// LoaderModel wraps the spinner for loading animation
// and a message to display next to it.
type LoaderModel struct {
	spinner spinner.Model
	text    string
}

func NewLoaderModel(text string) LoaderModel {
	return LoaderModel{
		spinner: spinner.New(),
		text:    text,
	}
}
