package tui

import (
	"lxsavage/tuner1/internal/statusbar"

	"github.com/charmbracelet/lipgloss"
)

const (
	blue      = lipgloss.Color("#458588")
	darkGray  = lipgloss.Color("#3c3836")
	lightGray = lipgloss.Color("#ebdbb2")
)

var StyleCentered = lipgloss.NewStyle().Align(lipgloss.Center)

var StyleActiveSpeakerSegment = statusbar.StyleDefaultSegment.
	Background(blue)

var StyleTuningsBindingBox = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	Padding(0, 1)

var StyleStringBlock = StyleCentered.
	Width(5).
	Padding(0, 1)

var StyleStringHighlighted = StyleStringBlock.
	Background(lipgloss.Color(blue)).
	Foreground(lipgloss.Color(lightGray))
