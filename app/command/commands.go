package command

import (
	"fmt"
	"io"
	"strings"

	"os/exec"
	"slices"
)

type Shell interface {
	WorkingDir() string
	SetWorkingDir(dir string)
	PathList() []string
	GetStdout() io.Writer
	GetStderr() io.Writer
	SetStdout(io.Writer)
	SetStderr(io.Writer)
	History() []string
	UpdateHistory(string)
}

var builtinCommands = []string{
	"echo",
	"exit",
	"type",
	"pwd",
	"cd",
	"history",
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
	command := strings.Join(append([]string{cmd}, args...), " ")
	cmdPath, isExec := CmdPath(cmd, s.PathList())
	switch {
	case isBuiltin(cmd):
		execBuiltinCmd(s, cmd, args)
	case isExec:
		execCmd(s, cmd, cmdPath, args)
	default:
		fmt.Printf("%s: command not found\n", cmd)
	}
	s.UpdateHistory(command)
	return
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
		EchoCmd(s, args)
	case "pwd":
		PwdCmd(s, args)
	case "cd":
		CdCmd(s, args)
	case "history":
		HistoryCmd(s, args)
	default:
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func execCmd(s Shell, cmd string, cmdPath string, args []string) {
	eCmd := exec.Command(cmd, args...)
	eCmd.Path = cmdPath
	eCmd.Stdout = s.GetStdout()
	eCmd.Stderr = s.GetStderr()
	eCmd.Run()
	return
}

func display(s Shell, lines []string) {
	for _, line := range lines {
		_, err := fmt.Fprintln(s.GetStdout(), line)
		if err != nil {
			panic(err)
		}
	}
}
