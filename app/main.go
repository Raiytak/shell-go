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

  fmt.Print("$ ")
  text, err := reader.ReadString('\n')
  if err != nil {
    fmt.Print("Error happened with the input")
    os.Exit(1)
   }

   // if command not found
   fmt.Printf("%s command not found", text[:len(text)-1])
   os.Exit(1)
}
