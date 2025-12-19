package main

type UIModel struct {
	choices  []Note
	cursor   int
	selected int
	a4       float64
}

func initialModel(tuning []Note, a4 float64) UIModel {
	return UIModel{
		choices:  tuning,
		cursor:   0,
		selected: -1,
		a4:       a4,
	}
}
