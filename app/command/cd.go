package command

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func CdCmd(s Shell, args []string) {
	dir := strings.Join(args, "/")
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
		fmt.Printf("cd: %s: No such file or directory\n", wDir)
		return
	}
	err = os.Chdir(wDir)
	if err != nil {
		fmt.Print("error while changing directory")
	}
	s.SetWorkingDir(wDir)
}
