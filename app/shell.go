package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "slices"
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
  stdout := os.Stdout
	for {
    os.Stdout = stdout
		fmt.Print("$ ")

    cmd, args, err := processUserInput(s)
    if err != nil {
      fmt.Printf("error processing: %s %s", cmd, args)
      continue
    }

    // Redirect output
    if len(args) >= 2 && (args[len(args)-2] == ">" || args[len(args)-2] == "1>") {
      file, err := os.Create(args[len(args) -1])
      if err != nil {
        fmt.Printf("error redirecting output: %s", err)
        continue
      }
      defer file.Close()

      os.Stdout = file
      args = slices.Delete(args, len(args)-2, len(args))
      // args = args[:len(args)-2]
      // eCmd.Stderr = os.Stderr
    }
		command.RunCommand(s, cmd, args)
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
  s.wDir  = dir
}

func (s *Shell) PathList() []string {
  return s.pathList
}

