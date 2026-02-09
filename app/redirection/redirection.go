package redirection

import (
	"os"
	"slices"
)

var redirectionSymbols = []string{">", "1>", "2>", ">>", "2>>"}

func SetRedirection(args []string, openFiles []*os.File) []string {
	if len(args) <= 1 {
		return args
	}

	var f *os.File
	for {
		if hasRedirection(args) {
			symbol, filePath := args[len(args)-2], args[len(args)-1]
			if isStdoutRedirection(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, &os.Stdout)
			} else if isStderrRedirection(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, &os.Stderr)
			} else if isStdoutAppend(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, &os.Stdout)
			} else if isStderrAppend(symbol) {
				f = setRedirection(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, &os.Stderr)
			}
			openFiles = append(openFiles, f)
			args = slices.Delete(args, len(args)-2, len(args))
		} else {
			return args
		}
	}
}

func setRedirection(filePath string, flag int, out **os.File) *os.File {
	f, err := os.OpenFile(filePath, flag, 0644)
	if err != nil {
		panic(err)
	}
	*out = f
	return f
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
	return arg == ">>"
}

func isStderrAppend(arg string) bool {
	return arg == "2>>"
}
