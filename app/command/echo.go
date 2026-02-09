package command

import (
	"fmt"
	"strings"
)

func EchoCmd(args []string) {
	firstString, lastString := args[0], args[len(args)-1]
	firstChar, lastChar := firstString[0], lastString[len(lastString)-1]

	if firstChar == '"' || firstChar == '\'' {
		if firstChar != lastChar {
			fmt.Print("Unclosed string\n")
			fmt.Printf("first string %s, last string %s", firstString, lastString)
			return
		}
		args[0] = firstString[1:]
		args[len(args)-1] = lastString[:len(lastString)-1]
	}
	fmt.Println(strings.Join(args, " "))
}
