package main

import (
	"bufio"
	"fmt"
	"os"

	wegex "dict-parser/parsers"
)

func main() {

	// Get Dict Data in buffer
	file, error := os.Open("resources/dict.txt")

	var phase = wegex.PreStart

	if error != nil {
		fmt.Printf("I/O error: %s", error)
		return
	}

	// Close file on function exit
	defer file.Close()

	// Bring the file into a buffer
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	// i := 0

	for scanner.Scan() {

		// if i > 290 {
		// 	break
		// } else {
		// 	i++
		// }

		// if phase == parsingDefinition {
		// }
		if wegex.ParseDict(&phase, scanner.Text()) == 0 {
			break
		}

	}
}
