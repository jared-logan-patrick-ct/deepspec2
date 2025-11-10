package deepspec

import "github.com/charmbracelet/lipgloss"

// Color palette
const (
	ctYellow     = "#FFC806"
	ctGreen      = "#13C1C1"
	ctGreenLight = "#6ec2c2"
	ctPurple     = "#6259FE"
	errorRed     = "#FF8C8C" // Salmon/pink red
)

// Pre-defined styles
var (
	ErrorStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color(errorRed))
	UserMessageLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(ctYellow))
	UserMessageTextStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(ctGreenLight))
	AssistantBoldStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("252"))
	AssistantStyle        = lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("15"))
	LoaderStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color(ctYellow)).Bold(true)
)
