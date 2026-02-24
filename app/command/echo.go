package command

import (
	"io"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/context"
)

type Echo struct{}

func (c *Echo) Run(_ *context.Shell, ctx *context.Command, args []string) error {
	_, err := io.WriteString(ctx.Stdout, strings.Join(args, " ")+"\n")
	return err
}
