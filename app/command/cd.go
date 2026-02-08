package command

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func CdCmd(s Shell, args string) {
	fields := strings.Fields(args)
	if len(fields) == 0 {
		return
	}

	wDir := ""
	if args[0] == '/' {
		wDir = args
	} else if args[0] == '~' {
		wDir = os.Getenv("HOME") + args[1:]
	} else {
		wDir = path.Clean(path.Join(s.WorkingDir(), args))
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
