package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/chzyer/readline"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/context"
	"github.com/codecrafters-io/shell-starter-go/app/history"
	"github.com/codecrafters-io/shell-starter-go/app/redirection"
)

type Shell struct {
	Context context.Shell
	Reader  *readline.Instance
}

func NewShell(stdin io.Reader, stdout io.Writer, stderr io.Writer) *Shell {
	reader, err := readline.New("$ ")
	if err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic("error gathering the working directory")
	}
	os.Setenv("HISTFILE", "test_history")
	histFile := os.Getenv("HISTFILE")
	fmt.Print(histFile)
	history, err := history.Initialize(histFile)
	if err != nil {
		panic(err)
	}

	return &Shell{
		Context: context.Shell{
			Dir:      dir,
			PathList: strings.Split(os.Getenv("PATH"), string(os.PathListSeparator)),
			History:  history,
			HistFile: histFile,
			Stdin:    stdin,
			Stdout:   stdout,
			Stderr:   stderr,
		},
		Reader: reader,
	}
}

func (s *Shell) Run() {
	for {
		line, err := s.Reader.Readline()
		if err != nil {
			io.WriteString(s.Context.Stderr, err.Error()+"\n")
			os.Exit(1)
		}
		if line == "" {
			continue
		}

		cmds := redirection.Subcommands(line)
		if err := runPipeline(cmds, &s.Context); err != nil {
			io.WriteString(s.Context.Stderr, err.Error()+"\n")
			continue
		}
	}
}

func runPipeline(cmds []string, ctxSh *context.Shell) (err error) {
	var name string
	var openedFiles []*os.File
	var args, fields []string
	prevReader := ctxSh.Stdin
	var wg sync.WaitGroup
	errCh := make(chan error, len(cmds))
	for i, cmd := range cmds {
		isLast := i == len(cmds)-1
		ctxCmd := context.Command{Stdin: prevReader, Stdout: io.Discard, Stderr: io.Discard}
		name, fields, err = command.Tokenize(cmd)
		if err != nil {
			return err
		}
		if !isLast {
			r, w, err := os.Pipe()
			ctxCmd.Stdout = w
			prevReader = r

			args, openedFiles, err = redirection.SetRedirection(&ctxCmd, fields)
			if err != nil {
				return err
			}

			nameCopy := name
			argsCopy := args
			ctxCmdCopy := ctxCmd

			wg.Add(1)

			go func(ctxSh *context.Shell, ctxCmd *context.Command, name string, args []string) {
				defer wg.Done()
				defer w.Close()

				if err := command.Run(name, args, ctxSh, ctxCmd); err != nil {
					errCh <- err
				}
			}(ctxSh, &ctxCmdCopy, nameCopy, argsCopy)

		} else {
			if !redirection.RedirectedStdout(fields) {
				ctxCmd.Stdout = ctxSh.Stdout
			}
			if !redirection.RedirectedStderr(fields) {
				ctxCmd.Stderr = ctxSh.Stderr
			}
			args, openedFiles, err = redirection.SetRedirection(&ctxCmd, fields)
			if err != nil {
				return err
			}

			wg.Add(1)
			go func(ctxSh *context.Shell, ctxCmd *context.Command, args []string) {
				defer wg.Done()

				if err := command.Run(name, args, ctxSh, ctxCmd); err != nil {
					errCh <- err
				}
			}(ctxSh, &ctxCmd, args)
		}
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	CloseFiles(openedFiles)
	return nil
}

func CloseFiles(files []*os.File) error {
	for _, f := range files {
		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}
