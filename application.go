package main

import (
	"os"

	"dict-parser/parsers"
)

func main() {

	// Get Dict Data in buffer
	var filename string
	if len(os.Args) > 0 {
		filename = os.Args[0]
	} else {
		filename = "./resources/input.txt"
	}

	parsers.ParseFileToJSON(filename)
}
