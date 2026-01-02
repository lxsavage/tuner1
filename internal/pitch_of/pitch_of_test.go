package pitch_of

import (
	"lxsavage/tuner1/internal/common"
	"math"
	"strings"
	"testing"
)

func TestPitchOfC0(t *testing.T) {
	// C0 should be approximately 16.35 Hz when A4 is 440 Hz
	want := 16.35
	of, err := PitchOf(common.Note{Pitch: "C", Octave: 0}, 440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("pitchOf(note{pitch: \"C\", octave: 0}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}

func TestPitchOfFSharp2(t *testing.T) {
	// F#2 should be approximately 92.50 Hz when A4 is 440 Hz
	want := 92.50
	of, err := PitchOf(common.Note{Pitch: "F#", Octave: 2}, 440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("pitchOf(note{pitch: \"F#\", octave: 2}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}

func TestPitchOfC5(t *testing.T) {
	// C5 should be approximately 523.25 Hz when A4 is 440 Hz
	want := 523.25
	of, err := PitchOf(common.Note{Pitch: "C", Octave: 5}, 440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("pitchOf(note{pitch: \"C\", octave: 5}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}

func TestPitchOfB8(t *testing.T) {
	// B8 should be approximately 7902.13 Hz when A4 is 440 Hz
	want := 7902.13
	of, err := PitchOf(common.Note{Pitch: "B", Octave: 8}, 440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("pitchOf(note{pitch: \"B\", octave: 8}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}
func TestInvalidNoteName(t *testing.T) {
	want_err := "Invalid note name: H"
	of, err := PitchOf(common.Note{Pitch: "H", Octave: 8}, 440.0)
	if err == nil {
		t.Fatalf("pitchOf(note{pitch: \"H\", octave: 8}, 440.0) = %.2f; want contains(error(...), \"%s\")", of, want_err)
	}

	if !strings.Contains(err.Error(), want_err) {
		t.Fatalf("pitchOf(note{pitch: \"H\", octave: 8}, 440.0) = %.2f; want contains(error(...), \"%s\")", of, want_err)
	}
}
