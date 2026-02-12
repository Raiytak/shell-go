package command

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func CdCmd(s Shell, args []string) (stdout []string, stderr []string) {
	if len(args) == 0 {
		return
	}

	dir := strings.Join(args, "/")
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
		return stdout, []string{fmt.Sprintf("cd: %s: No such file or directory", wDir)}
	}
	err = os.Chdir(wDir)
	if err != nil {
		return stdout, []string{fmt.Sprintf("error while changing directory")}
	}
	s.SetWorkingDir(wDir)
	return stdout, stderr
}
