package utils

import "os"

func ReadInput(path string) []string {
	file, _ := os.ReadFile(path)
	return splitLines(string(file))

}

func splitLines(s string) []string {
	var lines []string
	currentLine := ""
	for _, char := range s {
		if char == '\n' {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLine += string(char)
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}
