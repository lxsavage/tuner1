package tui

import (
	"fmt"
	"strings"
)

func renderTuningBox(m model) string {
	var index_line strings.Builder
	var note_line strings.Builder
	var interaction_line strings.Builder

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

		index_line.WriteString(top)
		note_line.WriteString(middle)
		interaction_line.WriteString(bottom)
	}

	return StyleTuningsBindingBox.Render(
		fmt.Sprintf("%s\n%s\n%s",
			index_line.String(),
			note_line.String(),
			interaction_line.String(),
		),
	)
}
