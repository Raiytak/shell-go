package command

import (
	"errors"
	"fmt"
	"io"

	"github.com/codecrafters-io/shell-starter-go/app/context"
)

type Type struct{}

func (c *Type) Run(ctxSh *context.Shell, ctxCmd *context.Command, args []string) (err error) {
	if len(args) == 0 {
		return errors.New(": not found")
	}

	name := args[0]

	// Built-in Function
	_, builtin := Builtin[name]
	if builtin {
		io.WriteString(ctxCmd.Stdout, fmt.Sprintf("%s is a shell builtin\n", name))
		return err
	}

	// Function Found in PATH
	path, found := findCommand(name, ctxSh.PathList)
	if found {
		io.WriteString(ctxCmd.Stdout, fmt.Sprintf("%s is %s\n", name, path))
		return err
	}

	// Command not found
	errMsg := fmt.Sprintf("%s: not found\n", name)
	io.WriteString(ctxCmd.Stderr, errMsg)
	return errors.New(errMsg)
}
