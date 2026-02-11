package command

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func CdCmd(s Shell, args []string) {
	dir := strings.Join(args, "/")
	var lines []string
	if len(dir) == 0 {
		return
	}

	wDir := ""
	if dir[0] == '/' {
		wDir = dir
	} else if dir[0] == '~' {
		wDir = os.Getenv("HOME") + dir[1:]
	} else {
		wDir = path.Clean(path.Join(s.WorkingDir(), dir))
	}
	_, err := os.Stat(wDir)
	if err != nil {
		lines = []string{fmt.Sprintf("cd: %s: No such file or directory", wDir)}
		display(s, lines)
		return
	}
	err = os.Chdir(wDir)
	if err != nil {
		lines = []string{fmt.Sprintf("error while changing directory")}
		display(s, lines)
	}
	s.SetWorkingDir(wDir)
}
