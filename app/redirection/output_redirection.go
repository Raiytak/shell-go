package redirection

import (
	"io"
	"os"
	"slices"

	"github.com/codecrafters-io/shell-starter-go/app/context"
)

var redirectionSymbol = []string{">", "1>", "2>", ">>", "1>>", "2>>"}
var truncFile = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
var appendFile = os.O_CREATE | os.O_WRONLY | os.O_APPEND

func SetRedirection(ctxCmd *context.Command, fields []string) (args []string, openedFiles []*os.File, err error) {
	var f *os.File
	fStdout := []io.Writer{}
	fStderr := []io.Writer{}
	args = make([]string, len(fields))
	copy(args, fields)

	if len(args) <= 1 {
		return args, openedFiles, err
	}

	if !redirected(args) {
		return args, openedFiles, err
	}
	for {
		if redirected(args) {
			symbol, filename := args[len(args)-2], args[len(args)-1]
			switch symbol {
			case ">", "1>":
				f, err = openFile(filename, truncFile)
				var w io.Writer = f
				fStdout = append(fStdout, w)
			case "2>":
				f, err = openFile(filename, truncFile)
				var w io.Writer = f
				fStderr = append(fStderr, w)
			case ">>", "1>>":
				f, err = openFile(filename, appendFile)
				var w io.Writer = f
				fStdout = append(fStdout, w)
			case "2>>":
				f, err = openFile(filename, appendFile)
				var w io.Writer = f
				fStderr = append(fStderr, w)
			}
			if err != nil {
				return args, openedFiles, err
			}
			args = slices.Delete(args, len(args)-2, len(args))
			openedFiles = append(openedFiles, f)
		} else {
			ctxCmd.Stdout = io.MultiWriter(append(fStdout, ctxCmd.Stdout)...)
			ctxCmd.Stderr = io.MultiWriter(append(fStderr, ctxCmd.Stderr)...)
			return args, openedFiles, err
		}
	}
}

func openFile(filename string, flag int) (*os.File, error) {
	return os.OpenFile(filename, flag, 0644)
}

func redirected(args []string) bool {
	if len(args) <= 1 {
		return false
	}
	return isOutputRedirection(args[len(args)-2])
}

func isOutputRedirection(arg string) bool {
	return slices.Contains(redirectionSymbol, arg)
}

func RedirectedStdout(args []string) bool {
	return slices.Contains(args, ">") || slices.Contains(args, "1>") || slices.Contains(args, "1>>") || slices.Contains(args, ">>")
}

func RedirectedStderr(args []string) bool {
	return slices.Contains(args, "2>") || slices.Contains(args, "2>>")
}
