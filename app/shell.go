package main

import (
	"bufio"
	"os"
  "fmt"
  "strings"
)

type Shell struct {
	reader   *bufio.Reader
	pathList []string
}

func NewShell() *Shell {
	pathEnv := os.Getenv("PATH")

	pathList := strings.Split(pathEnv, string(os.PathListSeparator))
	return &Shell{
		reader:   bufio.NewReader(os.Stdin),
		pathList: pathList,
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
		command := fields[0]
		args := fields[1:]

		RunCommand(command, args, s.pathList)
	}
}
