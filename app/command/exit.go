package command

import (
	"os"
)

func ExitCmd(s Shell) {
	writeHistory(s, s.GetHistoryFile())
	os.Exit(0)
}
