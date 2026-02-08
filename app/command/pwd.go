package command

import (
	"fmt"
)

func PwdCmd(s Shell, args []string) {
	if len(args) != 0 {
		fmt.Print("pwd: too many arguments\n")
		return
	}
	fmt.Printf("%s\n", s.WorkingDir())
}
