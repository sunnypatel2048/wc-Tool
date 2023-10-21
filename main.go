package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	linesPtr := flag.Bool("l", false, "Measure the number of lines in the file.")
	wordsPtr := flag.Bool("w", false, "Measure the number of words in the file.")
	bytesPtr := flag.Bool("c", false, "Measure the number of bytes in the file.")
	charPtr := flag.Bool("m", false, "Measure the number of characters in the file.")
	flag.Parse()

	if fileName := flag.Arg(0); fileName != "" {
		switch {
		case *bytesPtr:
			CalculateBytes(fileName)
		case *linesPtr:
			CalculateLines(fileName)
		case *wordsPtr:
			CalculateWords(fileName)
		case *charPtr:
			CalculateCharacters(fileName)
		default:
			fmt.Println("Out")
		}
	} else {
		fmt.Println("FileName argument not provided.")
		os.Exit(1)
	}
}
