package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var path_std_file = filepath.Join("config", "standards.txt")

func main() {
	template := flag.Bool("ls", false, "list templates")
	tuning := flag.String("tuning", "", "a CSV list of notes for a tuning, or '+' followed by a template name")
	reference := flag.Float64("A4", 440.0, "the reference pitch to tune to")
	flag.Parse()

	if *template {
		std_file, err := os.Open(path_std_file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(SprintStandards(std_file))
		std_file.Close()

		os.Exit(0)
	}

	if len(*tuning) == 0 {
		log.Fatal("Please pass in a tuning specifier")
	}

	var tuning_csv string
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
	} else {
		tuning_csv = *tuning
	}

	tunings, err := getTuning(tuning_csv)
	if err != nil {
		log.Fatalf("Failed to parse tuning: %s\n", err)
	}

	// This is an infinite ui render loop
	ui(tunings, *reference)
}
