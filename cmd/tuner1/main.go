package main

import (
	"errors"
	"flag"
	"fmt"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/tuner"
	"os"
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

	if err := tuner.Execute(Version, a4_pitch, show_version, list_templates, edit_standards,
		debug_mode, tuning_template, standards, wave_type); err != nil {
		fmt.Fprintln(os.Stderr, err)

		var err_exit common.ExitError
		if errors.As(err, &err_exit) {
			os.Exit(err_exit.Code)
		}
	}
}
