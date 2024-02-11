package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	preStart int = iota
	beforeFirstWord
	parsingWord
	parsingPronounciationFirstLine
	parsingPronounciation
	parsingDefinitionFirstLine
	parsingDefinition
	newLineInDefinition
)

const (
	WordRegexStr      = `^[A-Z](?:[A-Z0-9\. ';-]*)$`
	DictEndRegexStr   = `^\s*[\*]{3} END.*$`
	DictStartRegexStr = `^\s*[\*]{3} START.*$`
	NewLineRegexStr   = "^$"
)

var (
	WordRegex      = regexp.MustCompile(WordRegexStr)
	DictEndRegex   = regexp.MustCompile(DictEndRegexStr)
	DictStartRegex = regexp.MustCompile(DictStartRegexStr)
	NewLineRegex   = regexp.MustCompile(NewLineRegexStr)
)

func main() {

	// Get Dict Data in buffer
	file, error := os.Open("resources/dict.txt")

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
	// i := 0

	for scanner.Scan() {

		// if i > 290 {
		// 	break
		// } else {
		// 	i++
		// }

		// if phase == parsingDefinition {
		// }
		if ParseDict(&phase, scanner.Text()) == 0 {
			break
		}

	}
}

func ParseDict(phase *int, line string) int {
	if *phase == preStart && DictStartRegex.MatchString(line) {
		*phase = beforeFirstWord
		return 1
	} else if *phase == preStart {
		return 1
	}

	if DictEndRegex.MatchString(line) {
		fmt.Println("\t\"\n}")
		fmt.Printf("\n\tfinished reading\n")
		return 0
	}

	if NewLineRegex.MatchString(line) {

		if *phase == parsingDefinition {
			*phase = newLineInDefinition
			fmt.Printf(`"`)
		} else if *phase == parsingPronounciation {
			*phase = parsingDefinitionFirstLine
			fmt.Printf(`",`)
		}

	} else if WordRegex.MatchString(line) {
		if *phase == newLineInDefinition {
			fmt.Printf("\n\t]")
		} else if *phase != beforeFirstWord {
			fmt.Printf("]\"\n},\n")
		}

		fmt.Printf("\n},\n{\n\t\"word\": \"%s\",", line)
		*phase = parsingPronounciationFirstLine

	} else if *phase == parsingPronounciationFirstLine {
		fmt.Printf("\n\t"+`"pronounciation": "%s`, strings.Replace(line, `"`, `'`, -1))
		*phase = parsingPronounciation
	} else if *phase == parsingPronounciation {
		fmt.Printf(" %s", line)
	} else if *phase == parsingDefinitionFirstLine {
		fmt.Printf(
			"\n\t\"definition\": [\n\t\"%s",
			strings.Replace(line, `"`, `'`, -1))

		*phase = parsingDefinition

	} else if *phase == parsingDefinition || *phase == newLineInDefinition {

		if *phase == newLineInDefinition {
			fmt.Printf(",\n\t\"")
			*phase = parsingDefinition
		}

		fmt.Printf("%s", strings.Replace(line, `"`, `'`, -1))
	}
	return 1
}
