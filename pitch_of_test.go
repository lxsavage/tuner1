package main

import (
	"math"
	"testing"
)

func TestPitchOfC0(t *testing.T) {
	// C0 should be approximately 16.35 Hz when A4 is 440 Hz
	want := 16.35
	of, err := pitchOf(Note{pitch: "C", octave: 0}, 440.0)
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
	of, err := pitchOf(Note{pitch: "F#", octave: 2}, 440.0)
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
	of, err := pitchOf(Note{pitch: "C", octave: 5}, 440.0)
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
	of, err := pitchOf(Note{pitch: "B", octave: 8}, 440.0)
	if err != nil {
		t.Fatal(err)
	}

	if diff := math.Abs(of - want); diff > 0.01 {
		t.Errorf("pitchOf(note{pitch: \"B\", octave: 8}, 440.0) = %.2f; want approximately %.2f", of, want)
	}
}
