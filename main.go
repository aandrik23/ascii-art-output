package main

import (
	"fmt"
	"os"
	"output/funcs"
	"strings"
)

func main() {
	// Check if the number of command-line arguments is correct
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
		return
	}
	var text string
	var output string
	styleBanner := "standard"
	var outputfile string
	hasOutput := false
	if len(os.Args) == 4 {
		output = os.Args[1]
		text = os.Args[2]
		styleBanner = strings.ToLower(os.Args[3])
		hasOutput = true

		if !strings.HasPrefix(output, "--output=") || !strings.HasSuffix(output, ".txt") {
			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
			return
		}

		outputfile = output[9:] // Skips the "--output=" part
	}

	if len(os.Args) == 3 {
		output = os.Args[1]
		if strings.HasPrefix(output, "--output=") && strings.HasSuffix(output, ".txt") {

			text = os.Args[2]
			outputfile = output[9:] // Skips the "--output=" part
			hasOutput = true
		} else if strings.HasPrefix(output, "--output=") && !strings.HasSuffix(output, ".txt") {

			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
			return
		} else {

			text = os.Args[1]
			styleBanner = strings.ToLower(os.Args[2])
		}
	}

	if len(os.Args) == 2 {
		text = os.Args[1]
	}

	sepText := strings.Split(text, "\\n")

	file, err := os.ReadFile("banners/" + styleBanner + ".txt")
	if err != nil {
		fmt.Println(" banner does not exist.")
		return
	}

	str := string(file)
	str = strings.Replace(str, "\r\n", "\n", -1)
	lines := strings.Split(str, "\n")
	if !hasOutput {

		funcs.PrintAsciiArt(sepText, lines)
	} else {
		// If output file is specified, create or overwrite the file
		file, err := os.Create(outputfile)
		if err != nil {
			fmt.Println("ERROR! Usage: go run . [OPTION] [STRING] [BANNER] \nEX: go run . --output=<fileName.txt> \"something\" standard")
			return
		}

		funcs.PrintToFile(file, sepText, lines)
	}
}
