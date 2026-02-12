package command

import "os"

func ExitCmd(s Shell) {
  appendHistory(s, os.Getenv("HISTFILE"))
	os.Exit(0)
}
