package command

import "fmt"

func HistoryCmd(s Shell) {
	history := s.History()
	display(s, fmt.Sprintf("history"))
	for i, h := range history {
		display(s, fmt.Sprintf("    %d  %s", i, h))
	}
	return
}
