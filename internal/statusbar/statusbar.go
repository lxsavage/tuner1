package statusbar

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Model contains the state of a status bar view
type Model struct {
	width    int
	style    lipgloss.Style
	segments []segment
	cache    string
}
type statusBarOption func(m *Model)

// Create a new status bar with the provided options
func StatusBar(options ...statusBarOption) Model {
	bar := Model{
		style: StyleDefaultStatusBar,
	}

	for _, opt := range options {
		opt(&bar)
	}

	return bar
}

// Specify the width of a status bar
func WithWidth(w int) statusBarOption {
	return func(m *Model) {
		m.width = w
	}
}

// Specify the style of the background outside of segments in a status bar
func WithBackgroundStyle(s lipgloss.Style) statusBarOption {
	return func(m *Model) {
		m.style = s
	}
}

// Specify the segments to include in the status bar
func WithSegments(c ...segment) statusBarOption {
	return func(m *Model) {
		m.segments = append(m.segments, c...)
	}
}

// Set the width of the status bar
func (s *Model) SetWidth(w int) {
	s.width = max(0, w)
	s.cache = ""
}

// Gets the segment instance inside a status bar by its ID; returns the segment
// and a bool that is true if the segment was found
func (s Model) GetSegmentById(id string) (segment, bool) {
	idx := s.indexOf(id)
	if idx >= 0 {
		return s.segments[idx], true
	}
	return segment{}, false
}

// Replaces a segment instance inside a status bar with a new one. Returns a
// bool that is true if the segment was successfully replaced.
//
// AddSegmentOptionsById is the preferred way of implementing segments with
// dynamic data, but this has been included in the case of segment data that has
// a fixed set of possible visual states (e.g., editor modes on Vim) where there
// would be a performance/memory benefit to keeping an instance of each possible
// segment state and switching them out on the status bar.
func (s *Model) SetSegmentById(id string, c segment) bool {
	idx := s.indexOf(id)
	if idx < 0 {
		return false
	}

	s.cache = ""
	s.segments[idx] = c
	return true
}

// Adds additional options to a specific segment in a status bar using its ID.
//
// Use this when a segment needs to change with dynamic data; in the case of a
// fixed set of possible states, use SetSegmentById instead.
func (s *Model) AddSegmentOptionsById(id string, options ...segmentOption) bool {
	idx := s.indexOf(id)
	if idx < 0 {
		return false
	}

	s.cache = ""
	newSegment := SegmentWithBase(s.segments[idx], options...)
	s.segments[idx] = newSegment
	return true
}

// Render the status bar view's current state
func (s *Model) View() string {
	if len(s.cache) > 0 {
		return s.cache
	}

	outsideW := 0
	centerW := 0
	var lbar, cbar, rbar strings.Builder
	for _, component := range s.segments {
		c := component.View()
		width := lipgloss.Width(c)

		switch component.position {
		case lipgloss.Right:
			rbar.WriteString(c)
			outsideW += width
		case lipgloss.Center:
			centerW += width
			cbar.WriteString(c)
		case lipgloss.Left:
			fallthrough
		default:
			outsideW += width
			lbar.WriteString(c)
		}
	}

	lbarStr := lbar.String()
	cbarStr := cbar.String()
	rbarStr := rbar.String()

	leftBg := ""
	rightBg := ""

	innerPaddingNeeded := s.width - outsideW - centerW
	if innerPaddingNeeded > 0 {
		lbarWidth := lipgloss.Width(lbarStr)
		cbarWidth := lipgloss.Width(cbarStr)
		rbarWidth := lipgloss.Width(rbarStr)

		paddingNeededCenterSegmentsOnly := (s.width - cbarWidth) / 2
		leftCount := paddingNeededCenterSegmentsOnly - lbarWidth
		rightCount := paddingNeededCenterSegmentsOnly - rbarWidth

		// Compensate imperfect centering if the needed padding is odd width
		if (s.width-cbarWidth)%2 != 0 {
			leftCount++
		}

		// Guard rails to prevent panics when width is less than minimum width to render status bar
		if leftCount > 0 {
			leftBg = s.style.Render(strings.Repeat(" ", leftCount))
		}
		if rightCount > 0 {
			rightBg = s.style.Render(strings.Repeat(" ", rightCount))
		}
	}

	res := lbarStr + leftBg + cbarStr + rightBg + rbarStr
	s.cache = res
	return res
}

func (s Model) indexOf(id string) int {
	for i, c := range s.segments {
		if c.id == id {
			return i
		}
	}
	return -1
}
