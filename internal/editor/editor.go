// Package editor provides a function to edit a file using the system's default editor.
package editor

import (
	"os"
	"os/exec"
	"runtime"
)

func EditFile(stdFilePath string) error {
	if _, err := os.Stat(stdFilePath); err != nil {
		return err
	}

	var cmd string
	var args []string

	if editor := os.Getenv("EDITOR"); len(editor) != 0 {
		cmd = editor
		args = []string{stdFilePath}
	} else {
		switch runtime.GOOS {
		case "windows":
			cmd = "cmd"
			args = []string{"/c", "start", "", stdFilePath}
		case "darwin":
			cmd = "open"
			args = []string{"-t", stdFilePath}
		default:
			cmd = "xdg-open"
			args = []string{stdFilePath}
		}
	}
	proc := exec.Command(cmd, args...)
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	return proc.Run()
}
