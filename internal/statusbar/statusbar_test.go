package statusbar

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestViewCentersAndWidths(t *testing.T) {
	bar := StatusBar(
		WithWidth(21),
		WithBackgroundStyle(lipgloss.NewStyle()),
		WithSegments(
			Segment("L",
				WithPosition(lipgloss.Left),
				WithStyle(lipgloss.NewStyle()),
			),
			Segment("C",
				WithPosition(lipgloss.Center),
				WithStyle(lipgloss.NewStyle()),
			),
			Segment("R",
				WithPosition(lipgloss.Right),
				WithStyle(lipgloss.NewStyle()),
			),
		),
	)

	v := bar.View()
	if got := lipgloss.Width(v); got != 21 {
		t.Fatalf("expected width 21, got %d; view=%q", got, v)
	}

	if !strings.Contains(v, "L") || !strings.Contains(v, "C") || !strings.Contains(v, "R") {
		t.Fatalf("view does not contain expected segments: %q", v)
	}

	// Verify center segment is roughly centered: distance from left edge to 'C'
	// and from 'C' to right edge should differ by at most 1 rune.
	leftIdx := strings.Index(v, "C")
	if leftIdx < 0 {
		t.Fatalf("center segment not found in view: %q", v)
	}
	leftSpace := leftIdx
	rightSpace := lipgloss.Width(v) - (leftIdx + len("C"))
	diff := leftSpace - rightSpace
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		t.Fatalf("centered asymmetry too large: left=%d right=%d view=%q", leftSpace, rightSpace, v)
	}
}

func TestCacheAndSetSegment(t *testing.T) {
	bar := StatusBar(
		WithWidth(10),
		WithBackgroundStyle(lipgloss.NewStyle()),
		WithSegments(
			Segment("old",
				WithId("seg"),
				WithStyle(lipgloss.NewStyle()),
			),
		),
	)

	p := &bar
	v1 := p.View()
	if !strings.Contains(v1, "old") {
		t.Fatalf("initial view missing text: %q", v1)
	}

	// Calling View again should hit the cache and return identical string
	v2 := p.View()
	if v1 != v2 {
		t.Fatalf("cached view mismatch; v1=%q v2=%q", v1, v2)
	}

	// Replace the segment; SetSegmentById should clear cache and reflect change
	ok := p.SetSegmentById("seg",
		Segment("new",
			WithId("seg"),
			WithStyle(lipgloss.NewStyle()),
		),
	)
	if !ok {
		t.Fatalf("SetSegmentById returned false")
	}
	v3 := p.View()
	if !strings.Contains(v3, "new") {
		t.Fatalf("updated view missing new text: %q", v3)
	}
	if strings.Contains(v3, "old") {
		t.Fatalf("updated view still contains old text: %q", v3)
	}

	// Ensure SetSegmentById returns false for unknown id
	if p.SetSegmentById("does-not-exist", Segment("x")) {
		t.Fatalf("SetSegmentById returned true for unknown id")
	}
}

func TestGetAndAddSegmentOptions(t *testing.T) {
	bar := StatusBar(
		WithWidth(15),
		WithBackgroundStyle(lipgloss.NewStyle()),
		WithSegments(
			Segment("base",
				WithId("id1"),
				WithStyle(lipgloss.NewStyle()),
			),
		),
	)

	// GetSegmentById should find the segment
	s, ok := bar.GetSegmentById("id1")
	if !ok {
		t.Fatalf("expected to find segment by id")
	}
	if s.text != "base" {
		t.Fatalf("unexpected segment text: %q", s.text)
	}

	// AddSegmentOptionsById should modify the segment (we change the text)
	p := &bar
	ok = p.AddSegmentOptionsById("id1",
		WithText("changed"),
	)
	if !ok {
		t.Fatalf("AddSegmentOptionsById returned false")
	}

	s2, ok := p.GetSegmentById("id1")
	if !ok {
		t.Fatalf("expected to find segment after modification")
	}
	if s2.text != "changed" {
		t.Fatalf("segment text not updated: %q", s2.text)
	}

	// Ensure AddSegmentOptionsById returns false for unknown id
	if p.AddSegmentOptionsById("nope", WithText("x")) {
		t.Fatalf("AddSegmentOptionsById returned true for unknown id")
	}
}

func TestSetWidthClearsCache(t *testing.T) {
	bar := StatusBar(
		WithWidth(8),
		WithBackgroundStyle(lipgloss.NewStyle()),
		WithSegments(
			Segment("A",
				WithStyle(lipgloss.NewStyle()),
			),
		),
	)

	p := &bar
	v1 := p.View()
	w1 := lipgloss.Width(v1)
	if w1 != 8 {
		t.Fatalf("initial width expected 8, got %d view=%q", w1, v1)
	}

	// Change width and ensure View reflects new width (cache must be cleared)
	p.SetWidth(12)
	v2 := p.View()
	w2 := lipgloss.Width(v2)
	if w2 != 12 {
		t.Fatalf("after SetWidth expected width 12, got %d view=%q", w2, v2)
	}
}
