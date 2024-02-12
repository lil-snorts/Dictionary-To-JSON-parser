package parsers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	WordRegex      = regexp.MustCompile(`^[A-Z](?:[A-Z0-9\. ';-]*)$`)
	DictEndRegex   = regexp.MustCompile(`^\s*[*]{3} END.*$`)
	DictStartRegex = regexp.MustCompile(`^\s*[*]{3} START.*$`)
	EmptyLineRegex = regexp.MustCompile("^$")
	name           string
	pronounciation string
	descriptions   []string
)

type DictWord struct {
	Name           string
	Pronounciation string
	Descriptions   []string
}

func ParseFileToJSON(filepath string) {
	entireDict := []DictWord{}

	file, error := os.Open(filepath)

	if error != nil {
		fmt.Printf("Cant find %s\n", filepath)
		return
	}
	defer file.Close()

	// Bring the file into a buffer
	scanner := bufio.NewScanner(file)

	active := false
	bigText := ""

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), `"`, `'`)

		if !active {
			if DictStartRegex.MatchString(line) {
				active = true
			}
			continue
		}

		if WordRegex.MatchString(line) {
			if name != "" {
				entireDict = append(entireDict, DictWord{
					Name:           name,
					Pronounciation: pronounciation,
					Descriptions:   descriptions,
				})
			}
			name = line
			pronounciation = ""
			descriptions = []string{}
			fmt.Printf("%s, ", name)
		} else if EmptyLineRegex.MatchString(line) {
			// This only happens when its multi lined defn.
			// or the transition from pronounce to defn.
			if pronounciation == "" {
				pronounciation = bigText
			} else {
				descriptions = append(descriptions, bigText)
			}
			bigText = ""

		} else if DictEndRegex.MatchString(line) {
			active = false
		} else {
			bigText += line
		}
	}
	text, _ := json.Marshal(entireDict)

	// 0666 is Read, Write but not execute
	err := os.WriteFile("./output/parsed.json", text, 0666)

	if err != nil {
		fmt.Print(err)
	}
}
