package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/redirection"
)

type Shell struct {
	reader   *bufio.Reader
	pathList []string
	wDir     string
  stdout *os.File
  stderr *os.File
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
    stdout: &os.Stdout,
    stderr: &os.Stderr,
	}
}

func (s *Shell) Run() {
	stdout, stderr := os.Stdout, os.Stderr
	var openFiles []*os.File
	for {
		os.Stdout = stdout
		os.Stderr = stderr
		fmt.Print("$ ")

		cmd, args, err := processUserInput(s)
		if err != nil {
			fmt.Printf("error processing: %s %s", cmd, args)
			continue
		}

		args = redirection.SetRedirection(args, openFiles)

		command.RunCommand(s, cmd, args)
		closeFiles(openFiles)
	}
}

func processUserInput(s *Shell) (cmd string, args []string, err error) {
	line, err := s.reader.ReadString('\n')
	if err != nil {
		return
	}

	fields := strings.Fields(line)
	cmd = fields[0]
	if len(fields) > 1 {
		args = fields[1:]
	}
	return
}

func (s *Shell) WorkingDir() string {
	return s.wDir
}

func (s *Shell) SetWorkingDir(dir string) {
	s.wDir = dir
}

func (s *Shell) PathList() []string {
	return s.pathList
}

func closeFiles(openFiles []*os.File) {
	for _, file := range openFiles {
		file.Close()
	}
	openFiles = []*os.File{}
}
