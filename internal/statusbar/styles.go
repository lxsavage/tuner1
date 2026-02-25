package statusbar

import "github.com/charmbracelet/lipgloss"

// Gruvbox colors
const (
	// Dark theme

	bg1 = "#3c3836"
	bg2 = "#504945"
	fg2 = "#d5c4a1"
	fg4 = "#a89984"

	// Light theme

	// TODO - implement
)

var StyleDefaultStatusBar = lipgloss.NewStyle().
	Background(lipgloss.Color(bg1)).
	Foreground(lipgloss.Color(fg4))

var StyleDefaultSegment = lipgloss.NewStyle().
	Padding(0, 1).
	Background(lipgloss.Color(bg2)).
	Foreground(lipgloss.Color(fg2))
