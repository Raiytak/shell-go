package command

import (
	"fmt"
	"slices"
)

func TypeCmd(s Shell, args []string) (stdout []string, stderr []string) {
	if len(args) == 0 {
		return stdout, []string{fmt.Sprintln(": not found")}
	}

	cmd := args[0]
	pathList := s.PathList()

	// Built-in Function
	if ok := slices.Contains(builtinCommands, cmd); ok {
		return []string{fmt.Sprintf("%s is a shell builtin", cmd)}, stderr
	}

	// Function Found in PATH
	cmdPath, isExec := CmdPath(cmd, pathList)
	if isExec {
		return []string{fmt.Sprintf("%s is %s", cmd, cmdPath)}, stderr
	}

	// Command not found
	return stdout, []string{fmt.Sprintf("%s: not found", cmd)}
}
