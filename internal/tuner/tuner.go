// Package tuner provides the core tuning functionality, including parsing tuning specifications,
// managing tuning standards/templates, and coordinating the tuning interface.
package tuner

import (
	"bufio"
	"errors"
	"fmt"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/editor"
	"lxsavage/tuner1/internal/synth"
	"lxsavage/tuner1/internal/tui"
	"lxsavage/tuner1/pkg/sysexit"
	"os"
	"path/filepath"
	"strings"
)

const a4 = 44100 // Hz

func getStandardsConfigFilepath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", common.ExitError{
			Code:    sysexit.EX_IOERR,
			Message: "Failed to get user config directory: " + err.Error(),
		}
	}

	return filepath.Join(configDir, "tuner1", "standards.txt"), nil
}

func getStandardsFileContents(stdFilePath string) ([]string, error) {

	stdFile, err := os.Open(stdFilePath)
	if err != nil {
		return nil, common.ExitError{
			Code:    sysexit.EX_IOERR,
			Message: "Failed to open standards file: " + err.Error(),
		}
	}
	defer stdFile.Close()

	var res []string
	sc := bufio.NewScanner(stdFile)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 {
			continue
		}

		res = append(res, line)
	}

	return res, nil
}

func listTemplates(stdFilePath string) error {
	standards, err := getStandardsFileContents(stdFilePath)
	if err != nil {
		return common.ExitError{
			Code:    sysexit.EX_IOERR,
			Message: "Failed to read standards file: " + err.Error(),
		}
	}

	fmt.Print(sprintStandards(standards))
	return nil
}

func Execute(d Config) error {
	if d.ShowVersion {
		fmt.Println(d.ProgramVersion)
		return nil
	}

	isTemplate := len(d.TuningTemplate) > 0 && (d.TuningTemplate)[0] == '+'
	shouldUseDefaultStdFile := d.EditStandards || d.ListTemplates || isTemplate

	stdFilePath := d.StandardsPath
	if shouldUseDefaultStdFile && len(stdFilePath) == 0 {
		var err error
		stdFilePath, err = getStandardsConfigFilepath()
		if err != nil {
			return err
		}
	}

	if d.ListTemplates {
		return listTemplates(stdFilePath)
	}

	if d.EditStandards {
		if err := editor.EditFile(stdFilePath); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return common.ExitError{
					Code: sysexit.EX_IOERR,
					Message: fmt.Sprintf("Unable to find standards file at \"%s\"; "+
						"perform a reinstall with the script to get the default one",
						stdFilePath),
				}
			}
			return common.ExitError{
				Code:    sysexit.EX_UNAVAILABLE,
				Message: "Unable to launch editor: " + err.Error(),
			}
		}
		return nil
	}

	if len(d.TuningTemplate) == 0 {
		return common.ExitError{
			Code: sysexit.EX_USAGE,
			Message: fmt.Sprintf("Please pass in a tuning specifier\n\nTry:\n"+
				"  %s -ls\n", os.Args[0]),
		}
	}

	tuningName := "user-defined tuning"
	rawTuning := d.TuningTemplate
	if (d.TuningTemplate)[0] == '+' {
		standards, err := getStandardsFileContents(stdFilePath)
		if err != nil {
			return common.ExitError{
				Code:    sysexit.EX_IOERR,
				Message: "Failed to open standards file: " + err.Error(),
			}
		}

		standardRawTuning, err := getStandard(standards, d.TuningTemplate)
		if err != nil {
			return common.ExitError{
				Code:    sysexit.EX_DATAERR,
				Message: "Failed to load template: " + err.Error(),
			}
		}

		tuningName = d.TuningTemplate
		rawTuning = standardRawTuning
	}

	stringNotes, err := getTuning(rawTuning)
	if err != nil {
		return common.ExitError{
			Code:    sysexit.EX_CONFIG,
			Message: "Failed to parse tuning: " + err.Error(),
		}
	}

	return tui.StartTUI(tui.Config{
		A4:          d.A4,
		Version:     d.ProgramVersion,
		StringNotes: stringNotes,
		TuningName:  tuningName,
		Synth:       synth.NewSynth(d.WaveType, a4),
		DebugMode:   d.DebugMode,
	})
}
