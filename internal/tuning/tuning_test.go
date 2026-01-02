package tuning

import (
	"lxsavage/tuner1/pkg/note"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSprintStandards(t *testing.T) {
	standards := []string{
		"e-standard:E2,A2,D3,G3,B3,E4",
		"eb-standard:Eb2,Ab2,Db3,Gb3,Bb3,Eb4",
		"d-standard:D2,G2,C3,F3,A3,D4",
		"c-standard:C2,F2,Bb2,Eb3,G3,C4",
	}
	want := "+e-standard\n+eb-standard\n+d-standard\n+c-standard\n"
	of := SprintStandards(standards)

	if of != want {
		t.Fatalf("SprintStandards(\"...\") = %s\nwant %s", of, want)
	}
}

func TestGetStandardExists(t *testing.T) {
	standards := []string{
		"e-standard:E2,A2,D3,G3,B3,E4",
		"eb-standard:Eb2,Ab2,Db3,Gb3,Bb3,Eb4",
		"d-standard:D2,G2,C3,F3,A3,D4",
		"c-standard:C2,F2,Bb2,Eb3,G3,C4",
	}
	name := "+eb-standard"

	want := "Eb2,Ab2,Db3,Gb3,Bb3,Eb4"
	of, err := GetStandard(standards, name)
	if err != nil {
		t.Fatalf("GetStandards(\"...\") ERR: %s", err)
	}

	if of != want {
		t.Fatalf("GetStandards(\"...\") = %s\nwant %s", of, want)
	}
}

func TestGetStandardNotExists(t *testing.T) {
	standards := []string{
		"e-standard:E2,A2,D3,G3,B3,E4",
		"eb-standard:Eb2,Ab2,Db3,Gb3,Bb3,Eb4",
		"d-standard:D2,G2,C3,F3,A3,D4",
		"c-standard:C2,F2,Bb2,Eb3,G3,C4",
	}
	name := "+eb-standard-7"
	want_err := "standard " + name + " not found"
	of, err := GetStandard(standards, name)
	if err == nil {
		t.Fatalf("GetStandards(\"...\") = %s\nwant contains(error(...),\"%s\")", of, want_err)
	}

	if !strings.Contains(err.Error(), want_err) {
		t.Fatalf("GetStandards(\"...\") = error(\"%s\")\nwant contains(error(...),\"%s\")", err.Error(), want_err)
	}
}

func TestGetStandardIllegalDefinition(t *testing.T) {
	standards := []string{
		"e-standard:E2,A2,D3,G3,B3,E4",
		"eb-standard!Eb2,Ab2,Db3,Gb3,Bb3,Eb4",
		"d-standard:D2,G2,C3,F3,A3,D4",
		"c-standard:C2,F2,Bb2,Eb3,G3,C4",
	}
	name := "+eb-standard-7"
	want_err := "illegal standard definition"
	of, err := GetStandard(standards, name)

	if err == nil {
		t.Fatalf("GetStandards(\"...\") = %s\nwant error()", of)
	}

	if !strings.Contains(err.Error(), want_err) {
		t.Fatalf("GetStandards(\"...\") = %s\nwant contains(error(...),\"%s\")", of, want_err)
	}
}

func TestGetTuningValid(t *testing.T) {
	csv := "A2,D3,G3,Bb3,E#4"
	want := []note.Note{
		{Pitch: "A", Octave: 2},
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

func TestGetTuningInvalidNoteName(t *testing.T) {
	csv := "H4"
	want_err := "\n- The note name must be represented by an uppercase A-G"

	of, err := GetTuning(csv)
	if err == nil {
		var of_str strings.Builder
		of_str.WriteString("[\n")
		for n := range of {
			note := of[n]
			of_str.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(note.Octave) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant contains(error(...),%s)", csv, of_str.String(), want_err)
	}

	if !strings.Contains(err.Error(), want_err) {
		t.Fatalf("getTuning(\"%s\") = error(%s)\nwant contains(error(...),%s)", csv, err.Error(), want_err)
	}
}

func TestGetTuningInvalidNoteAccidentalCapitalFlat(t *testing.T) {
	csv := "AB4"
	want_err := "\n- A flat accidental must be represented by a lowercase \"b\""

	of, err := GetTuning(csv)
	if err == nil {
		var of_str strings.Builder
		of_str.WriteString("[\n")
		for n := range of {
			note := of[n]
			of_str.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(note.Octave) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant contains(error(...),%s)", csv, of_str.String(), want_err)
	}

	if !strings.Contains(err.Error(), want_err) {
		t.Fatalf("getTuning(\"%s\") = error(%s)\nwant contains(error(...),%s)", csv, err.Error(), want_err)
	}
}

func TestGetTuningInvalidNoteCatchall(t *testing.T) {
	csv := "AbE"
	want_err := "invalid note"

	of, err := GetTuning(csv)
	if err == nil {
		var of_str strings.Builder
		of_str.WriteString("[\n")
		for n := range of {
			note := of[n]
			of_str.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(note.Octave) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant contains(error(...),%s)", csv, of_str.String(), want_err)
	}

	if !strings.Contains(err.Error(), want_err) {
		t.Fatalf("getTuning(\"%s\") = error(%s)\nwant contains(error(...),%s)", csv, err.Error(), want_err)
	}
}
