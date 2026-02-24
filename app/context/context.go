package context

import (
	"io"
)

type Shell struct {
	Dir      string
	PathList []string
	History  []string
	HistFile string
	Stdin    io.Reader
	Stdout   io.Writer
	Stderr   io.Writer
}

type Command struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
