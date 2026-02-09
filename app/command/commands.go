package command

import (
	"fmt"
	"os"

	//	"os"
	"os/exec"
	"slices"
)

type Shell interface {
	WorkingDir() string
	SetWorkingDir(dir string)
	PathList() []string
}

// Other functions
func joinArgs(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := args[0]
	for i := 1; i < len(args); i++ {
		result += " " + args[i]
	}
	return result
}

func RunCommand(s Shell, cmd string, args []string) {
	if isBuiltin(cmd) {
		execBuiltinCmd(s, cmd, args)
		return
	} else if cmdPath, ok := CmdPath(cmd, s.PathList()); ok {
		execCmd(s, cmd, cmdPath, args)
	} else {
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func isBuiltin(cmd string) bool {
	return slices.Contains(builtinCommands, cmd)
}

func execBuiltinCmd(s Shell, cmd string, args []string) {
	switch cmd {
	case "exit":
		ExitCmd()
	case "type":
		TypeCmd(s, args)
	case "echo":
		EchoCmd(args)
	case "pwd":
		PwdCmd(s, args)
	case "cd":
		CdCmd(s, args)
	default:
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func execCmd(s Shell, cmd string, cmdPath string, args []string) {
	eCmd := exec.Command(cmd, args...)
	eCmd.Path = cmdPath
	eCmd.Stdout = os.Stdout
	eCmd.Stderr = os.Stderr
	eCmd.Run()
	return
}
