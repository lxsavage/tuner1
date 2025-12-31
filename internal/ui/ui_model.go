package ui

import (
	"lxsavage/tuner1/internal/common"
)

type UIModel struct {
	choices  []common.Note
	cursor   int
	selected int
	a4       float64
}

func InitialUIModel(tuning []common.Note, a4 float64) UIModel {
	return UIModel{
		choices:  tuning,
		cursor:   0,
		selected: -1, // -1 denotes muted
		a4:       a4,
	}
}
