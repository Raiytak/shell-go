package redirection

import (
	"fmt"
	"io"
	"os"
)

type Shell interface {
	WorkingDir() string
	SetWorkingDir(dir string)
	PathList() []string
	SetStdout(io.Writer)
	SetStderr(io.Writer)
	GetStdout() io.Writer
	GetStderr() io.Writer
	SetOpenFiles([]*os.File)
}

func Redirect(s Shell, stdout string, stderr string) (err error) {
	var text string
	var output io.Writer
	if len(stdout) > 0 {
		text = stdout
		output = s.GetStdout()
		_, err = fmt.Fprint(output, text)
	}
	if len(stderr) > 0 {
		text = stderr
		output = s.GetStderr()
		_, err = fmt.Fprint(output, text)
	}
	return err
}
