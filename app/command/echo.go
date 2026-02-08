package command

import (
  "fmt"
  "strings"
)

func EchoCmd(args []string) {
  firstChar, lastChar := args[0][0], args[len(args)][len(len(args))]
  if firstChar == "'" || firstChar == "\"" {
    if firstChar != lastChar {
      fmt.Print("Unclosed string\n")
      return
    }
    args[0] = args[0][1:]
    args[len(args)] = args[len(args)][:len(args[len(args])]
  }
	fmt.Println(strings.Join(args, " "))

}
