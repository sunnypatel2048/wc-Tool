package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	output := ""

	linesPtr := flag.Bool("l", false, "Measure the number of lines in the file.")
	wordsPtr := flag.Bool("w", false, "Measure the number of words in the file.")
	bytesPtr := flag.Bool("c", false, "Measure the number of bytes in the file.")
	charPtr := flag.Bool("m", false, "Measure the number of characters in the file.")
	flag.Parse()
	fileName := flag.Arg(0)

	var fileData []byte

	if fileName != "" {
		file, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		fileData = file

	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		fileData = stdin
	}

	switch {
	case *bytesPtr:
		output += fmt.Sprintf("%d ", CalculateBytes(fileData))
	case *linesPtr:
		output += fmt.Sprintf("%d ", CalculateLines(fileData))
	case *wordsPtr:
		output += fmt.Sprintf("%d ", CalculateWords(fileData))
	case *charPtr:
		output += fmt.Sprintf("%d ", CalculateCharacters(fileData))
	default:
		output += fmt.Sprintf("    %d   %d  %d ", CalculateLines(fileData), CalculateWords(fileData), CalculateBytes(fileData))
	}

	if fileName != "" {
		output += fileName
	}

	fmt.Println(output)
}
