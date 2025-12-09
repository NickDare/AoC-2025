package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type BoxPair struct {
	dist int
	boxA JunctionBox
	boxB JunctionBox
}

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

type JunctionBox struct {
	x, y, z int
}

func parseJunctionBoxes(lines []string) []JunctionBox {
	boxes := make([]JunctionBox, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		boxes[i] = JunctionBox{x: x, y: y, z: z}
	}
	return boxes
}

func euclideanDistance(a, b JunctionBox) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

type Circuit struct {
	boxes []JunctionBox
}

func partA(lines []string, numPairs int) int {
	boxes := parseJunctionBoxes(lines)

	// Create a circuit for each box (use pointers for mutation)
	circuits := make([]*Circuit, len(boxes))
	boxToCircuit := make(map[JunctionBox]*Circuit)
	for i, box := range boxes {
		circuits[i] = &Circuit{boxes: []JunctionBox{box}}
		boxToCircuit[box] = circuits[i]
	}

	// Generate all pairs and sort by distance
	var pairs []BoxPair
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, BoxPair{
				dist: euclideanDistance(boxes[i], boxes[j]),
				boxA: boxes[i],
				boxB: boxes[j],
			})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})

	// Process exactly numPairs pairs (not successful connections!)
	for i := 0; i < numPairs && i < len(pairs); i++ {
		pair := pairs[i]
		circuitA := boxToCircuit[pair.boxA]
		circuitB := boxToCircuit[pair.boxB]

		// Only merge if different circuits
		if circuitA != circuitB {
			circuitA.boxes = append(circuitA.boxes, circuitB.boxes...)
			for _, box := range circuitB.boxes {
				boxToCircuit[box] = circuitA
			}
			circuitB.boxes = nil
		}
	}

	// Collect circuit sizes
	var sizes []int
	for _, c := range circuits {
		if len(c.boxes) > 0 {
			sizes = append(sizes, len(c.boxes))
		}
	}

	// Sort descending and return product of 3 largest
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func partB(lines []string) int {
	boxes := parseJunctionBoxes(lines)

	// Create a circuit for each box
	circuits := make([]*Circuit, len(boxes))
	boxToCircuit := make(map[JunctionBox]*Circuit)
	for i, box := range boxes {
		circuits[i] = &Circuit{boxes: []JunctionBox{box}}
		boxToCircuit[box] = circuits[i]
	}

	// Generate all pairs and sort by distance
	var pairs []BoxPair
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, BoxPair{
				dist: euclideanDistance(boxes[i], boxes[j]),
				boxA: boxes[i],
				boxB: boxes[j],
			})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})

	numCircuits := len(boxes)
	var lastPair BoxPair

	// Process pairs until all boxes are in one circuit
	for _, pair := range pairs {
		if numCircuits == 1 {
			break
		}

		circuitA := boxToCircuit[pair.boxA]
		circuitB := boxToCircuit[pair.boxB]

		// Only merge if different circuits
		if circuitA != circuitB {
			lastPair = pair
			circuitA.boxes = append(circuitA.boxes, circuitB.boxes...)
			for _, box := range circuitB.boxes {
				boxToCircuit[box] = circuitA
			}
			circuitB.boxes = nil
			numCircuits--
		}
	}

	return lastPair.boxA.x * lastPair.boxB.x
}

func main() {
	eInput := readInput("eInput.txt")
	myInput := readInput("myInput.txt")

	eRes := partA(eInput, 10)
	println("Example Part A:", eRes)
	myRes := partA(myInput, 1000)
	println("My Part A:", myRes)

	eResB := partB(eInput)
	println("Example Part B:", eResB)
	myResB := partB(myInput)
	println("My Part B:", myResB)
}
