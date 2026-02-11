package command

import (
	"fmt"
)

func PwdCmd(s Shell, args []string) {
	var lines []string
	if len(args) != 0 {
		lines = []string{("pwd: too many arguments\n")}
	} else {
		lines = []string{fmt.Sprintf("%s\n", s.WorkingDir())}
	}
	display(s, lines)
}
