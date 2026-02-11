package command

import (
	"fmt"
)

func PwdCmd(s Shell, args []string) {
	var lines []string
	if len(args) != 0 {
		lines = []string{("pwd: too many arguments")}
	} else {
		lines = []string{s.WorkingDir()}
	}
	display(s, lines)
}
