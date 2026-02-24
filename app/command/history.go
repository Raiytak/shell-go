package command

import (
	"errors"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/context"
	"github.com/codecrafters-io/shell-starter-go/app/history"
)

type History struct{}

func (c History) Run(ctxSh *context.Shell, ctxCmd *context.Command, args []string) (err error) {
	var output string
	switch {
	case len(args) == 0:
		output, err = history.All(ctxSh.History)
	case len(args) == 1:
		limit, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		output, err = history.Limit(ctxSh.History, limit)
	case len(args) == 2:
		action := args[0]
		filename := args[1]
		updatedHistory, err := history.Persist(ctxSh.History, action, filename)
		if action == "-r" {
			updatedHistory = slices.Insert(updatedHistory, 0, strings.Join(args, " "))
		}
		ctxSh.History = updatedHistory
		if err != nil {
			return err
		}
		ctxSh.History = updatedHistory
	default:
		return errors.New("wrong argument")
	}
	if err != nil {
		return err
	}
	if len(output) > 0 {
		io.WriteString(ctxCmd.Stdout, output+"\n")
	}
	return err
}
