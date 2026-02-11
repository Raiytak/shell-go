package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HistoryCmd(s Shell, args []string) (stdout []string, stderr []string) {
	history := s.History()
	switch {
	case len(args) == 0:
		stdout, stderr = completeHistory(history)
	case len(args) == 1:
		limit := args[0]
		stdout, stderr = limitHistory(history, limit)
	case len(args) == 2:
		action := args[0]
		filename := args[1]
		stdout, stderr = historyPersistence(s, action, filename)
	default:
		return stdout, []string{"wrong argument"}
	}
	return stdout, stderr
}

func completeHistory(history []string) (stdout []string, stderr []string) {
	stdout, stderr = getHistory(history, 0)
	return stdout, stderr
}

func limitHistory(history []string, limit string) (stdout []string, stderr []string) {
	l, err := strconv.Atoi(limit)
	if err != nil {
		return stdout, []string{fmt.Sprintf("%s: argument not handled", limit)}
	}

	start := len(history) - l
	if start < 0 {
		start = 0
	}
	stdout, stderr = getHistory(history, start)
	return stdout, stderr
}

func historyPersistence(s Shell, action string, filename string) (stdout []string, stderr []string) {
	switch action {
	case "-r":
		stderr = importHistory(s, filename)
	case "-w":
		stderr = writeHistory(s, filename)
		s.ResetHistory()
	case "-a":
		stderr = appendHistory(s, filename)
		s.ResetHistory()
	}
	return stdout, stderr
}

func importHistory(s Shell, filename string) (stderr []string) {
	lines, err := ReadHistory(filename)
	if err != nil {
		stderr = []string{"error reading history file"}
	}
	for _, line := range lines {
		s.UpdateHistory(line)
	}
	return stderr
}

func ReadHistory(filename string) (lines []string, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return lines, errors.New("error opening file")
	}
	content := string(data)
	lines = strings.Split(content, "\n")
	return lines, err
}

func writeHistory(s Shell, filename string) (stderr []string) {
	return saveHistory(filename, s.History())
}

func appendHistory(s Shell, filename string) (stderr []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{fmt.Sprintf("error opening file: %s", filename)}
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	for _, l := range s.History() {
		lines = append(lines, l)
	}
	return saveHistory(filename, lines)
}

func saveHistory(filename string, lines []string) (stderr []string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return []string{fmt.Sprintf("error opening file: %s", filename)}
	}
	defer f.Close()

	_, err = f.WriteString(strings.Join(append(lines, ""), "\n"))
	if err != nil {
		return []string{fmt.Sprintf("error writing file: %s", filename)}
	}
	return stderr
}

func getHistory(history []string, start int) (stdout []string, stderr []string) {
	if start < 0 || len(history) < start {
		return stdout, []string{"no such event"}
	}
	for i := start; i < len(history); i++ {
		stdout = append(stdout, fmt.Sprintf("    %d  %s", i+1, history[i]))
	}
	return stdout, stderr
}
