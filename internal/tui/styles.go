package tui

import (
	"lxsavage/tuner1/internal/statusbar"

	"github.com/charmbracelet/lipgloss"
)

const (
	green     = lipgloss.Color("#8ec07c")
	blue      = lipgloss.Color("#558387")
	darkGray  = lipgloss.Color("#3c3836")
	lightGray = lipgloss.Color("#e8dcb7")
)

var StyleActiveSpeakerSegment = statusbar.StyleDefaultSegment.
	Background(blue)

var StyleTuningHighlighted = lipgloss.NewStyle().
	Background(blue).
	Foreground(lightGray)
