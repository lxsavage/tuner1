package tuner

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
	of := sprintStandards(standards)

	if of != want {
		t.Fatalf("sprintStandards(\"...\") = %s\nwant %s", of, want)
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
	of, err := getStandard(standards, name)
	if err != nil {
		t.Fatalf("getStandards(\"...\") ERR: %s", err)
	}

	if of != want {
		t.Fatalf("getStandards(\"...\") = %s\nwant %s", of, want)
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
	wantErr := "standard " + name + " not found"
	of, err := getStandard(standards, name)
	if err == nil {
		t.Fatalf("getStandard(\"...\") = %s\nwant contains(error(...),\"%s\")", of, wantErr)
	}

	if !strings.Contains(err.Error(), wantErr) {
		t.Fatalf("getStandard(\"...\") = error(\"%s\")\nwant contains(error(...),\"%s\")", err.Error(), wantErr)
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
	wantErr := "illegal standard definition"
	of, err := getStandard(standards, name)

	if err == nil {
		t.Fatalf("getStandard(\"...\") = %s\nwant error()", of)
	}

	if !strings.Contains(err.Error(), wantErr) {
		t.Fatalf("getStandard(\"...\") = %s\nwant contains(error(...),\"%s\")", of, wantErr)
	}
}

func TestGetTuningValid(t *testing.T) {
	csv := "A2,D#3,G3,Bb3,E4"
	want := []note.Note{
		{Pitch: "A", Octave: 2},
		{Pitch: "D#", Octave: 3},
		{Pitch: "G", Octave: 3},
		{Pitch: "Bb", Octave: 3},
		{Pitch: "E", Octave: 4},
	}

	var wantStr strings.Builder
	wantStr.WriteString("[\n")
	for _, note := range want {
		wantStr.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(int(note.Octave)) + "}\n")
	}
	wantStr.WriteRune(']')

	of, err := getTuning(csv)
	if err != nil {
		t.Fatalf("getTuning(\"%s\") = error(\"%s\"), want %s", csv, err, wantStr.String())
	}

	if !reflect.DeepEqual(of, want) {
		var ofStr strings.Builder

		ofStr.WriteString("[\n")
		for _, note := range of {
			ofStr.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(int(note.Octave)) + "}\n")
		}
		ofStr.WriteRune(']')

		t.Fatalf("getTuning(\"%s\") = %s\nwant %s", csv, ofStr.String(), wantStr.String())
	}
}

func TestGetTuningInvalidNoteName(t *testing.T) {
	csv := "H4"
	wantErr := "\n- The note name must be represented by an uppercase A-G"

	of, err := getTuning(csv)
	if err == nil {
		var ofStr strings.Builder
		ofStr.WriteString("[\n")
		for _, note := range of {
			ofStr.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(int(note.Octave)) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant contains(error(...),%s)", csv, ofStr.String(), wantErr)
	}

	if !strings.Contains(err.Error(), wantErr) {
		t.Fatalf("getTuning(\"%s\") = error(%s)\nwant contains(error(...),%s)", csv, err.Error(), wantErr)
	}
}

func TestGetTuningInvalidNoteAccidentalCapitalFlat(t *testing.T) {
	csv := "AB4"
	wantErr := "\n- A flat accidental must be represented by a lowercase \"b\""

	of, err := getTuning(csv)
	if err == nil {
		var ofStr strings.Builder
		ofStr.WriteString("[\n")
		for _, note := range of {
			ofStr.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(int(note.Octave)) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant contains(error(...),%s)", csv, ofStr.String(), wantErr)
	}

	if !strings.Contains(err.Error(), wantErr) {
		t.Fatalf("getTuning(\"%s\") = error(%s)\nwant contains(error(...),%s)", csv, err.Error(), wantErr)
	}
}

func TestGetTuningInvalidNoteCatchall(t *testing.T) {
	csv := "AbE"
	wantErr := "invalid note"

	of, err := getTuning(csv)
	if err == nil {
		var ofStr strings.Builder
		ofStr.WriteString("[\n")
		for _, note := range of {
			ofStr.WriteString("{pitch: " + note.Pitch + ", octave: " + strconv.Itoa(int(note.Octave)) + "}\n")
		}

		t.Fatalf("getTuning(\"%s\") = %s\nwant contains(error(...),%s)", csv, ofStr.String(), wantErr)
	}

	if !strings.Contains(err.Error(), wantErr) {
		t.Fatalf("getTuning(\"%s\") = error(%s)\nwant contains(error(...),%s)", csv, err.Error(), wantErr)
	}
}
