package day01

import (
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	lists := make([][]int, 2)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "   ")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		lists[0] = append(lists[0], left)
		lists[1] = append(lists[1], right)
	}

	return lists
}

func SolvePartOne(input string) int {
	lists := parseInput(input)
	leftList, rightList := lists[0], lists[1]

	sort.Ints(leftList)
	sort.Ints(rightList)

	distances := []int{}
	for i, left := range leftList {
		right := rightList[i]
		distance := 0
		if left > right {
			distance = left - right
		} else {
			distance = right - left
		}
		distances = append(distances, distance)
	}

	sum := 0
	for _, distance := range distances {
		sum += distance
	}
	return sum
}

func SolvePartTwo(input string) int {
	lists := parseInput(input)
	leftList, rightList := lists[0], lists[1]

	similarities := []int{}
	for _, left := range leftList {
		count := 0
		for _, right := range rightList {
			if left == right {
				count++
			}
		}
		similarities = append(similarities, left*count)
	}

	sum := 0
	for _, similarity := range similarities {
		sum += similarity
	}
	return sum
}
