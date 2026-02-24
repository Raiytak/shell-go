package command

import (
	"errors"
	"io"

	"github.com/codecrafters-io/shell-starter-go/app/context"
)

type Pwd struct{}

func (c Pwd) Run(ctxSh *context.Shell, ctxCmd *context.Command, args []string) error {
	if len(args) != 0 {
		return errors.New("too many arguments")
	}
	io.WriteString(ctxCmd.Stdout, ctxSh.Dir+"\n")
	return nil
}
