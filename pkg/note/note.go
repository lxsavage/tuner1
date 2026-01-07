// Package note provides a type and functions for working with musical notes.
package note

import (
	"errors"
	"fmt"
	"math"
)

// Pitch offsets from A; octave in standard notation starts a C and ends one octave higher
var pitch_offsets = map[string]int{
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
	if _, valid_note := pitch_offsets[pitch]; !valid_note {
		return Note{}, errors.New("invalid note name: " + pitch)
	}
	n := Note{pitch, octave}
	return n, nil
}

func (n Note) String() string {
	format_specifier := "%s%d"
	if len(n.Pitch) == 1 {
		format_specifier = "%s %d"
	}
	return fmt.Sprintf(format_specifier, n.Pitch, n.Octave)
}

// Determine the pitch in Hertz of a note in scientific notation with a given A4 reference.
//
// Generally, A4=440.0 Hz, but it may be different in some contexts.
func (note Note) PitchOf(a4 float64) (float64, error) {
	var note_offset int
	var note_exists bool
	if note_offset, note_exists = pitch_offsets[note.Pitch]; !note_exists {
		return 0.0, errors.New("Invalid note name: " + note.Pitch)
	}

	octave_offset := int(note.Octave) - 4
	semitone_offset := note_offset + (octave_offset * 12)

	// A4 * 2^(n/12) where n is the number of semitones away from A4
	final_pitch_hz := a4 * math.Pow(2, float64(semitone_offset)/12)

	return final_pitch_hz, nil
}
