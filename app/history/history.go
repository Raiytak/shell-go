package history

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

func All(history []string) (string, error) {
	return Limit(history, len(history))
}

func Limit(history []string, limit int) (output string, err error) {
	start := len(history) - limit
	if start < 0 {
		start = 0
	}
	return Display(history, start)
}

func Persist(history []string, action string, path string) (updatedHistory []string, err error) {
	switch action {
	case "-r":
		updatedHistory, err = Import(path)
	case "-w":
		err = writeHistory(history, path)
	case "-a":
		err = appendHistory(history, path)
	}
	return updatedHistory, err
}

func Initialize(path string) ([]string, error) {
	history, err := Import(path)
	if err != nil {
		f, err := os.OpenFile(path, os.O_CREATE, 0644)
		if err != nil {
			return history, err
		}
		return history, f.Close()
	}
	return history, err
}

func Import(path string) (history []string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return history, err
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	history = slices.DeleteFunc(lines, EmptyLine)
	return history, err
}

func writeHistory(history []string, path string) error {
	return saveLines(path, history)
}

func appendHistory(history []string, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]
	for _, l := range history {
		lines = append(lines, l)
	}
	return saveLines(path, lines)
}

func saveLines(path string, lines []string) (err error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(strings.Join(append(lines, ""), "\n"))
	return err
}

func Display(history []string, start int) (output string, err error) {
	if start < 0 || len(history) < start {
		return output, errors.New("no such event")
	}
	for i := start; i < len(history); i++ {
		output += fmt.Sprintf("    %d  %s", i+1, history[i])
		if i < len(history)-1 {
			output += "\n"
		}
	}
	return output, err
}

func Append(history []string, name string, args []string) []string {
	return append(history, name+" "+strings.Join(args, " "))
}

func reset(history []string) {
	history = []string{}
}

func EmptyLine(s string) bool {
	return (s == "\n" || s == "")
}
