package command

import "os"

func ExitCmd(s Shell) {
  writeHistory(s, os.Getenv("HISTFILE"))
	os.Exit(0)
}
