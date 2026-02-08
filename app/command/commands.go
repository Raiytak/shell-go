package command

import (
	"fmt"
	"os"
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
	if ok := slices.Contains(builtinCommands, cmd); ok {
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
		return
	}

	// Command executable
	execCmd(s, cmd, args)
}

func execCmd(s Shell, cmd string, args []string) {
	cmdPath, isExec := FindInPath(cmd, s.PathList())

	if isExec {
		eCmd := exec.Command(cmd, args...)
		eCmd.Path = cmdPath
		eCmd.Stdout = os.Stdout
		eCmd.Stderr = os.Stderr
		err := eCmd.Run()
		if err != nil {
			fmt.Printf("Command failed: %v\n", err)
		}
	} else {
		fmt.Printf("%s: command not found\n", cmd)
	}
	return
}
