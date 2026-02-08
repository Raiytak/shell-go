package command

import (
	"os"
	"path/filepath"
)

func CmdPath(cmd string, pathList []string) (string, bool) {
	for _, dir := range pathList {
		fullPath := filepath.Join(dir, cmd)

		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		if ok := isExecutable(info); ok {
			return fullPath, true
		}
	}
	return "", false
}

func isExecutable(info os.FileInfo) bool {
	return info.Mode()&0111 != 0
}
