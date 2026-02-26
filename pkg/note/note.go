// Package note provides a type and functions for working with musical notes.
package note

import (
	"errors"
	"fmt"
	"math"
)

// Pitch offsets from A; octave in standard notation starts a C and ends one octave higher
var pitchOffsets = map[string]int{
	"C":  -9,
	"C#": -8,
	"Db": -8,
	"D":  -7,
	"D#": -6,
	"Eb": -6,
	"E":  -5,
	"F":  -4,
	"F#": -3,
	"Gb": -3,
	"G":  -2,
	"G#": -1,
	"Ab": -1,
	"A":  0,
	"A#": 1,
	"Bb": 1,
	"B":  2,
}

type Note struct {
	Pitch  string
	Octave uint
}

func New(pitch string, octave uint) (Note, error) {
	if _, ok := pitchOffsets[pitch]; !ok {
		return Note{}, errors.New("invalid note name: " + pitch)
	}
	n := Note{pitch, octave}
	return n, nil
}

func (n Note) String() string {
	formatSpecifier := "%s%d"
	if len(n.Pitch) == 1 {
		formatSpecifier = "%s %d"
	}
	return fmt.Sprintf(formatSpecifier, n.Pitch, n.Octave)
}

// PitchOf determines the pitch in Hertz of a note in scientific notation with a given A4 reference.
//
// Generally, A4=440.0 Hz, but it may be different in some contexts.
func (n Note) PitchOf(a4 float64) (float64, error) {
	noteOffset, ok := pitchOffsets[n.Pitch]
	if !ok {
		return 0.0, errors.New("Invalid note name: " + n.Pitch)
	}

	octaveOffset := int(n.Octave) - 4
	semitoneOffset := noteOffset + (octaveOffset * 12)

	// A4 * 2^(n/12) where n is the number of semitones away from A4
	finalPitchHz := a4 * math.Pow(2, float64(semitoneOffset)/12)

	return finalPitchHz, nil
}
