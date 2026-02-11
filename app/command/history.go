package command

import "fmt"

func HistoryCmd(s Shell) {
	history := s.History()
	for i, h := range history {
		display(s, fmt.Sprintf("    %d  %s", i+1, h))
	}
	return
}
