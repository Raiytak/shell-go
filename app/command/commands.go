package command

import (
	"fmt"
	"io"
	"os"
	"strings"

	"os/exec"
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
	GetHistoryFile() string
	IsBuiltin(string) bool
	Commands() map[string]Command
}

type Command func(s Shell, args []string) (stdout []string, stderr []string)

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
	s.UpdateHistory(command)

	cmdPath, cmdFound := findCommand(cmd, s.PathList())
	switch {
	case s.IsBuiltin(cmd):
		stdout, stderr = execBuiltinCmd(s, cmd, args)
	case cmdFound:
		stdout, stderr = execCmd(s, cmd, cmdPath, args)
	default:
		stderr = []string{fmt.Sprintf("%s: command not found", cmd)}
	}
	return stdout, stderr
}

func execBuiltinCmd(s Shell, cmd string, args []string) (stdout []string, stderr []string) {
	command, ok := s.Commands()[cmd]
	if !ok {
		return stdout, []string{fmt.Sprintf("no builtin command %s", cmd)}
	}
	return command(s, args)
}

func execCmd(s Shell, cmd string, cmdPath string, args []string) (stdout []string, stderr []string) {
	eCmd := exec.Command(cmd, args...)
	eCmd.Path = cmdPath
	eCmd.Stdout = s.GetStdout()
	eCmd.Stderr = s.GetStderr()
	eCmd.Run()
	return stdout, stderr
}

func EmptyLine(s string) bool {
	return (s == "\n" || s == "")
}

func EnsureFileExists(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}
