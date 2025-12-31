package tuning

import (
	"lxsavage/tuner1/internal/common"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestGetTuning(t *testing.T) {
	csv := "E2,A2,D3,G3,Bb3,E#4"
	want := []common.Note{
		{Pitch: "A", Octave: 2},
		{Pitch: "E", Octave: 2},
		{Pitch: "D", Octave: 3},
		{Pitch: "G", Octave: 3},
		{Pitch: "Bb", Octave: 3},
		{Pitch: "E#", Octave: 4},
	}

	of, err := GetTuning(csv)
	if err != nil {
		t.Fatalf("pitchOf(\"csv\") ERR: %s", err)
	}

	if !reflect.DeepEqual(of, want) {
		var want_str strings.Builder
		var of_str strings.Builder

		want_str.WriteString("[\n")
		of_str.WriteString("[\n")

		for n := range want {
			note := want[n]
			want_str.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(note.Octave) + "}\n")
		}

		for n := range of {
			note := of[n]
			of_str.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(note.Octave) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant %s", csv, of_str.String(), want_str.String())
	}
}
