package command

import (
	"slices"

	"github.com/google/shlex"
)

var delimiter = []byte{'"', '\''}

func Tokenize(line string) (name string, fields []string, err error) {
	fields, err = extractFields(line)
	if err != nil {
		return name, fields, err
	}

	name = fields[0]
	return name, fields[1:], err
}

func extractFields(line string) (args []string, err error) {
	args, err = shlex.Split(line)
	return args, err
}

func isDelimiter(c byte) bool {
	return slices.Contains(delimiter, c)
}
