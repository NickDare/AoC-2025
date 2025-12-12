package main

import (
	"fmt"
	"strings"

	"github.com/NickDare/AoC-2025/utils"
)

type DeviceName string

type Device struct {
	name    DeviceName
	outputs []DeviceName
}

func partA(input []string) int {
	devices := make(map[string]*Device)

	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], ":")
		deviceName := parts[0]
		outputList := strings.Fields(parts[1])

		device := &Device{
			name:    DeviceName(deviceName),
			outputs: make([]DeviceName, 0),
		}
		for _, out := range outputList {
			device.outputs = append(device.outputs, DeviceName(out))
		}
		devices[deviceName] = device
	}

	totalPaths := 0
	for _, output := range devices["you"].outputs {
		if output == "out" {
			totalPaths++
		} else {
			totalPaths += countPaths(devices, output, "out")
		}

	}
	return totalPaths
}

func countPaths(devices map[string]*Device, current DeviceName, target DeviceName) int {
	if current == target {
		return 1
	}
	totalPaths := 0
	for _, output := range devices[string(current)].outputs {
		totalPaths += countPaths(devices, output, target)
	}
	return totalPaths
}

func countPathsWithDacAndfft(devices map[string]*Device, current DeviceName, target DeviceName, dacUsed bool, fftUsed bool, cachePaths map[string]int) int {
	if current == target {
		if dacUsed && fftUsed {
			return 1
		}
		return 0
	}

	cacheKey := fmt.Sprintf("%s|%t|%t", current, dacUsed, fftUsed)
	if val, ok := cachePaths[cacheKey]; ok {
		return val
	}

	totalPaths := 0
	for _, output := range devices[string(current)].outputs {
		newDacUsed := dacUsed
		newFftUsed := fftUsed
		if output == "dac" {
			newDacUsed = true
		}
		if output == "fft" {
			newFftUsed = true
		}
		totalPaths += countPathsWithDacAndfft(devices, output, target, newDacUsed, newFftUsed, cachePaths)
	}
	cachePaths[cacheKey] = totalPaths

	return totalPaths
}

func partB(input []string) int {
	devices := make(map[string]*Device)

	for i := 0; i < len(input); i++ {
		parts := strings.Split(input[i], ":")
		deviceName := parts[0]
		outputList := strings.Fields(parts[1])

		device := &Device{
			name:    DeviceName(deviceName),
			outputs: make([]DeviceName, 0),
		}
		for _, out := range outputList {
			device.outputs = append(device.outputs, DeviceName(out))
		}
		devices[deviceName] = device
	}

	cachePaths := make(map[string]int)
	totalPaths := countPathsWithDacAndfft(devices, "svr", "out", false, false, cachePaths)
	return totalPaths
}

func main() {

	eInput := utils.ReadInput("eInput.txt")
	eInput2 := utils.ReadInput("eInput2.txt")
	myInput := utils.ReadInput("myInput.txt")

	partAResult := partA(eInput)
	fmt.Println("Part A Example:", partAResult)
	partAResult = partA(myInput)
	fmt.Println("Part A My Input:", partAResult)
	partBResult := partB(eInput2)
	fmt.Println("Part B Example:", partBResult)
	partBResult = partB(myInput)
	fmt.Println("Part B My Input:", partBResult)

}
