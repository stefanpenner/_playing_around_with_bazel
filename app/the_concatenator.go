package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go_concat <output_file> <input_file1> [input_file2 ...]")
		os.Exit(1)
	}

	output := os.Args[1]
	inputFiles := os.Args[2:]

	destinationFile, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer destinationFile.Close()

	writer := bufio.NewWriter(destinationFile)
	defer writer.Flush()

	_, err = writer.WriteString("// Compiled\n")
	if err != nil {
		fmt.Printf("Error writing header: %v\n", err)
		os.Exit(1)
	}

	for _, input := range inputFiles {
		sourceFile, err := os.Open(input)
		if err != nil {
			fmt.Printf("Error opening input file %s: %v\n", input, err)
			os.Exit(1)
			continue
		}

		reader := bufio.NewReader(sourceFile)
		_, err = reader.WriteTo(writer)
		if err != nil {
			fmt.Printf("Error copying from %s: %v\n", input, err)
			os.Exit(1)
		}
		writer.WriteString("\n")

		sourceFile.Close()
	}
}
