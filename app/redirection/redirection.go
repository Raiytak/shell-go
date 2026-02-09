package redirection

import (
	"os"
	"slices"
)

var redirectionSymbols = []string{">", "1>", "2>"}

func SetRedirection(args []string, openFiles []os.File) []string {
	if len(args) <= 1 {
		return args
	}

	for i := 0; i < 2; i++ {
		if hasRedirection(args) {
			symbol, file := args[len(args)-2], args[len(args)-1]
			if isStdoutRedirectionSymbol(symbol) {
				setStdoutRedirection(file, openFiles)
			} else if isStderrRedirectionSymbol(symbol) {
				setStderrRedirection(file, openFiles)
			}
			args = slices.Delete(args, len(args)-2, len(args))
		}
	}
  return args
}

func setStdoutRedirection(filePath string, openFiles []os.File) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	os.Stdout = file
	openFiles = append(openFiles, *file)
}

func setStderrRedirection(filePath string, openFiles []os.File) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	os.Stderr = file
	openFiles = append(openFiles, *file)
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

func isStdoutRedirectionSymbol(arg string) bool {
	return (arg == ">" || arg == "1>")
}

func isStderrRedirectionSymbol(arg string) bool {
	return arg == "2>"
}
