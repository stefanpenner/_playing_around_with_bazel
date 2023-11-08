package main
import (
  "os"
  "bufio"
  "io"
  "fmt"
)

func main() {
  output := os.Args[2]
  input := os.Args[3]
  sourceFile, err := os.Open(input)

  if err != nil {
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
