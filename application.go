package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	// Get Dict Data in buffer
	file, error := os.Open("resources/dict.txt")

	wordRegex := regexp.MustCompile("^[A-Z][A-Z0-9\\. '-]*$")
	dictEndRegex := regexp.MustCompile("^\\s*[\\*]{3} END.*$")
	dictStartRegex := regexp.MustCompile("^\\s*[\\*]{3} START.*$")
	newLineRegex := regexp.MustCompile("^$")

	const (
		preStart int = iota
		beforeFirstWord
		parsingWord
		parsingPronounciation
		parsingDefinitionFirstLine
		parsingDefinition
		newLineInDefinition
	)

	var phase = preStart

	if error != nil {
		fmt.Printf("I/O error: %s", error)
		return
	}

	// Close file on function exit
	defer file.Close()

	// Bring the file into a buffer
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	i := 0

	for scanner.Scan() {

		if i > 190 {
			break
		} else {
			i++
		}

		// read new line from buffer
		line := scanner.Text()

		if phase == preStart && dictStartRegex.MatchString(line) {
			phase = beforeFirstWord
			continue
		} else if phase == preStart {
			continue
		}

		if dictEndRegex.MatchString(line) {
			fmt.Println("\t\"\n}")
			fmt.Printf("\n\tfinished reading\n")
			break
		}

		if wordRegex.MatchString(line) {
			if phase != beforeFirstWord {
				fmt.Printf("\"]\"\n},\n")
			}

			fmt.Printf("{\n\t\"word\": \"%s\",", line)
			phase = parsingPronounciation

		} else if phase == parsingPronounciation {
			fmt.Printf("\n\t\"pronounciation\": \"%s\",", strings.Replace(line, "\"", "\\\"", -1))
			phase = parsingDefinitionFirstLine

		} else if newLineRegex.MatchString(line) {

			if phase == parsingDefinition {
				fmt.Printf("\n")
				phase = newLineInDefinition
			}

			fmt.Printf(",")

		} else if phase == parsingDefinitionFirstLine {
			fmt.Printf("\n\t\"definition\": [\n\t\t\"%s", strings.Replace(line, "\"", "\\\"", -1))
			phase = parsingDefinition

		} else if phase == parsingDefinition || phase == newLineInDefinition {
			if phase == newLineInDefinition {
				fmt.Printf("\",\n\t\"")
				phase = parsingDefinition
			}
			fmt.Printf("%s", strings.Replace(line, "\"", "\\\"", -1))
		}

		// fmt.Println("\n\t\t\t", phase)
	}
}
