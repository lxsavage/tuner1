package tui

import (
	"lxsavage/tuner1/pkg/note"

	"github.com/charmbracelet/bubbles/help"
)

type model struct {
	choices  []note.Note
	help     help.Model
	keys     keyMap
	cursor   int
	selected int
	a4       float64
	debug    bool
}

func InitialUIModel(tuning []note.Note, a4 float64, debug bool) model {
	return model{
		choices:  tuning,
		help:     help.New(),
		keys:     keys,
		cursor:   0,
		selected: -1, // -1 denotes muted
		a4:       a4,
		debug:    debug,
	}
}
