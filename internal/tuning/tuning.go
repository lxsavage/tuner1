package tuning

import (
	"errors"
	"fmt"
	"lxsavage/tuner1/pkg/note"
	"regexp"
	"strconv"
	"strings"
)

// A scientific-notation note is in the format <Note uppercase><accidental?><octave number>, i.e., C#4
var re_valid_note = regexp.MustCompile(`^([A-G][#b]?)(\d+)`)

func SprintStandards(standards []string) string {
	var result strings.Builder

	for _, line := range standards {
		key := strings.Split(line, ":")[0]
		result.WriteString("+" + key + "\n")
	}

	return result.String()
}

func GetStandard(standards []string, name string) (string, error) {
	check := name
	if name[0] == '+' {
		check = name[1:]
	}

	for _, line := range standards {
		kv := strings.Split(line, ":")

		if len(kv) != 2 {
			return "", errors.New("illegal standard definition \"" + line + "\"")
		}

		if kv[0] != check {
			continue
		}

		return kv[1], nil
	}

	return "", fmt.Errorf("standard +%s not found", check)
}

func GetTuning(tuning_csv string) ([]note.Note, error) {
	tunings_raw := strings.Split(tuning_csv, ",")
	var tunings []note.Note
	for _, note_name := range tunings_raw {
		matches := re_valid_note.FindStringSubmatch(note_name)
		if len(matches) != 3 {
			msg := "invalid note: " + note_name
			if len(note_name) > 0 && (note_name[0] < 65 /*A*/ || note_name[0] > 71 /*G*/) {
				msg += "\n- The note name must be represented by an uppercase A-G"
			}

			if len(note_name) > 1 && strings.Contains(note_name[1:], "B") {
				msg += "\n- A flat accidental must be represented by a lowercase \"b\""
			}
			return nil, errors.New(msg)
		}

		pitch := matches[1]
		octave, err := strconv.Atoi(matches[2])
		if err != nil {
			// This is added as a safety net, should never be hit since the regexp
			// ensures that the octave portion of the split is only digits
			return nil, err
		}

		tunings = append(tunings, note.Note{Pitch: pitch, Octave: octave})
	}

	return tunings, nil
}
