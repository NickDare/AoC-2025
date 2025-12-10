package main

import (
	"fmt"
	"time"

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

type RowRanges struct {
	minX, maxX int
}

type FilledRows map[int]RowRanges

func (f FilledRows) createRanges(y, x1, x2 int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if ranges, exists := f[y]; exists {
		if x1 < ranges.minX {
			ranges.minX = x1
		}
		if x2 > ranges.maxX {
			ranges.maxX = x2
		}
		f[y] = ranges
	} else {
		f[y] = RowRanges{x1, x2}
	}
}

func (f FilledRows) isValidRectangle(minX, maxX, minY, maxY int) bool {
	for y := minY; y <= maxY; y++ {
		bounds, exists := f[y]
		if !exists || bounds.minX > minX || bounds.maxX < maxX {
			return false
		}
	}
	return true
}

func partB(input []string) {
	startTotal := time.Now()
	tileCords := make([][2]int, 0)
	for line := range input {
		x, y := parseCoordinates(input[line])
		tileCords = append(tileCords, [2]int{x, y})
	}

	filledRows := make(FilledRows)
	for i := 0; i < len(tileCords); i++ {
		for j := i + 1; j < len(tileCords); j++ {
			x1, y1 := tileCords[i][0], tileCords[i][1]
			x2, y2 := tileCords[j][0], tileCords[j][1]

			if x1 == x2 {
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				for y := y1; y <= y2; y++ {
					filledRows.createRanges(y, x1, x1)
				}
			} else if y1 == y2 {
				filledRows.createRanges(y1, x1, x2)
			}
		}
	}

	maxArea := 0
	for i := 0; i < len(tileCords); i++ {
		for j := i + 1; j < len(tileCords); j++ {
			x1, y1 := tileCords[i][0], tileCords[i][1]
			x2, y2 := tileCords[j][0], tileCords[j][1]

			minX, maxX := min(x1, x2), max(x1, x2)
			minY, maxY := min(y1, y2), max(y1, y2)
			area := (maxX - minX + 1) * (maxY - minY + 1)

			if area > maxArea && filledRows.isValidRectangle(minX, maxX, minY, maxY) {
				maxArea = area
			}
		}
	}
	fmt.Printf("Max area: %d (took %v)\n", maxArea, time.Since(startTotal))
}

func main() {
	eInput := utils.ReadInput("eInput.txt")
	myInput := utils.ReadInput("myInput.txt")
	partA(eInput)
	partA(myInput)
	partB(eInput)
	partB(myInput)
}
