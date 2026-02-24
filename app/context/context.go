package context

import (
	"io"

	"github.com/codecrafters-io/shell-starter-go/app/history"
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

func (c *Shell) SaveHistory() error {
	_, err := history.Persist(c.History, "-w", c.HistFile)
	return err
}

type Command struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
