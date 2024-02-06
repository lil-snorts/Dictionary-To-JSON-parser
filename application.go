package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	// Get Dict Data in buffer
	file, error := os.Open("resources/dict.txt")
	_DICTIONARY_WORD_REGEX, _ := regexp.Compile("^[A-Z][A-Z0-9\\. -]*$")
	_DICTIONARY_END_REGEX, _ := regexp.Compile("^\\s*[\\*]{3} END.*$")
	if error != nil {
		fmt.Printf("I/O error: %s", error)
		return
	}
	// Close file on function exit
	defer file.Close()

	// Bring the file into a buffer
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	var lastWord = "$"
	fmt.Println("{\n\t")
	for scanner.Scan() {
		line := scanner.Text()
		if _DICTIONARY_END_REGEX.MatchString(line) {
			// fmt.Printf(line)
			fmt.Printf("\tfinished reading\n")
			break
		}

		if _DICTIONARY_WORD_REGEX.MatchString(line) && lastWord != line {
			fmt.Printf("\"\n},\n{\n\t\"word\": \"%s\",\n\t\"definition\": \"", line)
			lastWord = line
		} else {
			fmt.Printf("\t%s", line)
		}
	}
	fmt.Println("\t\"\n}")

	// fmt.Println(lastWord)
}
