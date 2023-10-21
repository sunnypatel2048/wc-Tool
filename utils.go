package main

import (
	"bytes"
)

func CalculateBytes(fileData []byte) int {
	return len(fileData)
}

func CalculateLines(fileData []byte) int {
	return bytes.Count(fileData, []byte("\n"))
}

func CalculateWords(fileData []byte) int {
	return len(bytes.Fields(fileData))
}

func CalculateCharacters(fileData []byte) int {
	return len(bytes.Runes(fileData))
}
