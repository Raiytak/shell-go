package command

import (
	"os"
)

func ExitCmd(s Shell, _ []string) (_ []string, _ []string) {
	writeHistory(s, s.GetHistoryFile())
	os.Exit(0)
	return
}
