package command

import (
	"fmt"
	"strings"
)

func PwdCmd(s Shell, args string) {
	fields := strings.Fields(args)
	if len(fields) != 0 {
		fmt.Print("pwd: too many arguments\n")
		return
	}
	fmt.Printf("%s\n", s.WorkingDir())
}
