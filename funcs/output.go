package funcs

import (
	"log"
	"os"
)

func PrintToFile(newFile *os.File, sentences []string, textFile []string) {
	for i, word := range sentences {
		if word == "" {
			if i != 0 {
				_, err := newFile.WriteString("\n")
				if err != nil {
					log.Fatal(err)
				}
			}
			continue
		}
		for h := 1; h < 9; h++ {
			for k := 0; k < len(word); k++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[k])-32)*9+h {
						_, err := newFile.WriteString(line)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
			_, err := newFile.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
