package note

import (
	"math"
	"strings"
	"testing"
)

func TestNewNoteValid(t *testing.T) {
	want := Note{"A", 4}
	of, err := New("A", 4)
	if err != nil {
		t.Fatalf("New(\"A\", 4) = error(\"%s\"); want %s", err.Error(), want)
	}

	if want != of {
		t.Fatalf("New(\"A\", 4) = %s; want %s", of, want)
	}
}

func TestNewNoteInvalidNote(t *testing.T) {
	want_str := "invalid note name"
	of, err := New("H", 4)
	if err == nil {
		t.Fatalf("New(\"H\", 4) = %s; want contains(error(...), \"%s\")", of, want_str)
	}

	if !strings.Contains(err.Error(), want_str) {
		t.Fatalf("New(\"H\", 4) = error(\"%s\"); want contains(error(...), \"%s\")", err.Error(), want_str)
	}
}

func TestPitchOfC0(t *testing.T) {
	// C0 should be approximately 16.35 Hz when A4 is 440 Hz
	want := 16.35
	of, err := Note{Pitch: "C", Octave: 0}.PitchOf(440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("PitchOf(Note{pitch: \"C\", octave: 0}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}

func TestPitchOfFSharp2(t *testing.T) {
	// F#2 should be approximately 92.50 Hz when A4 is 440 Hz
	want := 92.50
	of, err := Note{Pitch: "F#", Octave: 2}.PitchOf(440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("PitchOf(Note{pitch: \"F#\", octave: 2}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}

func TestPitchOfC5(t *testing.T) {
	// C5 should be approximately 523.25 Hz when A4 is 440 Hz
	want := 523.25
	of, err := Note{Pitch: "C", Octave: 5}.PitchOf(440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("PitchOf(Note{pitch: \"C\", octave: 5}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}

func TestPitchOfB8(t *testing.T) {
	// B8 should be approximately 7902.13 Hz when A4 is 440 Hz
	want := 7902.13
	of, err := Note{Pitch: "B", Octave: 8}.PitchOf(440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("PitchOf(Note{pitch: \"B\", octave: 8}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}
func TestInvalidNoteName(t *testing.T) {
	want_err := "Invalid note name: H"
	of, err := Note{Pitch: "H", Octave: 8}.PitchOf(440.0)
	if err == nil {
		t.Fatalf("PitchOf(Note{pitch: \"H\", octave: 8}, 440.0) = %.2f; want contains(error(...), \"%s\")", of, want_err)
	}

	if !strings.Contains(err.Error(), want_err) {
		t.Fatalf("PitchOf(Note{pitch: \"H\", octave: 8}, 440.0) = %.2f; want contains(error(...), \"%s\")", of, want_err)
	}
}
