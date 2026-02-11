package command

import (
	"fmt"
	"slices"
)

func TypeCmd(s Shell, args []string) {
	var lines []string
	if len(args) == 0 {
		lines = []string{fmt.Sprintln(": not found")}
		display(s, lines)
		return
	}

	cmd := args[0]
	pathList := s.PathList()

	// Built-in Function
	if ok := slices.Contains(builtinCommands, cmd); ok {
		lines = []string{fmt.Sprintf("%s is a shell builtin", cmd)}
		display(s, lines)
		return
	}

	// Function Found in PATH
	cmdPath, isExec := CmdPath(cmd, pathList)
	if isExec {
		lines = []string{fmt.Sprintf("%s is %s", cmd, cmdPath)}
		display(s, lines)
		return
	}

	lines = []string{fmt.Sprintf("%s: not found", cmd)}
	display(s, lines)
}
