package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/chzyer/readline"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/redirection"
)

var delimiter = []byte{'"', '\''}

type Shell struct {
	reader    *readline.Instance
	pathList  []string
	wDir      string
	stdin     io.Reader
	stdout    io.Writer
	stderr    io.Writer
	history   []string
	histFile  string
	openFiles []*os.File
}

func NewShell(stdin io.Reader, stdout io.Writer, stderr io.Writer) *Shell {
	pathEnv := os.Getenv("PATH")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("error gathering the working directory")
		os.Exit(1)
	}
	reader, err := readline.New("$ ")
  histFile := os.Getenv("HISTFILE")
  if histFile == "" {
    histFile = ".bash_history"
  }
	err = command.EnsureFileExists(histFile)
	if err != nil {
		panic(err)
	}
	history := command.ReadHistory(histFile)
	history = slices.DeleteFunc(history, command.EmptyLine)

	pathList := strings.Split(pathEnv, string(os.PathListSeparator))
	return &Shell{
		reader:    reader,
		pathList:  pathList,
		wDir:      dir,
		stdin:     stdin,
		stdout:    stdout,
		stderr:    stderr,
		history:   history,
		histFile:  histFile,
		openFiles: []*os.File{},
	}
}

func (s *Shell) Run() {
	var stdout, stderr []string
	defaultStdout := s.stdout
	defaultStderr := s.stderr
	defer closeFiles(s)

	for {
		s.SetStdout(defaultStdout)
		s.SetStderr(defaultStderr)
		fmt.Print("$ ")

		cmd, args, err := getCommand(s)
		if err != nil {
			fmt.Printf("error processing: %s %s\n", cmd, args)
			continue
		}

		args = redirection.SetRedirection(s, args)

		stdout, stderr = command.RunCommand(s, cmd, args)
		err = display(s, stdout, stderr)
		if err != nil {
			panic(err)
		}

		closeFiles(s)
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

	line, err := s.reader.Readline()
	if err != nil {
		return cmd, args, err
	}

	fields, err := tokenize(line)
	if err != nil || len(fields) == 0 {
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

	if len(line) == 0 {
		return fields, nil
	}

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

func closeFiles(s *Shell) {
	for _, file := range s.GetOpenFiles() {
		file.Close()
	}
	s.SetOpenFiles([]*os.File{})
}

func (s *Shell) History() []string {
	return s.history
}

func (s *Shell) UpdateHistory(line string) {
	s.history = append(s.history, line)
}

func (s *Shell) ResetHistory() {
	s.history = []string{}
}

func (s *Shell) GetOpenFiles() []*os.File {
	return s.openFiles
}

func (s *Shell) SetOpenFiles(openFiles []*os.File) {
	s.openFiles = openFiles
}

func (s *Shell) GetHistoryFile() string {
  return s.histFile
}

func display(s *Shell, stdout []string, stderr []string) (err error) {
	var lines []string
	var output io.Writer
	switch {
	case len(stdout) > 0:
		lines = stdout
		output = s.GetStdout()
	case len(stderr) > 0:
		lines = stderr
		output = s.GetStderr()
	}
	for _, line := range lines {
		_, err := fmt.Fprintln(output, line)
		if err != nil {
			return err
		}
	}
	return err
}
