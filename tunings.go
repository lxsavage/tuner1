package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// A scientific-notation note is in the format <Note uppercase><accidental?><octave number>, i.e., C#4
var re_valid_note = regexp.MustCompile(`^([A-G][#b]?)(\d+)`)

func SprintStandards(std_file *os.File) string {
	var result strings.Builder
	sc := bufio.NewScanner(std_file)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 {
			continue
		}

		key := strings.Split(line, ":")[0]
		result.WriteString("+" + key + "\n")
	}

	return result.String()
}

func getStandard(std_file *os.File, name string) (string, error) {
	check := name
	if name[0] == '+' {
		check = name[1:]
	}

	sc := bufio.NewScanner(std_file)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 {
			continue
		}

		kv := strings.Split(line, ":")

		if len(kv) != 2 {
			return "", errors.New("illegal standard definition \"" + line + "\"")
		}

		if kv[0] != check {
			continue
		}

		return kv[1], nil
	}

	return "", fmt.Errorf("standard +%s not found", name)
}

func getTuning(tuning_csv string) ([]Note, error) {
	tunings_raw := strings.Split(tuning_csv, ",")
	var tunings []Note
	for i := range tunings_raw {
		note := tunings_raw[i]
		matches := re_valid_note.FindStringSubmatch(note)
		if len(matches) != 3 {
			msg := "invalid note: " + note
			if len(note) > 0 && (note[0] < 65 || note[0] > 90) {
				msg += "\n- The note name must be represented by an uppercase A-G"
			}

			if len(note) > 1 && strings.Contains(note[1:], "B") {
				msg += "\n- A flat accidental must be represented by a lowercase \"b\""
			}
			return nil, errors.New(msg)
		}

		fmt.Println("SDF", matches, note)

		pitch := matches[1]
		octave, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, errors.New("invalid note (failed to parse octave #): " + note)
		}

		tunings = append(tunings, Note{pitch, octave})
	}

	return tunings, nil
}
