package main

import (
	"flag"
	"fmt"
)

func main() {
	linesPtr := flag.Bool("l", false, "Measure the number of lines.")
	wordsPtr := flag.Bool("w", false, "Measure the number of words.")

	flag.Parse()

	fmt.Println("Lines:", *linesPtr)
	fmt.Println("Words:", *wordsPtr)
	fmt.Println("\nArgs:", flag.Args())
}
