package main

import (
	"errors"
	"flag"
	"fmt"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/tuner"
	"os"
	"slices"
)

// Version will be set during a CI build to the version string.
var Version = "localbuild"

func main() {
	show_version := flag.Bool("version", false, "display the program version")
	list_templates := flag.Bool("ls", false, "list templates")
	edit_standards := flag.Bool("edit-standards", false,
		"open the standards file with the default EDITOR.")
	tuning_template := flag.String("tuning", "",
		"a CSV list of notes for a tuning, or '+' followed by a template name")
	a4_pitch := flag.Float64("A4", 440.0,
		"the reference pitch to tune A4 to in Hertz")
	standards := flag.String("standards", "",
		"an alternate path to a standards.txt template file")
	wave_type := flag.String("wave", "sine",
		"the synth wave type to use (sine, square, or sawtooth)")
	debug_mode := flag.Bool("debug", false,
		"display additional debug information during runtime")

	flag.Parse()

	wave := "sine"
	possible_waves := []string{"sine", "square", "sawtooth"}
	if slices.Contains(possible_waves, *wave_type) {
		wave = *wave_type
	}

	err := tuner.Execute(tuner.Config{
		ProgramVersion: Version,
		WaveType:       wave,
		A4:             *a4_pitch,
		TuningTemplate: *tuning_template,
		StandardsPath:  *standards,
		ShowVersion:    *show_version,
		ListTemplates:  *list_templates,
		EditStandards:  *edit_standards,
		DebugMode:      *debug_mode,
	})

	if err == nil {
		return
	}

	fmt.Fprintln(os.Stderr, err)

	var err_exit common.ExitError
	if errors.As(err, &err_exit) {
		os.Exit(err_exit.Code)
	}
}
