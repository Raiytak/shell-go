package command

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/context"
)

type Exit struct{}

func (c *Exit) Run(ctx *context.Shell, _ *context.Command, _ []string) error {
	err := ctx.SaveHistory()
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
