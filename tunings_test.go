package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestGetTuning(t *testing.T) {
	csv := "E2,A2,D3,G3,Bb3,E#4"
	want := []Note{
		{pitch: "E", octave: 2},
		{pitch: "A", octave: 2},
		{pitch: "D", octave: 3},
		{pitch: "G", octave: 3},
		{pitch: "Bb", octave: 3},
		{pitch: "E#", octave: 4},
	}

	of, err := getTuning(csv)
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
			want_str.WriteString("{pitch: " + note.pitch + ", octave: " + strconv.Itoa(note.octave) + "}\n")
		}

		for n := range of {
			note := of[n]
			of_str.WriteString("{pitch: " + note.pitch + ", octave: " + strconv.Itoa(note.octave) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant %s", csv, of_str.String(), want_str.String())
	}
}
