package main

import (
  "os"
  "bufio"
  "io"
  "fmt"
   "github.com/fatih/color"
)

func implementation(input string, output string) {
  sourceFile, err := os.Open(input)

  if err != nil {
    c := color.New(color.FgRed).Add(color.Underline)
    c.Printf("Error Opening source file: '%s'\n", input)
    panic(err)
  }

  defer sourceFile.Close()

  destinationFile, err := os.Create(output)
  defer destinationFile.Close()

  reader := bufio.NewReader(sourceFile)
  writer := bufio.NewWriter(destinationFile)
  writer.WriteString("// Compiled \n ")

  // Copy the content from source to destination
  _, err = io.Copy(writer, reader)
  if err != nil {
    fmt.Println("Error copying file:", err)
    return
  }

  // Flush the writer to ensure all data is written to the file
  err = writer.Flush()
  if err != nil {
    fmt.Println("Error flushing writer:", err)
    return
  }
}

