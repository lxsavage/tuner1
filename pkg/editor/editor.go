package editor

import (
	"os"
	"os/exec"
	"runtime"
)

func EditFile(path_std_file string) error {
	if _, err := os.Stat(path_std_file); err != nil {
		return err
	}

	var cmd string
	var args []string

	if editor := os.Getenv("EDITOR"); len(editor) != 0 {
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
