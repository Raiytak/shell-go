package redirection

import (
	"io"
	"os"
	"slices"
)

type Shell interface {
	WorkingDir() string
	SetWorkingDir(dir string)
	PathList() []string
	SetStdout(io.Writer)
	SetStderr(io.Writer)
}

var redirectionSymbols = []string{">", "1>", "2>", ">>", "1>>", "2>>"}

func SetRedirection(s Shell, args []string, openFiles []*os.File) []string {
	if len(args) <= 1 {
		return args
	}

	var f *os.File
	fStdout := []*os.File{}
	fStderr := []*os.File{}
  stdoutRedirected, stderrRedirected := false, false
	for {
		if hasRedirection(args) {
			symbol, filePath := args[len(args)-2], args[len(args)-1]
			if isStdoutRedirection(symbol) {
				f = openFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
				fStdout = append(fStdout, f)
        stdoutRedirected = true
			} else if isStderrRedirection(symbol) {
				f = openFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
				fStderr = append(fStderr, f)
        stderrRedirected = true
			} else if isStdoutAppend(symbol) {
				f = openFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
				fStdout = append(fStdout, f)
        stdoutRedirected = true
			} else if isStderrAppend(symbol) {
				f = openFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
				fStderr = append(fStderr, f)
        stderrRedirected = true
			}
			openFiles = append(openFiles, f)
			args = slices.Delete(args, len(args)-2, len(args))
		} else {
      if !stdoutRedirected {
        fStdout = append(fStdout, os.Stdout)
      }
      if !stderrRedirected {
        fStderr = append(fStderr, os.Stderr)
      }
			setStdout(s, fStdout)
			setStderr(s, fStderr)
			return args
		}
	}
}

func openFile(filePath string, flag int) *os.File {
	f, err := os.OpenFile(filePath, flag, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func setStdout(s Shell, files []*os.File) {
	writers := make([]io.Writer, 0, len(files))
	for _, f := range files {
		writers = append(writers, f)
	}
	w := io.MultiWriter(writers...)
	s.SetStdout(w)
}

func setStderr(s Shell, files []*os.File) {
	writers := make([]io.Writer, 0, len(files))
	for _, f := range files {
		writers = append(writers, f)
	}
	w := io.MultiWriter(writers...)
	s.SetStderr(w)
}

func hasRedirection(args []string) bool {
	if len(args) > 1 {
		if isRedirectionSymbol(args[len(args)-2]) {
			return true
		}
	}
	return false
}

func isRedirectionSymbol(arg string) bool {
	return slices.Contains(redirectionSymbols, arg)
}

func isStdoutRedirection(arg string) bool {
	return (arg == ">" || arg == "1>")
}

func isStderrRedirection(arg string) bool {
	return arg == "2>"
}

func isStdoutAppend(arg string) bool {
	return (arg == ">>" || arg == "1>>")
}

func isStderrAppend(arg string) bool {
	return arg == "2>>"
}
