package command

import (
	"strings"
)

func EchoCmd(_ Shell, args []string) ([]string, []string) {
	return []string{strings.Join(args, " ")}, []string{}
}
