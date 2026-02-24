package command

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/context"
)

type Cd struct{}

func (c Cd) Run(ctx *context.Shell, _ *context.Command, args []string) error {
	input := strings.Join(args, "/")
	input = strings.ReplaceAll(input, "~", home())
	var target string
	switch {
	case len(args) == 0:
		target = home()
	case input[0] == '/':
		target = input
	default:
		target = path.Clean(path.Join(ctx.Dir, input))
	}
	_, err := os.Stat(target)
	if err != nil {
		io.WriteString(ctx.Stderr, fmt.Sprintf("cd: %s: No such file or directory\n", target))
		return err
	}
	ctx.Dir = target
	return err
}

func home() string {
	return os.Getenv("HOME")
}
