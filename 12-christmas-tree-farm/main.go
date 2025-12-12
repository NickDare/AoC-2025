package main

import (
	"strconv"
	"strings"

	"github.com/NickDare/AoC-2025/utils"
)

type present struct {
	id   int
	size int
}

type puzzle struct {
	gridSize      int
	numOfPresents []int
}

func parsePresents(input []string) []present {
	current := present{-1, 0}
	presents := []present{}
	for _, line := range input {
		if line == "" || strings.Contains(line, "x") {
			continue
		}
		if strings.Contains(line, ":") {
			if current.id != -1 && current.size != 0 {
				presents = append(presents, current)
			}
			parts := strings.Split(line, ":")
			id, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			current = present{id, 0}
		}
		if strings.Contains(line, "#") {
			count := strings.Count(line, "#")
			current.size += count
		}
	}
	if current.id != -1 && current.size != 0 {
		presents = append(presents, current)
	}
	return presents
}

func parsePuzzles(input []string) []puzzle {
	puzzles := []puzzle{}
	for _, line := range input {
		if !strings.Contains(line, "x") {
			continue
		}
		parts := strings.Split(line, ":")
		gridVals := strings.Split(parts[0], "x")
		width, _ := strconv.Atoi(strings.TrimSpace(gridVals[0]))
		height, _ := strconv.Atoi(strings.TrimSpace(gridVals[1]))
		gridSize := width * height
		requiredPresents := strings.Fields(parts[1])
		requiredPresentsInts := []int{}
		for _, requiredCount := range requiredPresents {
			count, _ := strconv.Atoi(requiredCount)
			requiredPresentsInts = append(requiredPresentsInts, count)
		}
		puzzles = append(puzzles, puzzle{gridSize, requiredPresentsInts})
	}
	return puzzles
}

func partA(input []string) int {
	presents := parsePresents(input)
	puzzles := parsePuzzles(input)

	canFitCount := 0
	for _, p := range puzzles {
		spaceRequired := 0
		for i := 0; i < len(p.numOfPresents); i++ {
			spaceRequired += p.numOfPresents[i] * presents[i].size
		}
		if spaceRequired <= p.gridSize {
			canFitCount++
		}
	}

	return canFitCount
}

func main() {
	input := utils.ReadInput("eInput.txt")
	myInput := utils.ReadInput("myInput.txt")
	res := partA(input)
	println("Part A:", res)
	res = partA(myInput)
	println("Part A My Input:", res)

}
