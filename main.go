package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var Version = "localbuild"

func getStandardsConfigFilepath() string {
	config_dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(config_dir, "tuner1", "standards.txt")
}

func editStandards(path_std_file string) error {
	var cmd string
	var args []string

	editor := os.Getenv("EDITOR")
	if len(editor) != 0 {
		cmd = editor
		args = []string{path_std_file}
	} else {
		switch runtime.GOOS {
		case "windows":
			cmd = "cmd"
			args = []string{"/c", "start", "", path_std_file}
		case "darwin":
			cmd = "open"
			args = []string{"-t", path_std_file}
		default:
			cmd = "xdg-open"
			args = []string{path_std_file}
		}
	}

	proc := exec.Command(cmd, args...)
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	return proc.Run()
}

func listTemplates(path_std_file string) {
	std_file, err := os.Open(path_std_file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(SprintStandards(std_file))
	std_file.Close()
}

func main() {
	version := flag.Bool("version", false, "display the program version")
	template := flag.Bool("ls", false, "list templates")
	edit_standards := flag.Bool("edit-standards", false, "open the standards file with the default EDITOR.")
	tuning := flag.String("tuning", "", "a CSV list of notes for a tuning, or '+' followed by a template name")
	reference := flag.Float64("A4", 440.0, "the reference pitch to tune A4 to in Hertz")
	standards := flag.String("standards", "", "an alternate path to a non-standard standards.txt template file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage for %s %s:\n", os.Args[0], Version)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version)
		os.Exit(0)
	}

	is_template := len(*tuning) > 0 && (*tuning)[0] == '+'
	should_use_default_std_file := *edit_standards || *template || is_template

	path_std_file := *standards
	if len(path_std_file) == 0 && should_use_default_std_file {
		path_std_file = getStandardsConfigFilepath()
	}

	if *template {
		listTemplates(path_std_file)
		os.Exit(0)
	}

	if *edit_standards {
		if err := editStandards(path_std_file); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to launch editor: %s\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if len(*tuning) == 0 {
		fmt.Fprintf(os.Stderr, "Please pass in a tuning specifier\n\nTry:\n  %s -ls\n", os.Args[0])
		os.Exit(1)
	}

	tuning_csv := *tuning
	if (*tuning)[0] == '+' {
		std_file, err := os.Open(path_std_file)
		if err != nil {
			log.Fatal(err)
		}
		defer std_file.Close()

		csv, err := getStandard(std_file, *tuning)
		if err != nil {
			log.Fatalf("failed to load template: %s\n", err)
		}

		tuning_csv = csv
	}

	tunings, err := getTuning(tuning_csv)
	if err != nil {
		log.Fatalf("Failed to parse tuning: %s\n", err)
	}

	startUI(tunings, *reference)
}
