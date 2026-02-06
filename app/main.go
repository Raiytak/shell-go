package main

import (
	"fmt"
  "os"
  "bufio"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
  reader := bufio.NewReader(os.Stdin)

   // if command not found
  for {
    fmt.Print("$ ")
    raw_text, err := reader.ReadString('\n')
    text := raw_text[:len(raw_text) - 1]
    if err != nil {
      fmt.Print("Error happened with the input\n")
      os.Exit(1)
    }

    // Exit
    if (text == "exit") {
      os.Exit(0)
    } else if (text[:4] == "echo") {
      fmt.Print(text[5:], "\n")
    } else {
      fmt.Printf("%s: command not found\n", text)
    }
  }
}
