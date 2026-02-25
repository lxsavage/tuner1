package statusbar

import "github.com/charmbracelet/lipgloss"

type segment struct {
	style    lipgloss.Style
	id       string
	text     string
	position lipgloss.Position // Supports Left, Center, and Right; other values will be treated as Left
}
type segmentOption func(m *segment)

// Create a segment based off of another segment
func SegmentWithBase(base segment, options ...segmentOption) segment {
	for _, opt := range options {
		opt(&base)
	}

	return base
}

// Create a new segment
func Segment(text string, options ...segmentOption) segment {
	s := segment{
		text:  text,
		style: StyleDefaultSegment,
	}

	return SegmentWithBase(s, options...)
}

// Specify the ID of a segment; used for dynamically-updating the segment's data
// in a status bar
func WithId(id string) segmentOption {
	return func(m *segment) {
		m.id = id
	}
}

// Specify the text value of a segment
func WithText(t string) segmentOption {
	return func(m *segment) {
		m.text = t
	}
}

// Specify the alignment of a segment; Left, Center, and Right are the only
// values supported, with other values being treated as equivalent to Left
func WithPosition(a lipgloss.Position) segmentOption {
	return func(m *segment) {
		m.position = a
	}
}

// Specify the style of a segment
func WithStyle(s lipgloss.Style) segmentOption {
	return func(m *segment) {
		m.style = s
	}
}

// Render the segment's current state
func (s segment) View() string {
	return s.style.Render(s.text)
}
