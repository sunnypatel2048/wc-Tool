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

func CalculateCharacters(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Set scanner's split function to be `bufio.ScanRunes`.
	// A rune is a single character in a UTF-8 encoded string.
	scanner.Split(bufio.ScanRunes)

	charCount := 0
	for scanner.Scan() {
		charCount++
	}

	fmt.Println("  ", charCount, fileName)
}

func CalculateDefaults(fileName string) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fBytes := len(fileData)
	fLines := bytes.Count(fileData, []byte{'\n'})

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Set scanner's delimiter to be a space.
	scanner.Split(bufio.ScanWords)

	fWords := 0
	for scanner.Scan() {
		fWords++
	}

	fmt.Printf("    %d   %d  %d %s\n", fLines, fWords, fBytes, fileName)
}
