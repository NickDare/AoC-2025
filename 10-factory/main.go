package main

import (
	"fmt"
	"math"
	"math/bits"
	"slices"
	"strings"

	"github.com/NickDare/AoC-2025/utils"
)

type Diagram []bool
type Schematics [][]int
type Joltages []int

func parseDiagram(input string) Diagram {
	a := strings.Split(input, "[")
	b := strings.Split(a[1], "]")

	Diagram := make([]bool, 0)

	for _, char := range b[0] {
		if string(char) == "#" {
			Diagram = append(Diagram, true)
		} else {
			Diagram = append(Diagram, false)
		}
	}

	return Diagram

}

func parseSchematics(input string) Schematics {
	a := strings.Split(input, "]")
	b := strings.Split(a[1], "{")

	schematics := make([][]int, 0)
	b[0] = strings.TrimSpace(b[0])

	dirtySchematics := strings.Split(b[0], "(")
	dirtySchematics = dirtySchematics[1:] // remove first empty element
	for _, ds := range dirtySchematics {
		if ds == "" {
			continue
		}
		cleaned := strings.TrimSuffix(ds, ")")
		numStrings := strings.Split(cleaned, ",")
		nums := make([]int, 0)
		for _, ns := range numStrings {
			var num int
			fmt.Sscanf(ns, "%d", &num)
			nums = append(nums, num)
		}
		schematics = append(schematics, nums)
	}

	return schematics
}

func parseJoltages(input string) Joltages {
	a := strings.Split(input, "{")
	b := strings.Split(a[1], "}")

	joltages := make([]int, 0)
	numStrings := strings.Split(b[0], ",")
	for _, ns := range numStrings {
		var num int
		fmt.Sscanf(ns, "%d", &num)
		joltages = append(joltages, num)
	}

	return joltages
}

func parseManualInstructions(input string) (Diagram, Schematics, Joltages) {
	return parseDiagram(input), parseSchematics(input), parseJoltages(input)
}

func calcLowestInputsNeeded(d Diagram, s Schematics, j Joltages) int {
	totalSchematics := len(s)
	lowestInputs := totalSchematics + 1

	possibleCombinations := int(math.Pow(2, float64(totalSchematics)))
	for schematicCombo := range possibleCombinations {
		currentState := make([]bool, len(d))

		for schematicIndex := range totalSchematics {
			schematicMask := int(math.Pow(2, float64(schematicIndex)))
			if schematicCombo&schematicMask != 0 {
				for _, switchPosition := range s[schematicIndex] {
					currentState[switchPosition] = !currentState[switchPosition]
				}
			}
		}

		if slices.Equal(currentState, d) {
			count := bits.OnesCount(uint(schematicCombo))
			if count < lowestInputs {
				lowestInputs = count
			}
		}
	}

	return lowestInputs
}

func partA(input []string) {
	lowestInputs := []int{}
	for i := range input {
		dia, schem, jolts := parseManualInstructions(input[i])
		lowest := calcLowestInputsNeeded(dia, schem, jolts)
		lowestInputs = append(lowestInputs, lowest)
	}

	sumOfLowest := 0
	for _, li := range lowestInputs {
		sumOfLowest += li
	}

	fmt.Println("Part A:", sumOfLowest)
}

func partB(input []string) {
}

func main() {
	input := utils.ReadInput("eInput.txt")
	myInput := utils.ReadInput("myInput.txt")
	partA(input)
	partA(myInput)
	partB(input)
}
