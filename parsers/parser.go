package parsers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	wordRegexStr      = `^[A-Z](?:[A-Z0-9\. ';-]*)$`
	dictEndRegexStr   = `^\s*[*]{3} END.*$`
	dictStartRegexStr = `^\s*[*]{3} START.*$`
	newLineRegexStr   = "^$"
)

var (
	WordRegex      = regexp.MustCompile(wordRegexStr)
	DictEndRegex   = regexp.MustCompile(dictEndRegexStr)
	DictStartRegex = regexp.MustCompile(dictStartRegexStr)
	EmptyLineRegex = regexp.MustCompile(newLineRegexStr)
)

type DictWord struct {
	Name           string
	Pronounciation string
	Descriptions   []string
}

var (
	name           string
	pronounciation string
	descriptions   []string
)

func ParseFileToJSON(filepath string) {
	entireDict := []DictWord{}

	file, error := os.Open(filepath)

	if error != nil {
		fmt.Printf("Cant find %s\n", filepath)
		return
	}

	// Bring the file into a buffer
	scanner := bufio.NewScanner(file)
	// might need to defer this
	defer file.Close()

	active := false
	bigText := ""

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), `"`, "'")

		if len(entireDict) > 0 {
			continue
		}

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
			fmt.Println(name)
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

	fmt.Print(string(text))

	err := os.WriteFile("./output/parsed.json", text, 0666)

	if err != nil {
		fmt.Print(err)
	}

	return
}
