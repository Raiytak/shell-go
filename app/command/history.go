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
	display(s, lines)
	return nil
}

func completeHistory(history []string) []string {
	var lines []string
	for i, h := range history {
		lines = append(lines, fmt.Sprintf("    %d  %s", i+1, h))
	}
	lines = append(lines, fmt.Sprintf("    %d  history", len(history) + 1))
	return lines
}

func limitHistory(history []string, limit string) ([]string, error) {
	var lines []string
	l, err := strconv.Atoi(limit)
	if err != nil {
		return lines, errors.New(fmt.Sprintf("%s: argument not handled", limit))
	}

  start := len(history) - l + 1
  if start < 0 {
    start = 0
  }
	for i := start; i < len(history); i++ {
		lines = append(lines, fmt.Sprintf("    %d  %s", i+1, history[i]))
	}
	lines = append(lines, fmt.Sprintf("    %d  history %d", len(history) + 1, l))
	return lines, nil
}
