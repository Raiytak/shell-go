package command

import (
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
		stderr = readHistory(s, filename)
	case "-w":
		stderr = writeHistory(s, filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
	case "-a":
		stderr = writeHistory(s, filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
	}
	return stdout, stderr
}

func readHistory(s Shell, filename string) (stderr []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{"error opening file"}
	}

	content := string(data)
	history := strings.Split(content, "\n")
	for _, line := range history {
		if line != "" {
			s.UpdateHistory(line)
		}
	}
	return stderr
}

func writeHistory(s Shell, filename string, flag int) (stderr []string) {
	f, err := os.OpenFile(filename, flag, 0644)
	if err != nil {
		return []string{fmt.Sprintf("error opening file: %s", filename)}
	}
	defer f.Close()

	_, err = f.WriteString(append(strings.Join(s.History(), "\n"), ""))
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
