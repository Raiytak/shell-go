package redirection

import (
  "io"
	"os"
	"slices"
)

var redirectionSymbols = []string{">", "1>", "2>", ">>", "1>>", "2>>"}

func SetRedirection(args []string, openFiles []*os.File) []string {
	if len(args) <= 1 {
		return args
	}

	var f *os.File
  var fStdout, fStderr []*os.File
	for {
		if hasRedirection(args) {
			symbol, filePath := args[len(args)-2], args[len(args)-1]
			if isStdoutRedirection(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
        append(fStdout, f)
			} else if isStderrRedirection(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
        append(fStderr, f)
			} else if isStdoutAppend(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
        append(fStdout, f)
			} else if isStderrAppend(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
        append(fStderr, f)
			}
			openFiles = append(openFiles, f)
			args = slices.Delete(args, len(args)-2, len(args))
		} else {
      setMultiwriter(fStdout, &os.Stdout)
      setMultiwriter(fStderr, &os.Stderr)
			return args
		}
	}
}

func setRedirection(filePath string, flag int) *os.File {
	f, err := os.OpenFile(filePath, flag, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func setMultiwriter(f []*os.File, **os.File) {

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
