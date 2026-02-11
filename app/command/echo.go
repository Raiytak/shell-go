package command

import (
	"strings"
)

func EchoCmd(args []string) ([]string, []string) {
	return []string{strings.Join(args, " ")}, []string{}
}
