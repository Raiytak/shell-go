package command

import (
	"errors"
	"slices"
	"strings"
)

var delimiter = []byte{'"', '\''}

func Tokenize(line string) (name string, fields []string, err error) {
	fields, err = extractFields(line)
	if err != nil || len(fields) == 0 {
		return name, fields, err
	}

	name = fields[0]
	return name, fields[1:], err
}

func extractFields(line string) (fields []string, err error) {
	var token string
	var d byte

	if len(line) == 0 {
		return fields, nil
	}

	cleanedLine := strings.Join(strings.Fields(line), " ")
	for i := 0; i < len(cleanedLine); i++ {
		b := cleanedLine[i]
		switch {
		case (b == d || (b == ' ' && d == 0)):
			fields = append(fields, token)
			d = 0
			token = ""
		case (d == 0 && isDelimiter(b)):
			d = b
		case i == len(cleanedLine)-1:
			if b != '\n' {
				token += string(b)
			}
			fields = append(fields, token)
		default:
			token += string(b)
		}
	}
	if d != 0 {
		return []string{}, errors.New("unclosed quote\n")
	}
	return fields, nil
}

func isDelimiter(c byte) bool {
	return slices.Contains(delimiter, c)
}
