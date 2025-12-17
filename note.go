package main

import "fmt"

type Note struct {
	pitch  string
	octave int
}

func (n Note) String() string {
	return fmt.Sprintf("%s%d", n.pitch, n.octave)
}
