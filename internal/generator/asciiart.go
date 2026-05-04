package generator

import (
	"fmt"
	"os"
	"strings"
)

func AsciiGen(sentence, bannerFile string) string {
	// Construct the path to the banner file
	file := "./banners/" + bannerFile + ".txt"
	content, err := os.ReadFile(file)
	if err != nil {
		return fmt.Sprintf("Error: could not read file %s", err)
	}

	banner := strings.Split(string(content), "\n")

	sentence = strings.ReplaceAll(sentence, "\r\n", "\n")
	strSlice := strings.Split(sentence, "\n")

	var strBuild strings.Builder

	// Loop through each line of the input sentence
	for _, word := range strSlice {

		if word == "" {
			strBuild.WriteByte('\n')

			continue
		}

		for i := 1; i < 9; i++ {

			for _, ch := range word {
				if ch < ' ' || ch > '~' {
					return ""
				}
				// Calculate the starting index of that character in the banner file
				start := ((int(ch) - 32) * 9) + i

				// Append the corresponding ASCII art line for that character to the string builder
				strBuild.WriteString(banner[start])
			}
			strBuild.WriteByte('\n')
		}
	}

	return strBuild.String()
}
