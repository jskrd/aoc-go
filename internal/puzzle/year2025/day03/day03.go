package day03

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveLevel1(input string) int {
	banks := parseInput(input)

	var sum = 0
	for _, bank := range banks {
		sum += findLargestJoltage(bank)
	}

	return sum
}

func parseInput(input string) []string {
	return strings.Split(
		strings.Trim(input, "\n"),
		"\n",
	)
}

func findLargestJoltage(bank string) int {
	batteries := make([]int, len(bank))
	for i, battery := range bank {
		batteries[i] = int(battery - '0')
	}

	firstSet := batteries[:len(batteries)-1]
	var firstLargest int
	for i, battery := range firstSet {
		if battery > firstSet[firstLargest] {
			firstLargest = i
		}
	}

	secondSet := batteries[firstLargest+1:]
	var secondLargest int
	for i, battery := range secondSet {
		if battery > secondSet[secondLargest] {
			secondLargest = i
		}
	}

	largestJoltage, _ := strconv.Atoi(
		fmt.Sprintf("%d%d", firstSet[firstLargest], secondSet[secondLargest]),
	)
	return largestJoltage
}
