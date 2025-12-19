package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

// TODO - See if possible (and better) to move these into non-global variables
var (
	// Mutex for adjusting sine wave frequency
	mu sync.RWMutex

	// Current sine wave frequency in Hz
	freq = 0.0

	// The position of the sine wave to encode in the current sample
	pos = 0
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
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.selected != m.cursor {
				m.selected = m.cursor
				selection_changed = true
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			num, err := strconv.Atoi(msg.String())
			if err != nil {
				break
			}

			if num > 0 && num <= len(m.choices) {
				m.selected = num - 1
				selection_changed = true
			}
		}
	}

	if selection_changed && m.selected < len(m.choices) {
		var newFreq float64 = 0
		if m.selected >= 0 {
			var err error
			newFreq, err = pitchOf(m.choices[m.selected], m.a4)

			if err != nil {
				log.Fatal(err)
			}
		}

		mu.Lock()
		freq = newFreq
		mu.Unlock()
	}

	return m, nil
}

func (m UIModel) View() string {
	var s strings.Builder

	s.WriteString("tuner1")
	if m.selected >= 0 {
		s.WriteString(" 📢")
	}
	s.WriteString("\n\n")

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected == i {
			checked = "X"
		}

		fmt.Fprintf(&s, "%d %s [%s] %s\n", i+1, cursor, checked, choice)
	}

	s.WriteString("\n[")
	for i := range m.choices {
		s.WriteString(strconv.Itoa(i + 1))
	}

	s.WriteString("] - select string by #\n")
	s.WriteString("[← ↑ ↓ →/hjkl] - move, [space/enter] - select, m - mute, q - quit\n")

	return s.String()
}

func startUI(tunings []Note, a4 float64) {
	//#region Audio
	// TODO: Clean up sine wave streaming/playing code here
	const sampleRate = 44100
	sr := beep.SampleRate(sampleRate)
	speaker.Init(sr, sr.N(time.Second/10))

	streamer := beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		mu.RLock()
		f := freq
		mu.RUnlock()

		// Generate a sine wave
		for i := range samples {
			v := math.Sin(2 * math.Pi * float64(pos) * f / sampleRate)
			samples[i][0] = v
			samples[i][1] = v
			pos++
		}

		return len(samples), true
	})

	speaker.Play(streamer)
	//#endregion

	ui := tea.NewProgram(initialModel(tunings, a4))
	if _, err := ui.Run(); err != nil {
		log.Fatalf("Critial error when running the tuner1 TUI:\n%s", err)
	}
}
