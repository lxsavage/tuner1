package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var Version = "localbuild"

func getHomeStandardsFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, ".config", "tuner1", "standards.txt")
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
	tuning := flag.String("tuning", "", "a CSV list of notes for a tuning, or '+' followed by a template name")
	reference := flag.Float64("A4", 440.0, "the reference pitch to tune A4 to in Hertz")
	standards := flag.String("standards", "", "a path to a standards.txt template file. By default, this uses one in ~/.config/tuner1/standards.txt")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage for %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Printf("%s\n", Version)
		os.Exit(0)
	}

	path_std_file := *standards
	if len(path_std_file) == 0 && (*template || ((len(*tuning) > 0) && (*tuning)[0] == '+')) {
		path_std_file = getHomeStandardsFilePath()
	}

	if *template {
		listTemplates(path_std_file)
		os.Exit(0)
	}

	if len(*tuning) == 0 {
		fmt.Fprintf(os.Stderr, "Please pass in a tuning specifier\n")
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
