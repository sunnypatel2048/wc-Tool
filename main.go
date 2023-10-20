package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// linesPtr := flag.Bool("l", false, "Measure the number of lines.")
	// wordsPtr := flag.Bool("w", false, "Measure the number of words.")
	bytesPtr := flag.Bool("c", false, "Measure the number of bytes in a file.")
	flag.Parse()

	if *bytesPtr {
		if filename := flag.Arg(0); filename != "" {
			file, err := os.ReadFile(filename)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("  ", len(file), filename)
		} else {
			fmt.Println("Filename argument not provided.")
			os.Exit(1)
		}
	}
}
