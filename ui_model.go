package main

type UIModel struct {
	choices  []Note
	cursor   int
	selected int
	a4       float64
}

func InitialUIModel(tuning []Note, a4 float64) UIModel {
	return UIModel{
		choices:  tuning,
		cursor:   0,
		selected: -1, // -1 denotes muted
		a4:       a4,
	}
}
