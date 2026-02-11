package command

import "fmt"

func HistoryCmd(s Shell) {
	history := s.History()
	display(s, fmt.Sprintf("history\n"))
	for i, h := range history {
		display(s, fmt.Sprintf("    %d  %s\n", i, h))
	}
	return
}
