package main

import (
	"fmt"
	"os"
	"strings"
)

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

type operator string

const (
	add      operator = "+"
	multiply operator = "*"
)

func getOperator(col int, input []string) operator {
	line := input[len(input)-1]
	parts := strings.Fields(line)

	if parts[col] == "+" {
		return add
	} else {
		return multiply
	}

}

func cleanInput(input []string) [][]string {
	var cleaned [][]string

	for _, line := range input {
		data := strings.Split(line, " ")
		tempArray := []string{}
		for _, d := range data {
			if d != "" {
				tempArray = append(tempArray, d)
			}
		}

		for i, d := range tempArray {
			if len(cleaned) <= i {
				cleaned = append(cleaned, []string{})
			}
			cleaned[i] = append(cleaned[i], d)
		}
	}

	return cleaned

}

func partA(input []string) {

	clened := cleanInput(input)
	sumOfAllProblems := 0

	for i, line := range clened {
		op := getOperator(i, input)

		result := 0
		for j, item := range line[:len(line)-1] {
			var num int
			fmt.Sscanf(item, "%d", &num)

			if j == 0 {
				result = num
			} else {
				switch op {
				case add:
					result += num
				case multiply:
					result *= num
				}
			}
		}

		sumOfAllProblems += result
	}

	fmt.Println("Part A:", sumOfAllProblems)

}

func partB(input []string) {
	inputLines := input[:len(input)-1]
	maxLength := 0
	for _, line := range inputLines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	var problems [][]string
	var currentProblem []string

	for col := 0; col < maxLength; col++ {
		digits := ""
		for _, line := range inputLines {
			if col < len(line) && line[col] != ' ' {
				digits += string(line[col])
			}
		}

		if digits == "" {
			if len(currentProblem) > 0 {
				problems = append(problems, currentProblem)
				currentProblem = nil
			}
		} else {
			currentProblem = append(currentProblem, digits)
		}
	}
	if len(currentProblem) > 0 {
		problems = append(problems, currentProblem)
	}

	sumOfAllProblems := 0
	for i, problem := range problems {
		operator := getOperator(i, input)
		result := 0
		for j := len(problem) - 1; j >= 0; j-- {
			var num int
			fmt.Sscanf(problem[j], "%d", &num)

			if j == len(problem)-1 {
				result = num
			} else {
				switch operator {
				case "+":
					result += num
				case "*":
					result *= num
				default:
					fmt.Println("Unknown operator:", operator)
					return
				}
			}
		}

		sumOfAllProblems += result
	}

	fmt.Println("Part B:", sumOfAllProblems)
}

func main() {
	eInput := readInput("eInput.txt")
	myInput := readInput("myInput.txt")

	partA(eInput)
	partA(myInput)
	partB(eInput)
	partB(myInput)
}
