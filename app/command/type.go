package command

import (
	"fmt"
	"slices"
)

func TypeCmd(s Shell, cmd string) {
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

	display(s, fmt.Sprintf("%s: not found", cmd))
}
