package main

import (
	"os"

	"dict-parser/parsers"
)

func main() {

	// Get Dict Data in buffer
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "./resources/input.txt"
	}

	parsers.ParseFileToJSON(filename)
}
