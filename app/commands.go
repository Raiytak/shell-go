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
}

func echoCmd(args []string, _ []string) {
  fmt.Println(joinArgs(args))
}

func exitCmd(_ []string, _ []string) {
  os.Exit(0)

}

func typeCmd(args []string, pathList []string) {
  if len(args) == 0 {
    fmt.Println(": not found")
    return
  }

  cmd := args[0]

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

func RunCommand(cmd string, args []string, pathList []string) {
  // Builtin command
  if ok := slices.Contains(builtinCommands, cmd); ok {
    switch cmd {
    case "exit":
      exitCmd(args, pathList)
    case "type":
      typeCmd(args, pathList)
    case "echo":
      echoCmd(args, pathList)
    default:
      fmt.Printf("%s: command not found\n", cmd)
    }
    return
  }

  // Command executable
  cmdPath, isExec := FindInPath(cmd, pathList)
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
