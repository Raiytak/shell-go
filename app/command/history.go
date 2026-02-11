package command

import "fmt"

func HistoryCmd(s Shell) {
  history := s.History()
	fmt.Print("history\n")
	for i, h := range history {
		fmt.Printf("    %d  %s\n", i, h)
	}
	return
}
