package command

import (
	"fmt"
	"strconv"
)

func HistoryCmd(s Shell, args []string) (stdout []string, stderr []string) {
	history := s.History()
	switch {
	case len(args) == 0:
		stdout = completeHistory(history)
	case len(args) == 1:
		limit := args[0]
		stdout, stderr = limitHistory(history, limit)
	case len(args) > 1:
		return stdout, []string{"too many arguments"}
	}
	return stdout, stderr
}

func completeHistory(history []string) (stdout []string) {
	for i, h := range history {
		stdout = append(stdout, fmt.Sprintf("    %d  %s", i+1, h))
	}
	stdout = append(stdout, fmt.Sprintf("    %d  history", len(history)+1))
	return stdout
}

func limitHistory(history []string, limit string) (stdout []string, stderr []string) {
	l, err := strconv.Atoi(limit)
	if err != nil {
		return stdout, []string{fmt.Sprintf("%s: argument not handled", limit)}
	}

	start := len(history) - l + 1
	if start < 0 {
		start = 0
	}
	for i := start; i < len(history); i++ {
		stdout = append(stdout, fmt.Sprintf("    %d  %s", i+1, history[i]))
	}
	stdout = append(stdout, fmt.Sprintf("    %d  history %d", len(history)+1, l))
	return stdout, stderr
}
