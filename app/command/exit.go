package command

import (
	"fmt"
	"os"
)

func ExitCmd(s Shell) {
	fmt.Print(appendHistory(s, os.Getenv("HISTFILE")))
	os.Exit(0)
}
