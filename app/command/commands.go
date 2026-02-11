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
  ResetHistory()
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

func RunCommand(s Shell, cmd string, args []string) (stdout []string, stderr []string) {
	command := strings.Join(append([]string{cmd}, args...), " ")
	cmdPath, isExec := CmdPath(cmd, s.PathList())
	s.UpdateHistory(command)
	switch {
	case isBuiltin(cmd):
		stdout, stderr = execBuiltinCmd(s, cmd, args)
	case isExec:
		stdout, stderr = execCmd(s, cmd, cmdPath, args)
	default:
		stderr = []string{fmt.Sprintf("%s: command not found", cmd)}
	}
	return stdout, stderr
}

func isBuiltin(cmd string) bool {
	return slices.Contains(builtinCommands, cmd)
}

func execBuiltinCmd(s Shell, cmd string, args []string) (stdout []string, stderr []string) {
	switch cmd {
	case "exit":
		ExitCmd()
	case "type":
		stdout, stderr = TypeCmd(s, args)
	case "echo":
		stdout, stderr = EchoCmd(args)
	case "pwd":
		stdout, stderr = PwdCmd(s, args)
	case "cd":
		stderr = CdCmd(s, args)
	case "history":
		stdout, stderr = HistoryCmd(s, args)
	default:
		stderr = []string{fmt.Sprintf("%s: command not found", cmd)}
	}
	return stdout, stderr
}

func execCmd(s Shell, cmd string, cmdPath string, args []string) (stdout []string, stderr []string) {
	eCmd := exec.Command(cmd, args...)
	eCmd.Path = cmdPath
	eCmd.Stdout = s.GetStdout()
	eCmd.Stderr = s.GetStderr()
	eCmd.Run()
	return stdout, stderr
}
