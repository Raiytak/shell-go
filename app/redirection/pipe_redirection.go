package redirection

import "strings"

var pipeSymbol = "|"

func Subcommands(line string) []string {
	return strings.Split(line, pipeSymbol)
}
