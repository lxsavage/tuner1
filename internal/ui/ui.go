package ui

import (
	"fmt"
	"log"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/pitch_of"
	"lxsavage/tuner1/pkg/sysexit"
	"lxsavage/tuner1/pkg/ui_helpers"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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

const sample_rate = 44100 // Hz

var p_version string

var (
	freq     = 0.0 // Hz
	mu_freq  sync.RWMutex
	wave_pos int
)

// Generate a sine wave in place in samples. Returns the amount of samples
// after the operation, as well as if the operation was successful.
func generateSineWave(samples [][2]float64) (int, bool) {
	mu_freq.RLock()
	f := freq
	mu_freq.RUnlock()

	for i := range samples {
		v := math.Sin(2 * math.Pi * float64(wave_pos) * f / sample_rate)
		samples[i][0] = v
		samples[i][1] = v
		wave_pos++
	}

	return len(samples), true
}

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
		var newFreq float64 = 0
		if m.selected >= 0 {
			var err error
			newFreq, err = pitch_of.PitchOf(m.choices[m.selected], m.a4)

			if err != nil {
				log.Fatal(err)
			}
		}

		mu_freq.Lock()
		freq = newFreq
		mu_freq.Unlock()
	}

	return m, nil
}

func (m UIModel) View() string {
	term_col_count, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		panic(err)
	}

	var view_text strings.Builder

	// Title box
	var title_text strings.Builder

	title_text.WriteString("tuner1")
	if m.selected >= 0 {
		title_text.WriteString(" 📢")
	}
	fmt.Fprintf(&title_text, "\n%s", p_version)

	title_box := title_text.String()

	// Choice box
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
	choice_box = ui_helpers.WrapBox(choice_box, 1, 0)

	// Instruction box
	var instructions_text strings.Builder
	instructions_text.WriteRune('[')
	for i := range m.choices {
		instructions_text.WriteString(strconv.Itoa(i + 1))
	}

	instructions_text.WriteString("] - select string by #, [↑ ↓/jk] - next/previous string\n")
	instructions_text.WriteString("[← →/hl] - move, [space/enter] - select, m - mute, q - quit\n")

	instruction_box := instructions_text.String()

	// Create the view
	fmt.Fprintf(&view_text, "%s\n\n%s\n\n%s",
		title_box,
		choice_box,
		instruction_box)

	return ui_helpers.CenterBox(view_text.String(), term_col_count)
}

func StartUI(tunings []common.Note, a4 float64, version string) int {
	p_version = version

	sr := beep.SampleRate(sample_rate)
	speaker.Init(sr, sr.N(time.Second/10))

	streamer := beep.StreamerFunc(generateSineWave)
	speaker.Play(streamer)

	ui := tea.NewProgram(InitialUIModel(tunings, a4), tea.WithAltScreen())
	if _, err := ui.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Critial error when running the tuner1 TUI:\n%s", err)
		return sysexit.EX_SOFTWARE
	}

	return sysexit.EX_OK
}
