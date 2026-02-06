package main

import (
	"fmt"
  "os"
  "bufio"
  "strings"
  "slices"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
  reader := bufio.NewReader(os.Stdin)
  builtin_commands := []string{"echo", "exit", "type"}

   // if command not found
  for {
    fmt.Print("$ ")
    input, err := reader.ReadString('\n')
    fields := strings.Fields(input)
    command := fields[0]
    if err != nil {
      fmt.Print("Error happened with the input\n")
      os.Exit(1)
    }

    // Exit
    if (command == "exit") {
      os.Exit(0)
    } else if (command == "type") {
        evaluated_command := fields[1]
        if (slices.Contains(builtin_commands, evaluated_command)) {
          fmt.Printf("%s is a shell builtin\n", evaluated_command)
        } else {
          fmt.Printf("%s: not found\n", evaluated_command)
        }
    } else if (command == "echo") {
      fmt.Print(input[5:])
    } else {
      fmt.Printf("%s: command not found\n", command)
    }
  }
}
