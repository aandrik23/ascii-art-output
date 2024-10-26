package funcs

import "fmt"

func PrintAsciiArt(sentences []string, textFile []string) {
	for i, word := range sentences {
		if word == "" {
			if i != 0 {
				fmt.Println() // Print a new line for blank words
			}
			continue
		}
		for h := 1; h < 9; h++ { // ASCII art character height is 8
			for k := 0; k < len(word); k++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[k])-32)*9+h { // Map the character to ASCII art lines
						fmt.Print(line) // Print the corresponding line for the character
					}
				}
			}
			fmt.Println() // New line after each line of ASCII art
		}
	}
}
