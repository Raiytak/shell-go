package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "github.com/codecrafters-io/shell-starter-go/app/command"
)

type Shell struct {
	reader   *bufio.Reader
	pathList []string
	wDir     string
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
		wDir:     dir,
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
		line = strings.TrimRight(line, "\n")
		if line == "" {
			continue
		}

		trimmed_line := strings.TrimLeft(line, " \t")
		cmd, str_args, _ := strings.Cut(trimmed_line, " ")
    args := strings.Fields(str_args)

		command.RunCommand(s, cmd, args)
	}
}

func (s *Shell) WorkingDir() string {
  return s.wDir
}

func (s *Shell) SetWorkingDir(dir string) {
  s.wDir  = dir
}

func (s *Shell) PathList() []string {
  return s.pathList
}

