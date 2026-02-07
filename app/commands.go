package main

import (
  "fmt"
  "os"
  "slices"
  "os/exec"
)

var builtinCommands = []string{
  "echo",
  "exit",
  "type",
  "pwd",
}

// Builtin command
func echoCmd(args []string) {
  fmt.Println(joinArgs(args))
}

func exitCmd() {
  os.Exit(0)

}

func typeCmd(s *Shell, args []string) {
  if len(args) == 0 {
    fmt.Println(": not found")
    return
  }

  cmd := args[0]
  pathList := s.pathList

  // Built-in Function
  if ok := slices.Contains(builtinCommands, cmd); ok {
    fmt.Printf("%s is a shell builtin\n", cmd)
    return
  }

  // Function Found in PATH
  cmdPath, isExec := FindInPath(cmd, pathList)
  if isExec {
    fmt.Printf("%s is %s\n", cmd, cmdPath)
    return
  }

  fmt.Printf("%s: not found\n", cmd)
}

func pwdCmd(s *Shell, args []string) {
  fmt.Printf("%s\n", s.wDir)
}

// Other functions
func joinArgs(args []string) string {
  if len(args) == 0 {
    return ""
  }

  result := args[0]
  for i := 1; i < len(args); i++ {
    result += " " + args[i]
  }
  return result
}

func RunCommand(s *Shell, cmd string, args []string) {
  // Builtin command
  if ok := slices.Contains(builtinCommands, cmd); ok {
    switch cmd {
    case "exit":
      exitCmd()
    case "type":
      typeCmd(s, args)
    case "echo":
      echoCmd(args)
    case "pwd":
      pwdCmd(s, args)
    default:
      fmt.Printf("%s: command not found\n", cmd)
    }
    return
  }

  // Command executable
  cmdPath, isExec := FindInPath(cmd, s.pathList)
  if isExec {
    e_cmd := exec.Command(cmd, args...)
    e_cmd.Path = cmdPath
    e_cmd.Stdout = os.Stdout
    e_cmd.Stderr = os.Stderr
    err := e_cmd.Run()
    if err != nil {
      fmt.Printf("Command failed: %v\n", err)
    }
    return
  }

  // Not builtin function
  fmt.Printf("%s: command not found\n", cmd)
}
