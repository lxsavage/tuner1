package tuner

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"lxsavage/tuner1/internal/common"
	"lxsavage/tuner1/internal/ui"
	"lxsavage/tuner1/pkg/editor"
	"lxsavage/tuner1/pkg/sysexit"
	"os"
	"path/filepath"
	"strings"
)

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

func listTemplates(path_std_file string) {
	standards, err := getStandardsFileContents(path_std_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read standards file: %s\n", err)
		os.Exit(sysexit.EX_IOERR)
	}

	fmt.Print(sprintStandards(standards))
}

func Execute(program_version string, show_version *bool, list_templates *bool, edit_standards *bool, tuning_template *string, reference *float64, standards *string, wave_type *string) error {
	if *show_version {
		fmt.Println(program_version)
		return nil
	}
	is_template := len(*tuning_template) > 0 && (*tuning_template)[0] == '+'
	should_use_default_std_file := *edit_standards || *list_templates || is_template

	path_std_file := *standards
	if should_use_default_std_file && len(path_std_file) == 0 {
		path_std_file = getStandardsConfigFilepath()
	}

	if *list_templates {
		listTemplates(path_std_file)
		return nil
	}

	if *edit_standards {
		if err := editor.EditFile(path_std_file); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return common.ExitError{
					Code: sysexit.EX_IOERR,
					Message: fmt.Sprintf("Unable to find standards file at \"%s\"; "+
						"perform a reinstall with the script to get the default one",
						path_std_file),
				}
			}
			fmt.Fprintf(os.Stderr, "Unable to launch editor: %s\n", err)
			os.Exit(sysexit.EX_UNAVAILABLE)

			return common.ExitError{
				Code:    sysexit.EX_UNAVAILABLE,
				Message: "Unable to launch editor: " + err.Error(),
			}
		}
		return nil
	}

	if len(*tuning_template) == 0 {
		return common.ExitError{
			Code: sysexit.EX_USAGE,
			Message: fmt.Sprintf("Please pass in a tuning specifier\n\nTry:\n"+
				"  %s -ls\n", os.Args[0]),
		}
	}

	tuning_csv := *tuning_template
	if (*tuning_template)[0] == '+' {
		standards, err := getStandardsFileContents(path_std_file)
		if err != nil {
			return common.ExitError{
				Code:    sysexit.EX_IOERR,
				Message: "Failed to open standards file: " + err.Error(),
			}
		}

		csv, err := getStandard(standards, *tuning_template)
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

	return ui.StartTUI(program_version, tunings, *reference, *wave_type)
}
