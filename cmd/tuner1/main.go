package main

import (
	"errors"
	"flag"
	"fmt"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/tuner"
	"os"
)

var Version = "localbuild"

func main() {
	show_version := flag.Bool("version", false, "display the program version")
	list_templates := flag.Bool("ls", false, "list templates")
	edit_standards := flag.Bool("edit-standards", false,
		"open the standards file with the default EDITOR.")
	tuning_template := flag.String("tuning", "",
		"a CSV list of notes for a tuning, or '+' followed by a template name")
	reference := flag.Float64("A4", 440.0,
		"the reference pitch to tune A4 to in Hertz")
	standards := flag.String("standards", "",
		"an alternate path to a standards.txt template file")
	wave_type := flag.String("wave", "sine",
		"the synth wave type to use (sine or square)")

	flag.Parse()

	if err := tuner.Execute(Version, show_version, list_templates, edit_standards,
		tuning_template, reference, standards, wave_type); err != nil {
		fmt.Fprintln(os.Stderr, err)

		var err_exit common.ExitError
		if errors.As(err, &err_exit) {
			os.Exit(err_exit.Code)
		}
	}
}
