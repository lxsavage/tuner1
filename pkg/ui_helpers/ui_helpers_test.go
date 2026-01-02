package ui_helpers

import (
	"strings"
	"testing"
)

func TestLeftPadLine(t *testing.T) {
	want := "  asdf"
	of := LeftPadLine("asdf", 6, ' ')

	if want != of {
		t.Errorf("LeftPadLine(\"asdf\",6) = \"%s\"; want \"%s\"", of, want)
	}
}

func TestLeftPadLineNoPad(t *testing.T) {
	want := "asdf"
	of := LeftPadLine("asdf", 4, ' ')

	if want != of {
		t.Errorf("LeftPadLine(\"asdf\",4) = \"%s\"; want \"%s\"", of, want)
	}
}

func TestLeftPadLineNegativePad(t *testing.T) {
	want := "asdf"
	of := LeftPadLine("asdf", 3, ' ')

	if want != of {
		t.Errorf("LeftPadLine(\"asdf\",3) = \"%s\"; want \"%s\"", of, want)
	}
}

func TestRightPadLine(t *testing.T) {
	want := "asdf  "
	of := RightPadLine("asdf", 6, ' ')

	if want != of {
		t.Errorf("LeftPadLine(\"asdf\",6) = \"%s\"; want \"%s\"", of, want)
	}
}

func TestRightPadLineNoPad(t *testing.T) {
	want := "asdf"
	of := RightPadLine("asdf", 4, ' ')

	if want != of {
		t.Errorf("RightPadLine(\"asdf\",4) = \"%s\"; want \"%s\"", of, want)
	}
}

func TestRightPadLineNegativePad(t *testing.T) {
	want := "asdf"
	of := RightPadLine("asdf", 3, ' ')

	if want != of {
		t.Errorf("RightPadLine(\"asdf\",3) = \"%s\"; want \"%s\"", of, want)
	}
}

func TestWrapBoxTwoLinesEqualLen(t *testing.T) {
	want := "┌────┐\n│asdf│\n│fdsa│\n└────┘"
	of := WrapBox("asdf\nfdsa", 0, 0)

	if want != of {
		t.Errorf("WrapBox(\"asdf\",0,0) =\n%s\n; want\n%s\n", of, want)
	}
}

func TestWrapBoxTwoLinesNotEqualLen(t *testing.T) {
	want := "┌─────┐\n│asdf │\n│afdsa│\n└─────┘"
	of := WrapBox("asdf\nafdsa", 0, 0)

	if want != of {
		t.Errorf("WrapBox(\"asdf\",0,0) =\n%s\n; want\n%s\n", of, want)
	}
}

func TestWrapBoxTwoLinesEqualLenXPad(t *testing.T) {
	want := "┌──────┐\n│ asdf │\n│ fdsa │\n└──────┘"
	of := WrapBox("asdf\nfdsa", 1, 0)

	if want != of {
		t.Errorf("WrapBox(\"asdf\",1,0) =\n%s\n; want\n%s\n", of, want)
	}
}

func TestWrapBoxTwoLinesEqualLenYPad(t *testing.T) {
	want := "┌────┐\n│    │\n│asdf│\n│fdsa│\n│    │\n└────┘"
	of := WrapBox("asdf\nfdsa", 0, 1)

	if want != of {
		t.Errorf("WrapBox(\"asdf\",0,1) =\n%s\n; want\n%s\n", of, want)
	}
}

func TestWrapBoxTwoLinesEqualLenXYPad(t *testing.T) {
	want := "┌──────┐\n│      │\n│ asdf │\n│ fdsa │\n│      │\n└──────┘"
	of := WrapBox("asdf\nfdsa", 1, 1)

	if want != of {
		t.Errorf("WrapBox(\"asdf\",1,1) =\n%s\n; want\n%s\n", of, want)
	}
}

func TestWrapBoxTwoLinesNotEqualLenXYPad(t *testing.T) {
	want := "┌──────┐\n│      │\n│ adf  │\n│ fdsa │\n│      │\n└──────┘"
	of := WrapBox("adf\nfdsa", 1, 1)

	if want != of {
		t.Errorf("WrapBox(\"asdf\",1,1) =\n%s\n; want\n%s\n", of, want)
	}
}

func TestCenterBoxEvenCoverage(t *testing.T) {
	want := "  asdf\n  fdsa"
	of := CenterBox("asdf\nfdsa", 8)

	if want != of {
		want_spaces := strings.ReplaceAll(want, " ", ".")
		of_spaces := strings.ReplaceAll(of, " ", ".")
		t.Errorf("CenterBox(\"asdf\\nfdsa\",8) =\n%s\n; want\n%s\n", of_spaces, want_spaces)
	}
}
func TestCenterBoxNoPad(t *testing.T) {
	want := "asdf\nfdsa"
	of := CenterBox("asdf\nfdsa", 4)

	if want != of {
		want_spaces := strings.ReplaceAll(want, " ", ".")
		of_spaces := strings.ReplaceAll(of, " ", ".")
		t.Errorf("CenterBox(\"asdf\\nfdsa\",4) =\n%s\n; want\n%s\n", of_spaces, want_spaces)
	}
}
func TestCenterBoxNegativePad(t *testing.T) {
	want := "asdf\nfdsa"
	of := CenterBox("asdf\nfdsa", 1)

	if want != of {
		want_spaces := strings.ReplaceAll(want, " ", ".")
		of_spaces := strings.ReplaceAll(of, " ", ".")
		t.Errorf("CenterBox(\"asdf\\nfdsa\",4) =\n%s\n; want\n%s\n", of_spaces, want_spaces)
	}
}
