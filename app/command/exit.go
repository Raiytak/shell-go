package command

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/context"
)

type Exit struct{}

func (c *Exit) Run(_ *context.Shell, _ *context.Command, _ []string) error {
	os.Exit(0)
	return nil
}
