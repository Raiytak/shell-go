package command

import (
	"fmt"
	"slices"
)

func TypeCmd(s Shell, args []string) {
	if len(args) == 0 {
		display(s, fmt.Sprintln(": not found"))
		return
	}

	cmd := args[0]
	pathList := s.PathList()

	// Built-in Function
	if ok := slices.Contains(builtinCommands, cmd); ok {
		display(s, fmt.Sprintf("%s is a shell builtin", cmd))
		return
	}

	// Function Found in PATH
	cmdPath, isExec := CmdPath(cmd, pathList)
	if isExec {
		display(s, fmt.Sprintf("%s is %s", cmd, cmdPath))
		return
	}

	display(s, fmt.Sprintf(": not found", cmd))
}
