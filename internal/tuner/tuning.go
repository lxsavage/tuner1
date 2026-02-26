package tuner

import (
	"errors"
	"fmt"
	"lxsavage/tuner1/pkg/note"
	"regexp"
	"strconv"
	"strings"
)

// A scientific-notation note is in the format <Note uppercase><accidental?><octave number>, i.e., C#4
var validNoteRegexp = regexp.MustCompile(`^([A-G][#b]?)(\d+)`)

func sprintStandards(standards []string) string {
	var result strings.Builder

	for _, line := range standards {
		key := strings.Split(line, ":")[0]
		result.WriteString("+" + key + "\n")
	}

	return result.String()
}

func getStandard(standards []string, name string) (string, error) {
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

func getTuning(rawTuning string) ([]note.Note, error) {
	rawTuningNotes := strings.Split(rawTuning, ",")
	var resultNotes []note.Note
	for _, noteName := range rawTuningNotes {
		matches := validNoteRegexp.FindStringSubmatch(noteName)
		if len(matches) != 3 {
			msg := "invalid note: " + noteName
			if len(noteName) > 0 && (noteName[0] < 65 /*A*/ || noteName[0] > 71 /*G*/) {
				msg += "\n- The note name must be represented by an uppercase A-G"
			}

			if len(noteName) > 1 && strings.Contains(noteName[1:], "B") {
				msg += "\n- A flat accidental must be represented by a lowercase \"b\""
			}
			return nil, errors.New(msg)
		}

		pitch := matches[1]
		octave, err := strconv.Atoi(matches[2])
		if err != nil {
			// Added as a safety net, should never be hit since the regexp ensures
			// that the octave portion of the split is only digits
			return nil, err
		}

		resultNote, err := note.New(pitch, uint(octave))
		if err != nil {
			// Added as a safety net, negative octaves are not possible due to the regexp
			return nil, err
		}

		resultNotes = append(resultNotes, resultNote)
	}

	return resultNotes, nil
}
