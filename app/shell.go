package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	reader   *bufio.Reader
	pathList []string
  wDir string
}

func NewShell() *Shell {
	pathEnv := os.Getenv("PATH")
  dir, err := os.Getwd()
  if err != nil {
    fmt.Print("error gathering the working directory")
    os.Exit(1)
  }

	pathList := strings.Split(pathEnv, string(os.PathListSeparator))
	return &Shell{
		reader:   bufio.NewReader(os.Stdin),
		pathList: pathList,
    wDir: dir,
	}
}

func (s *Shell) Run() {
	for {
		fmt.Print("$ ")

		line, err := s.reader.ReadString('\n')
		if err != nil {
			fmt.Print("input error")
			return
		}

		line = strings.TrimLeft(line, " ")
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		cmd := fields[0]
		args := fields[1:]

		RunCommand(s, cmd, args)
	}
}
