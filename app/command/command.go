package command

import (
	"errors"
	"fmt"
	"io"

	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/context"
	"github.com/codecrafters-io/shell-starter-go/app/history"
)

type Command interface {
	Run(*context.Shell, *context.Command, []string) error
}

type Shell interface {
	ResetHistory()
	GetHistoryFile() string
}

var Builtin = map[string]Command{
	"cd":      &Cd{},
	"echo":    &Echo{},
	"exit":    &Exit{},
	"type":    &Type{},
	"pwd":     &Pwd{},
	"history": &History{},
}

func Run(name string, args []string, ctxSh *context.Shell, ctxCmd *context.Command) error {
	ctxSh.History = history.Append(ctxSh.History, name, args)
	bCmd, builtin := Builtin[name]
	path, found := findCommand(name, ctxSh.PathList)
	switch {
	case builtin:
		return bCmd.Run(ctxSh, ctxCmd, args)
	case found:
		cmd := exec.Command(name, args...)
		cmd.Path = path
		return runExecutable(cmd, ctxCmd)
	default:
		errMsg := fmt.Sprintf("%s: command not found\n", name)
		io.WriteString(ctxCmd.Stderr, errMsg)
		return errors.New(errMsg)
	}
}

func runExecutable(cmd *exec.Cmd, ctx *context.Command) error {
	cmd.Stdin = ctx.Stdin
	cmd.Stdout = ctx.Stdout
	cmd.Stderr = ctx.Stderr
	return cmd.Run()
}

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
