package main

import "os"

func readInput(path string) []string {

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

func partA(lines []string) int {

	startIndex := -1
	for i, line := range lines {
		for j, char := range line {
			if char == 'S' {
				lineBytes := []byte(lines[i+1])
				lineBytes[j] = '|'
				lines[i+1] = string(lineBytes)
				startIndex = i + 1
				break
			}

			if startIndex != -1 {
				break
			}
		}
		if startIndex != -1 {
			break
		}
	}

	totalSplits := 0
	for i := startIndex + 1; i < len(lines); i++ {
		line := lines[i]
		prevLine := lines[i-1]

		lineBytes := []byte(line)
		for j, char := range line {
			if char == '^' && prevLine[j] == '|' {
				if j-1 >= 0 {
					lineBytes[j-1] = '|'
				}
				if j+1 < len(line) {
					lineBytes[j+1] = '|'
				}
				totalSplits++
			}
		}
		lines[i] = string(lineBytes)

		for j, char := range prevLine {
			if char == '|' && lines[i][j] != '^' {
				lineBytes := []byte(lines[i])
				lineBytes[j] = '|'
				lines[i] = string(lineBytes)
			}
		}
	}

	return totalSplits
}

func partB(lines []string) int {

	startIndex := -1
	startCol := -1
	for i, line := range lines {
		for j, char := range line {
			if char == 'S' {
				lineBytes := []byte(lines[i+1])
				lineBytes[j] = '|'
				lines[i+1] = string(lineBytes)
				startIndex = i + 1
				startCol = j
				break
			}

			if startIndex != -1 {
				break
			}
		}
		if startIndex != -1 {
			break
		}
	}

	combinationMap := make(map[int]int)
	combinationMap[startCol] = 1

	for i := startIndex; i < len(lines); i++ {
		line := lines[i]

		newCombinationMap := make(map[int]int)

		for j, count := range combinationMap {
			switch line[j] {
			case '|':
				if i+1 < len(lines) {
					newCombinationMap[j] += count
				} else {
					newCombinationMap[j] += count
				}
			case '^':
				if j-1 >= 0 {
					newCombinationMap[j-1] += count
				}
				if j+1 < len(line) {
					newCombinationMap[j+1] += count
				}
			}
		}

		combinationMap = newCombinationMap
	}

	totalCombinations := 0
	for _, count := range combinationMap {
		totalCombinations += count
	}

	return totalCombinations
}

func main() {
	eInput := readInput("eInput.txt")
	myInput := readInput("myInput.txt")

	eRes := partA(eInput)
	println("Example Part A:", eRes)
	myRes := partA(myInput)
	println("My Part A:", myRes)

	eResB := partB(eInput)
	println("Example Part B:", eResB)
	myResB := partB(myInput)
	println("My Part B:", myResB)

}
