package generator

import (
	"fmt"
	"html"
	"os"
	"strings"
)

func AsciiGen(sentence, bannerFile, subString, color string) string {
	file := "./banners/" + bannerFile + ".txt"
	content, err := os.ReadFile(file)
	if err != nil {
		return fmt.Sprintf("Error: could not read file %s", err)
	}

	banner := strings.Split(string(content), "\n")
	sentence = strings.ReplaceAll(sentence, "\r\n", "\n")
	strSlice := strings.Split(sentence, "\n")
	var strBuild strings.Builder

	for _, word := range strSlice {
		if word == "" {
			strBuild.WriteByte('\n')
			continue
		}

		// Map which characters fall within the target substring
		colorMap := make([]bool, len(word))
		if subString != "" {
			idx := 0
			for {
				i := strings.Index(word[idx:], subString)
				if i == -1 {
					break
				}
				for j := 0; j < len(subString); j++ {
					colorMap[idx+i+j] = true
				}
				idx += i + len(subString)
			}
		}

		// Build the ASCII art row by row
		for i := 1; i < 9; i++ {
			for charIdx, ch := range word {
				if ch < ' ' || ch > '~' {
					continue // Skip non-printable characters
				}
				start := ((int(ch) - 32) * 9) + i

				// Escape the ASCII characters so they render safely in HTML
				escapedArt := html.EscapeString(banner[start])

				// If the character is part of the substring, wrap it in a colored span
				if colorMap[charIdx] {
					strBuild.WriteString(fmt.Sprintf("<span style=\"color: %s;\">%s</span>", color, escapedArt))
				} else {
					strBuild.WriteString(escapedArt)
				}
			}
			strBuild.WriteByte('\n')
		}
	}
	return strBuild.String()
}