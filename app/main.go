package main

import (
	"fmt"
  "os"
  "bufio"
  "slices"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
  reader := bufio.NewReader(os.Stdin)
  functions := []string{"exit"}

   // if command not found
  for {
    fmt.Print("$ ")
    raw_text, err := reader.ReadString('\n')
    text := raw_text[:len(raw_text) - 1]
    if err != nil {
      fmt.Print("Error happened with the input\n")
      os.Exit(1)
    }
    if !(slices.Contains(functions, text)) {
      fmt.Printf("%s: command not found\n", text)
    }
    if (text == "exit") {
      os.Exit(0)
    }
  }
}
