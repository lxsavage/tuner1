package tui

import (
	"fmt"
	"strings"
)

func renderTuningBox(m model) string {
	var stringNumberLine strings.Builder
	var noteNameLine strings.Builder
	var activeDotLine strings.Builder

	for i, choice := range m.choices {
		top := fmt.Sprintf("%d", i+1)
		middle := choice.String()

		bottom := ""
		if m.selected == i {
			bottom = "•"
		}

		if m.cursor == i {
			top = StyleStringHighlighted.Render(top)
			middle = StyleStringHighlighted.Render(middle)
			bottom = StyleStringHighlighted.Render(bottom)
		} else {
			top = StyleStringBlock.Render(top)
			middle = StyleStringBlock.Render(middle)
			bottom = StyleStringBlock.Render(bottom)
		}

		stringNumberLine.WriteString(top)
		noteNameLine.WriteString(middle)
		activeDotLine.WriteString(bottom)
	}

	return StyleTuningsBindingBox.Render(
		fmt.Sprintf("%s\n%s\n%s",
			stringNumberLine.String(),
			noteNameLine.String(),
			activeDotLine.String(),
		),
	)
}
