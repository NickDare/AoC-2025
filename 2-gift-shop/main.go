package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	exampleIds = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	ourIds     = "655-1102,2949-4331,885300-1098691,1867-2844,20-43,4382100-4484893,781681037-781860439,647601-734894,2-16,180-238,195135887-195258082,47-64,4392-6414,6470-10044,345-600,5353503564-5353567532,124142-198665,1151882036-1151931750,6666551471-6666743820,207368-302426,5457772-5654349,72969293-73018196,71-109,46428150-46507525,15955-26536,65620-107801,1255-1813,427058-455196,333968-391876,482446-514820,45504-61820,36235767-36468253,23249929-23312800,5210718-5346163,648632326-648673051,116-173,752508-837824"
)

func isInvalidIdPart1(id string) bool {
	idLength := len(id)
	if idLength%2 != 0 {
		return false
	}

	halfLength := idLength / 2
	firstHalf := id[:halfLength]
	secondHalf := id[halfLength:]

	return firstHalf == secondHalf
}

func isInvalidIdPart2(id string) bool {
	idLength := len(id)

	for subLenth := 1; subLenth <= idLength/2; subLenth++ {
		if idLength%subLenth != 0 {
			continue
		}
		subString := id[:subLenth]
		repeats := idLength / subLenth
		repeatedString := strings.Repeat(subString, repeats)
		if repeatedString == id {
			return true
		}
	}
	return false
}

type Part string

const (
	Part1 Part = "A"
	Part2 Part = "B"
)

func checkRangeForInvalidIds(startId string, endId string, part Part) []string {
	invalidIds := []string{}

	startValue, _ := strconv.Atoi(startId)
	endValue, _ := strconv.Atoi(endId)
	for i := startValue; i <= endValue; i++ {
		idStr := strconv.Itoa(i)
		if part == Part1 && isInvalidIdPart1(idStr) {
			invalidIds = append(invalidIds, idStr)
		} else if part == Part2 &&
			isInvalidIdPart2(idStr) {
			invalidIds = append(invalidIds, idStr)
		}
	}

	return invalidIds
}

func part1(ids []string) {
	sumOfInvalidIds := 0

	for _, idRange := range ids {
		bounds := strings.Split(idRange, "-")
		startId := bounds[0]
		endId := bounds[1]
		invalidIds := checkRangeForInvalidIds(startId, endId, Part1)
		for _, invalidId := range invalidIds {
			invalidIdValue, _ := strconv.Atoi(invalidId)
			sumOfInvalidIds += invalidIdValue
		}
	}

	fmt.Printf("Part 1 - Sum of invalid Ids: %d\n", sumOfInvalidIds)
}

func part2(ids []string) {
	sumOfInvalidIds := 0

	for _, idRange := range ids {
		bounds := strings.Split(idRange, "-")
		startId := bounds[0]
		endId := bounds[1]
		invalidIds := checkRangeForInvalidIds(startId, endId, Part2)
		for _, invalidId := range invalidIds {
			invalidIdValue, _ := strconv.Atoi(invalidId)
			sumOfInvalidIds += invalidIdValue
		}
	}

	fmt.Printf("Part 2 - Sum of invalid Ids: %d\n", sumOfInvalidIds)
}

func main() {

	exampleIds := strings.Split(exampleIds, ",")
	fmt.Println("Example Ids:")
	part1(exampleIds)
	part2(exampleIds)
	// Example answers
	// Part 1 - Sum of invalid Ids: 1227775554
	// Part 2 - Sum of invalid Ids: 4174379265

	ids := strings.Split(ourIds, ",")
	fmt.Println("Our Ids:")
	part1(ids)
	part2(ids)
	// Correct answers our set
	// Part 1 - Sum of invalid Ids: 21898734247
	// Part 2 - Sum of invalid Ids: 28915664389
}
