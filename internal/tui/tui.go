// Package tui provides the terminal user interface for the tuner1 application.
package tui

import (
	"fmt"
	"log"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/statusbar"
	"lxsavage/tuner1/internal/synth"
	"lxsavage/tuner1/pkg/note"
	"lxsavage/tuner1/pkg/sysexit"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

const (
	speakerSegmentId = "s-speaker"
	freqSegmentId    = "s-freq"
)

const (
	speakerPlaying = "🔊"
	speakerMuted   = "🔇"
)

var (
	segmentActiveSpeaker = statusbar.Segment(speakerPlaying,
		statusbar.WithId(speakerSegmentId),
		statusbar.WithStyle(StyleActiveSpeakerSegment),
	)
	segmentMutedSpeaker = statusbar.Segment(speakerMuted,
		statusbar.WithId(speakerSegmentId),
		statusbar.WithStyle(statusbar.StyleDefaultSegment),
	)
	segmentNoFrequency = statusbar.Segment("Note Frequency: 0.00 Hz",
		statusbar.WithId(freqSegmentId),
		statusbar.WithPosition(lipgloss.Left),
		statusbar.WithStyle(StyleSegmentNoBg),
	)
)

var (
	programVersion string
	waveSynth      synth.Synth
)

type model struct {
	width    int
	status   statusbar.Model
	choices  []note.Note
	help     help.Model
	keys     keyMap
	cursor   int
	selected int
	a4       float64
	debug    bool
}

func InitialUIModel(tuning []note.Note, a4 float64, tuningName string, debug bool) model {
	freqSegment := statusbar.Segment("",
		statusbar.WithStyle(statusbar.StyleDefaultStatusBar),
	)
	if debug {
		freqSegment = segmentNoFrequency
	}

	return model{
		status: statusbar.StatusBar(
			statusbar.WithSegments(
				segmentMutedSpeaker,
				freqSegment,
				statusbar.Segment("tuner1",
					statusbar.WithPosition(lipgloss.Center),
					statusbar.WithStyle(statusbar.StyleDefaultStatusBar),
				),
				statusbar.Segment(tuningName,
					statusbar.WithPosition(lipgloss.Right),
					statusbar.WithStyle(StyleSegmentNoBg),
				),
				statusbar.Segment(programVersion,
					statusbar.WithPosition(lipgloss.Right),
				),
			),
		),
		choices:  tuning,
		help:     help.New(),
		keys:     keys,
		cursor:   0,
		selected: -1, // -1 denotes muted
		a4:       a4,
		debug:    debug,
	}
}
func (m model) Init() tea.Cmd {
	return tea.WindowSize()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	selectionChanged := false
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.status.SetWidth(m.width)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Mute):
			if m.selected != -1 {
				m.selected = -1
				selectionChanged = true
			}
		case key.Matches(msg, m.keys.Next):
			m.selected = max((m.selected+1)%len(m.choices), 0)
			selectionChanged = true
		case key.Matches(msg, m.keys.Previous):
			if m.selected < 0 {
				m.selected = len(m.choices) - 1
			} else {
				m.selected = (m.selected - 1 + len(m.choices)) % len(m.choices)
			}

			selectionChanged = true
		case key.Matches(msg, m.keys.Left):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.keys.Right):
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.keys.Select):
			if m.selected != m.cursor {
				m.selected = m.cursor
				selectionChanged = true
			}
		case key.Matches(msg, m.keys.JumpToString):
			num, err := strconv.Atoi(msg.String())
			if err != nil {
				break
			}

			if num <= len(m.choices) {
				m.selected = num - 1
				selectionChanged = true
			}
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	if selectionChanged && m.selected < len(m.choices) {
		var newFreq float64 = 0
		if m.selected >= 0 {
			var err error
			newFreq, err = m.choices[m.selected].PitchOf(m.a4)

			if err != nil {
				log.Fatal(err)
			}
		}

		waveSynth.SetWaveFrequency(newFreq)
		m.updatePlayStatus()
	}

	return m, nil
}

func (m model) View() string {
	viewBox := fmt.Sprintf("%s\n\n%s\n\n%s",
		m.status.View(),
		renderTuningBox(m),
		m.help.View(m.keys),
	)

	return StyleCentered.
		Width(m.width).
		Render(viewBox)
}

func StartTUI(d Config) error {
	programVersion = d.Version
	waveSynth = d.Synth

	sr := beep.SampleRate(waveSynth.GetSampleRate())
	streamer := beep.StreamerFunc(waveSynth.SynthesizeWave)

	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(streamer)
	defer speaker.Close()

	tui := tea.NewProgram(InitialUIModel(d.StringNotes, d.A4, d.TuningName, d.DebugMode), tea.WithAltScreen())
	if _, err := tui.Run(); err != nil {
		return common.ExitError{
			Code:    sysexit.EX_SOFTWARE,
			Message: fmt.Sprintf("Critial error when running the tuner1 TUI:\n%s", err),
		}
	}

	return nil
}

func (m *model) updatePlayStatus() {
	if m.selected == -1 {
		m.status.SetSegmentById(speakerSegmentId, segmentMutedSpeaker)
		m.status.SetSegmentById(freqSegmentId, segmentNoFrequency)
		return
	}

	m.status.SetSegmentById(speakerSegmentId, segmentActiveSpeaker)

	freq, err := m.choices[m.selected].PitchOf(m.a4)
	if err != nil {
		return
	}

	m.status.AddSegmentOptionsById(freqSegmentId,
		statusbar.WithText(fmt.Sprintf("Note frequency: %.2f Hz", freq)),
	)
}
