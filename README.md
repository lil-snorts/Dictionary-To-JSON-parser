# Dictionary-To-JSON-parser

This is a GoLang script to parse the Gutenberg Websters dictionary https://www.gutenberg.org/ebooks/29765

## How to run

At the moment you need Go installed on your device

```bash
# Without the $filename it will default to resources/input.txt
go run application.go {$filename}
```

## What is this for?

This script was created to parse the Websters dictionary to be used in the flutter app, <a href="https://github.com/lil-snorts/daily-dict">daily dict</a>, as JSON files can be listed as resources in a flutter repo, and its easier to use struct->JSON libraies than code my own silly regex's
