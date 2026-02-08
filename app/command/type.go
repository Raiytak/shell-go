package command

import (
	"fmt"
	"slices"
)

var builtinCommands = []string{
	"echo",
	"exit",
	"type",
	"pwd",
	"cd",
}

func TypeCmd(s Shell, args []string) {
	if len(args) == 0 {
		fmt.Println(": not found")
		return
	}

	cmd := args[0]
	pathList := s.PathList()

	// Built-in Function
	if ok := slices.Contains(builtinCommands, cmd); ok {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}

	// Function Found in PATH
	cmdPath, isExec := CmdPath(cmd, pathList)
	if isExec {
		fmt.Printf("%s is %s\n", cmd, cmdPath)
		return
	}

	fmt.Printf("%s: not found\n", cmd)
}
