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
	showVersion := flag.Bool("version", false, "display the program version")
	listTemplates := flag.Bool("ls", false, "list templates")
	editStandards := flag.Bool("edit-standards", false,
		"open the standards file with the default EDITOR.")
	tuningTemplate := flag.String("tuning", "",
		"a CSV list of notes for a tuning, or '+' followed by a template name")
	pitchA4 := flag.Float64("A4", 440.0,
		"the reference pitch to tune A4 to in Hertz")
	standards := flag.String("standards", "",
		"an alternate path to a standards.txt template file")
	waveType := flag.String("wave", "sine",
		"the synth wave type to use (sine, square, or sawtooth)")
	debugMode := flag.Bool("debug", false,
		"display additional debug information during runtime")

	flag.Parse()

	wave := "sine"
	possibleWaves := []string{"sine", "square", "sawtooth"}
	if slices.Contains(possibleWaves, *waveType) {
		wave = *waveType
	}

	cfg := tuner.Config{
		ProgramVersion: Version,
		WaveType:       wave,
		A4:             *pitchA4,
		TuningTemplate: *tuningTemplate,
		StandardsPath:  *standards,
		ShowVersion:    *showVersion,
		ListTemplates:  *listTemplates,
		EditStandards:  *editStandards,
		DebugMode:      *debugMode,
	}
	if err := tuner.Execute(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)

		if exitErr, ok := errors.AsType[common.ExitError](err); ok {
			os.Exit(exitErr.Code)
		}
	}

}
