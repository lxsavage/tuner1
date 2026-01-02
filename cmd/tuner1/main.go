package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"lxsavage/tuner1/internal/tuning"
	"lxsavage/tuner1/internal/ui"
	"lxsavage/tuner1/pkg/sysexit"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var Version = "localbuild"

func getStandardsConfigFilepath() string {
	config_dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(config_dir, "tuner1", "standards.txt")
}

func getStandardsFileContents(path_std_file string) ([]string, error) {
	var res []string

	std_file, err := os.Open(path_std_file)
	if err != nil {
		return nil, err
	}
	defer std_file.Close()

	sc := bufio.NewScanner(std_file)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 {
			continue
		}

		res = append(res, line)
	}

	return res, nil
}

func launchStandardsEditor(path_std_file string) error {
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
	standards, err := getStandardsFileContents(path_std_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read standards file: %s\n", err)
		os.Exit(sysexit.EX_IOERR)
	}

	fmt.Print(tuning.SprintStandards(standards))
}

func main() {
	version := flag.Bool("version", false, "display the program version")
	template := flag.Bool("ls", false, "list templates")
	edit_standards := flag.Bool("edit-standards", false, "open the standards file with the default EDITOR.")
	tuning_template := flag.String("tuning", "", "a CSV list of notes for a tuning, or '+' followed by a template name")
	reference := flag.Float64("A4", 440.0, "the reference pitch to tune A4 to in Hertz")
	standards := flag.String("standards", "", "an alternate path to a non-standard standards.txt template file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage for %s %s:\n", os.Args[0], Version)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version)
		os.Exit(sysexit.EX_OK)
	}

	is_template := len(*tuning_template) > 0 && (*tuning_template)[0] == '+'
	should_use_default_std_file := *edit_standards || *template || is_template

	path_std_file := *standards
	if should_use_default_std_file && len(path_std_file) == 0 {
		path_std_file = getStandardsConfigFilepath()
	}

	if *template {
		listTemplates(path_std_file)
		os.Exit(sysexit.EX_OK)
	}

	if *edit_standards {
		if err := launchStandardsEditor(path_std_file); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to launch editor: %s\n", err)
			os.Exit(sysexit.EX_UNAVAILABLE)
		}

		os.Exit(sysexit.EX_OK)
	}

	if len(*tuning_template) == 0 {
		fmt.Fprintf(os.Stderr, "Please pass in a tuning specifier\n\nTry:\n  %s -ls\n", os.Args[0])
		os.Exit(sysexit.EX_USAGE)
	}

	tuning_csv := *tuning_template
	if (*tuning_template)[0] == '+' {
		standards, err := getStandardsFileContents(path_std_file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open standards file: %s\n", err)
			os.Exit(sysexit.EX_IOERR)
		}

		csv, err := tuning.GetStandard(standards, *tuning_template)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load template: %s\n", err)
			os.Exit(sysexit.EX_DATAERR)
		}

		tuning_csv = csv
	}

	tunings, err := tuning.GetTuning(tuning_csv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse tuning: %s\n", err)
		os.Exit(sysexit.EX_CONFIG)
	}

	if ex_code := ui.StartUI(tunings, *reference, Version); ex_code != sysexit.EX_OK {
		os.Exit(ex_code)
	}
}
