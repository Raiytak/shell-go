package main

import (
	"fmt"
  "os"
  "bufio"
  "path"
  "strings"
  "slices"
)


func displayType(c string, pathList []string) {
  // command in built-in commands
  if (slices.Contains(builtinCommands, c)) {
    fmt.Printf("%s is a shell builtin\n", c)
    return
  }

  // command in PATH
  for i := 0; i < len(pathList); i++ {
    p := pathList[i]
    _, err := os.Stat(p)
    // Path exists
    if err == nil {
      cmd_abs_path := path.Join(p, c)
      file_info, err := os.Stat(cmd_abs_path)
      // Command exists in path
      if err == nil {
        // Command is executable
        if isExecutable(file_info) {
          fmt.Printf("%s is %s\n", c, cmd_abs_path)
          return
        }
      }
    }
  }

  // command not found
  fmt.Printf("%s: not found\n", c)
}

func isExecutable(f os.FileInfo) bool{
  file_mode := f.Mode()
  is_executable := file_mode&0111 != 0
  return is_executable
}

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print
var builtinCommands = []string{"echo", "exit", "type"}

func main() {

  reader := bufio.NewReader(os.Stdin)

  _env_path, ok := os.LookupEnv("PATH")
  if !ok {
    fmt.Printf("%s is not set\n", _env_path)
  }
  env_paths := strings.Split(_env_path, string(os.PathListSeparator))

  for {
    fmt.Print("$ ")
    raw_input, err := reader.ReadString('\n')
    input := strings.TrimLeft(raw_input, " ")
    fields := strings.Fields(input)
    command := fields[0]
    if err != nil {
      fmt.Print("Error happened with the input\n")
      os.Exit(1)
    }

    switch command {
    case "exit":
      os.Exit(0)
    case "type":
      if len(fields) == 1 {
        fmt.Printf("%s: command not found\n", "")
      } else {
        evaluated_command := fields[1]
        displayType(evaluated_command, env_paths)
      }
    case "echo":
      fmt.Print(input[5:])
    default:
      fmt.Printf("%s: command not found\n", command)
    }
  }
}
