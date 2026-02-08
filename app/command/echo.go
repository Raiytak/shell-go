package command

import (
  "fmt"
  "strings"
)

func EchoCmd(args []string) {
	fmt.Println(strings.Join(args, " "))

}
