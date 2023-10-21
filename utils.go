package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func CalculateBytes(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("  ", fileInfo.Size(), fileName)
}

func CalculateLines(fileName string) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := bytes.Count(fileData, []byte{'\n'})
	fmt.Println("  ", lines, fileName)
}

func CalculateWords(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Set scanner's delimiter to be a space.
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	fmt.Println("  ", wordCount, fileName)
}
