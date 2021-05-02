package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args
  //  Print all arguments
  fmt.Println(args)

  programName := args[0]
  fmt.Println("The binary name is: %s \n", programName)

  //  Using slice to drop the first element
  otherArgs := args[1:]
  fmt.Println(otherArgs)

  for index, arg := range otherArgs {
    fmt.Println("Arg %d = %s \n", index, arg)
  }
}
