package main

import (
	"fmt"

	"github.com/NickDare/AoC-2025/utils"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseCoordinates(line string) (int, int) {
	var x, y int
	fmt.Sscanf(line, "%d,%d", &x, &y)
	return x, y
}

func partA(input []string) {
	tileCords := make([][2]int, 0)
	for line := range input {
		x, y := parseCoordinates(input[line])
		tileCords = append(tileCords, [2]int{x, y})
	}

	grid := make(map[int]map[int]bool)
	for _, cord := range tileCords {
		x, y := cord[0], cord[1]
		if _, exists := grid[y]; !exists {
			grid[y] = make(map[int]bool)
		}
		grid[y][x] = true
	}

	// // print the grid
	// gridXSize := 0
	// gridYSize := 0
	// for _, cord := range tileCords {
	// 	x, y := cord[0], cord[1]
	// 	if x > gridXSize {
	// 		gridXSize = x
	// 	}
	// 	if y > gridYSize {
	// 		gridYSize = y
	// 	}
	// }

	// for y := 0; y < gridYSize+1; y++ {
	// 	for x := 0; x < gridXSize+1; x++ {
	// 		if grid[y][x] {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	maxArea := 0
	for i := 0; i < len(tileCords); i++ {
		for j := i + 1; j < len(tileCords); j++ {
			x1, y1 := tileCords[i][0], tileCords[i][1]
			x2, y2 := tileCords[j][0], tileCords[j][1]

			width := abs(x2-x1) + 1
			height := abs(y2-y1) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Println("Max area:", maxArea)
}

func partB(input []string) {
	fmt.Println(input)
}

func main() {
	eInput := utils.ReadInput("eInput.txt")
	myInput := utils.ReadInput("myInput.txt")
	partA(eInput)
	partA(myInput)
	// partB(eInput)
	// partB(myInput)
}
