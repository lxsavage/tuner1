package tui

import (
	"lxsavage/tuner1/internal/synth"
	"lxsavage/tuner1/pkg/note"
)

type Config struct {
	A4        float64
	Version   string
	Tunings   []note.Note
	Synth     synth.Synth
	DebugMode bool
}
