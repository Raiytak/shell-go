package main

import "os"

func main() {
	shell := NewShell(os.Stdin, os.Stdout, os.Stderr)
	shell.Run()
}
