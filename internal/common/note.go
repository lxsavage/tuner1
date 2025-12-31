package common

import "fmt"

type Note struct {
	Pitch  string
	Octave int
}

func (n Note) String() string {
	format_specifier := "%s%d"
	if len(n.Pitch) == 1 {
		format_specifier = "%s %d"
	}
	return fmt.Sprintf(format_specifier, n.Pitch, n.Octave)
}
