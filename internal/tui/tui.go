package tui

import (
	"fmt"
	"log"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/synth"
	"lxsavage/tuner1/pkg/note"
	"lxsavage/tuner1/pkg/sysexit"
	"lxsavage/tuner1/pkg/ui_helpers"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

const (
	ANSI_RESET     = "\033[0m"
	ANSI_BG_BLUE   = "\033[44m"
	ANSI_TEXT_GRAY = "\033[1;30m"
)

var (
	p_version  string
	wave_synth synth.Synth
)

func (m UIModel) Init() tea.Cmd {
	return nil
}

func (m UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	selection_changed := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "m":
			if m.selected != -1 {
				m.selected = -1
				selection_changed = true
			}
		case "up", "k":
			m.selected = max((m.selected+1)%len(m.choices), 0)
			selection_changed = true
		case "down", "j":
			if m.selected < 0 {
				m.selected = len(m.choices) - 1
			} else {
				m.selected = (m.selected - 1 + len(m.choices)) % len(m.choices)
			}

			selection_changed = true
		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "l":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			num, err := strconv.Atoi(msg.String())
			if err != nil {
				break
			}

			if num <= len(m.choices) {
				m.selected = num - 1
				selection_changed = true
			}
		case "enter", " ":
			if m.selected != m.cursor {
				m.selected = m.cursor
				selection_changed = true
			}
		}
	}

	if selection_changed && m.selected < len(m.choices) {
		var new_freq float64 = 0
		if m.selected >= 0 {
			var err error
			new_freq, err = m.choices[m.selected].PitchOf(m.a4)

			if err != nil {
				log.Fatal(err)
			}
		}

		wave_synth.SetWaveFrequency(new_freq)
	}

	return m, nil
}

func renderTitle(m UIModel) string {
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

func renderChoices(m UIModel) string {
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

func renderKeymap(m UIModel) string {
	var instructions_text strings.Builder
	instructions_text.WriteRune('[')
	for i := range m.choices {
		instructions_text.WriteString(strconv.Itoa(i + 1))
	}

	instructions_text.WriteString("] - select string by #, [↑ ↓/jk] - next/previous string\n")
	instructions_text.WriteString("[← →/hl] - move, [space/enter] - select, m - mute, q - quit\n")

	return instructions_text.String()
}

func (m UIModel) View() string {
	term_col_count, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		panic(err)
	}

	// Render TUI sections
	title_section := renderTitle(m)
	choice_section := renderChoices(m)
	keymap_section := renderKeymap(m)

	// Create the view
	var view_box strings.Builder
	fmt.Fprintf(&view_box, "%s\n\n%s\n\n%s",
		title_section,
		choice_section,
		keymap_section)

	return ui_helpers.CenterBox(view_box.String(), term_col_count)
}

func StartTUI(version string, debug bool, tunings []note.Note, a4 float64, synth_impl synth.Synth) error {
	p_version = version
	wave_synth = synth_impl

	sr := beep.SampleRate(wave_synth.GetSampleRate())
	streamer := beep.StreamerFunc(wave_synth.SynthesizeWave)

	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(streamer)

	tui := tea.NewProgram(InitialUIModel(tunings, a4, debug), tea.WithAltScreen())
	if _, err := tui.Run(); err != nil {
		return common.ExitError{
			Code:    sysexit.EX_SOFTWARE,
			Message: fmt.Sprintf("Critial error when running the tuner1 TUI:\n%s", err),
		}
	}

	return nil
}
