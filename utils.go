package main

import (
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
