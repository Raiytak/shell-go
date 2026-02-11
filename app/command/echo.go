package command

import (
	"strings"
)

func EchoCmd(s Shell, args []string) {
	display(s, strings.Join(args, " "))
}
