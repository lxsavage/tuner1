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

const sample_rate = 44100 // Hz

func getStandardsConfigFilepath() (string, error) {
	config_dir, err := os.UserConfigDir()
	if err != nil {
		return "", common.ExitError{
			Code:    sysexit.EX_IOERR,
			Message: "Failed to get user config directory: " + err.Error(),
		}
	}

	return filepath.Join(config_dir, "tuner1", "standards.txt"), nil
}

func getStandardsFileContents(path_std_file string) ([]string, error) {
	var res []string

	std_file, err := os.Open(path_std_file)
	if err != nil {
		return nil, common.ExitError{
			Code:    sysexit.EX_IOERR,
			Message: "Failed to open standards file: " + err.Error(),
		}
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

func listTemplates(path_std_file string) error {
	standards, err := getStandardsFileContents(path_std_file)
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

	is_template := len(d.TuningTemplate) > 0 && (d.TuningTemplate)[0] == '+'
	should_use_default_std_file := d.EditStandards || d.ListTemplates || is_template

	path_std_file := d.StandardsPath
	if should_use_default_std_file && len(path_std_file) == 0 {
		var err error
		path_std_file, err = getStandardsConfigFilepath()
		if err != nil {
			return err
		}
	}

	if d.ListTemplates {
		return listTemplates(path_std_file)
	}

	if d.EditStandards {
		if err := editor.EditFile(path_std_file); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return common.ExitError{
					Code: sysexit.EX_IOERR,
					Message: fmt.Sprintf("Unable to find standards file at \"%s\"; "+
						"perform a reinstall with the script to get the default one",
						path_std_file),
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

	tuning_csv := d.TuningTemplate
	if (d.TuningTemplate)[0] == '+' {
		standards, err := getStandardsFileContents(path_std_file)
		if err != nil {
			return common.ExitError{
				Code:    sysexit.EX_IOERR,
				Message: "Failed to open standards file: " + err.Error(),
			}
		}

		csv, err := getStandard(standards, d.TuningTemplate)
		if err != nil {
			return common.ExitError{
				Code:    sysexit.EX_DATAERR,
				Message: "Failed to load template: " + err.Error(),
			}
		}

		tuning_csv = csv
	}

	tunings, err := getTuning(tuning_csv)
	if err != nil {
		return common.ExitError{
			Code:    sysexit.EX_CONFIG,
			Message: "Failed to parse tuning: " + err.Error(),
		}
	}

	return tui.StartTUI(tui.Config{
		A4:        d.A4,
		Version:   d.ProgramVersion,
		Tunings:   tunings,
		Synth:     synth.NewSynth(d.WaveType, sample_rate),
		DebugMode: d.DebugMode,
	})
}
