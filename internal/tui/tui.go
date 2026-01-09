// Package tui provides the terminal user interface for the tuner1 application.
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
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

var (
	p_version  string
	wave_synth synth.Synth
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	term_col_count, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		panic(err)
	}

	title_section := renderTitle(m)
	choice_section := renderChoices(m)
	keymap_section := renderKeymap(m)

	view_box := fmt.Sprintf("%s\n\n%s\n\n%s", title_section, choice_section, keymap_section)
	return ui_helpers.CenterBox(view_box, term_col_count)
}

func StartTUI(version string, debug bool, tunings []note.Note, a4 float64, synth_impl synth.Synth) error {
	p_version = version
	wave_synth = synth_impl

	sr := beep.SampleRate(wave_synth.GetSampleRate())
	streamer := beep.StreamerFunc(wave_synth.SynthesizeWave)

	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(streamer)
	defer speaker.Close()

	tui := tea.NewProgram(InitialUIModel(tunings, a4, debug), tea.WithAltScreen())
	if _, err := tui.Run(); err != nil {
		return common.ExitError{
			Code:    sysexit.EX_SOFTWARE,
			Message: fmt.Sprintf("Critial error when running the tuner1 TUI:\n%s", err),
		}
	}

	return nil
}
