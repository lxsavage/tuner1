package main

import "fmt"

type Note struct {
	pitch  string
	octave int
}

func (n Note) String() string {
	format_specifier := "%s%d"
	if len(n.pitch) == 1 {
		format_specifier = "%s %d"
	}
	return fmt.Sprintf(format_specifier, n.pitch, n.octave)
}
