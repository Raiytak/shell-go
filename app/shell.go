package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/redirection"
)

var delimiter = []byte{'"', '\''}

type Shell struct {
	reader   *bufio.Reader
	pathList []string
	wDir     string
	stdout   io.Writer
	stderr   io.Writer
	history  []string
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
		stdout:   os.Stdout,
		stderr:   os.Stderr,
		history:  []string{},
	}
}

func (s *Shell) Run() {
	stdout := os.Stdout
	stderr := os.Stderr
	var openFiles []*os.File
	for {
		s.SetStdout(stdout)
		s.SetStderr(stderr)
		fmt.Print("$ ")

		cmd, args, err := getCommand(s)
		if err != nil {
			fmt.Printf("error processing: %s %s\n", cmd, args)
			continue
		}

		args = redirection.SetRedirection(s, args, openFiles)

		command.RunCommand(s, cmd, args)
		closeFiles(openFiles)
	}
}

func (s *Shell) SetStdout(w io.Writer) {
	s.stdout = w
}

func (s *Shell) GetStdout() io.Writer {
	return s.stdout
}

func (s *Shell) SetStderr(w io.Writer) {
	s.stderr = w
}

func (s *Shell) GetStderr() io.Writer {
	return s.stderr
}

func getCommand(s *Shell) (string, []string, error) {
	var cmd string
	var args []string
	var err error

	line, err := s.reader.ReadString('\n')
	if err != nil {
		return cmd, args, err
	}

	fields, err := tokenize(line)
	if err != nil {
		return cmd, args, err
	}

	cmd = fields[0]
	if len(fields) > 1 {
		args = fields[1:]
	}
	return cmd, args, err
}

func tokenize(line string) ([]string, error) {
	var fields []string
	var token string
	var d byte

	cleanedLine := strings.Join(strings.Fields(line), " ")
	for i := 0; i < len(cleanedLine); i++ {
		b := cleanedLine[i]
		switch {
		case (b == d || (b == ' ' && d == 0)):
			fields = append(fields, token)
			d = 0
			token = ""
		case (d == 0 && isDelimiter(b)):
			d = b
		case i == len(cleanedLine)-1:
			if b != '\n' {
				token += string(b)
			}
			fields = append(fields, token)
		default:
			token += string(b)
		}
	}
	if d != 0 {
		return []string{}, errors.New("unclosed quote\n")
	}
	return fields, nil
}

func isDelimiter(c byte) bool {
	return slices.Contains(delimiter, c)
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

func (s *Shell) History() []string {
	return s.history
}

func (s *Shell) UpdateHistory(command string) {
	s.history = append(s.history, command)
}
