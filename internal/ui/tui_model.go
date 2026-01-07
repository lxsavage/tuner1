package ui

import "lxsavage/tuner1/pkg/note"

type UIModel struct {
	choices  []note.Note
	cursor   int
	selected int
	a4       float64
}

func InitialUIModel(tuning []note.Note, a4 float64) UIModel {
	return UIModel{
		choices:  tuning,
		cursor:   0,
		selected: -1, // -1 denotes muted
		a4:       a4,
	}
}
