package tui

import (
	"fmt"
	"lxsavage/tuner1/pkg/ui_helpers"
	"strconv"
	"strings"
)

const (
	ANSI_RESET     = "\033[0m"
	ANSI_BG_BLUE   = "\033[44m"
	ANSI_TEXT_GRAY = "\033[1;30m"
)

func renderTitle(m model) string {
	title_text := "tuner1"

	if m.selected >= 0 {
		title_text += " 📢"

		if m.debug {
			freq, err := m.choices[m.selected].PitchOf(m.a4)
			if err != nil {
				panic(err)
			}

			title_text += fmt.Sprintf(" - Note frequency: %.2f Hz", freq)
		}
	}

	return title_text + "\n" + p_version
}

func renderChoices(m model) string {
	var index_line strings.Builder
	var note_line strings.Builder
	var interaction_line strings.Builder

	for i, choice := range m.choices {
		line_highlight := ""
		if m.cursor == i {
			line_highlight = ANSI_BG_BLUE
		}

		checked := " "
		if m.selected == i {
			checked = "•"
		}

		padded_choice := ui_helpers.LeftPadLine(choice.String(), 3, ' ')
		fmt.Fprintf(&index_line, "%s  %d  %s", line_highlight, i+1, ANSI_RESET)
		fmt.Fprintf(&note_line, "%s %s %s", line_highlight, padded_choice, ANSI_RESET)
		fmt.Fprintf(&interaction_line, "%s  %s  %s", line_highlight, checked, ANSI_RESET)
	}

	choice_box := fmt.Sprintf("%s\n%s\n%s", index_line.String(), note_line.String(), interaction_line.String())
	return ui_helpers.WrapBox(choice_box, 1, 0)
}

func renderKeymap(m model) string {
	var instructions_text strings.Builder
	instructions_text.WriteRune('[')
	for i := range m.choices {
		instructions_text.WriteString(strconv.Itoa(i + 1))
	}

	instructions_text.WriteString("] - select string by #, [↑ ↓/jk] - next/previous string\n")
	instructions_text.WriteString("[← →/hl] - move, [space/enter] - select, m - mute, q - quit\n")

	return instructions_text.String()
}
