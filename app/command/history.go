package command

import (
	"errors"
	"fmt"
	"strconv"
)

func HistoryCmd(s Shell, args []string) error {
	history := s.History()
	var lines []string
	var err error
	switch {
	case len(args) == 0:
		lines = completeHistory(history)
	case len(args) == 1:
		limit := args[0]
		lines, err = limitHistory(history, limit)
		if err != nil {
			return err
		}
	case len(args) > 1:
		return errors.New("too many arguments")
	}
	lines = append(lines, fmt.Sprintf("    %d  history", len(history)))
	display(s, lines)
	return nil
}

func completeHistory(history []string) []string {
	var lines []string
	for i, h := range history {
		lines = append(lines, fmt.Sprintf("    %d  %s", i+1, h))
	}
	return lines
}

func limitHistory(history []string, limit string) ([]string, error) {
	var lines []string
	l, err := strconv.Atoi(limit)
	if err != nil {
		return lines, errors.New(fmt.Sprintf("%s: argument not handled", limit))
	}

	lHist := len(history)
	for i := lHist - l; i < lHist; i++ {
		lines = append(lines, fmt.Sprintf("    %d  %s", i, history[i]))
	}
	return lines, nil
}
