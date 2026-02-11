package command

import (
	"strings"
)

func EchoCmd(s Shell, args []string) {
	lines := []string{strings.Join(args, " ")}
	display(s, lines)
}
