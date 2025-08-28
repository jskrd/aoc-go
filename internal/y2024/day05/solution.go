package day05

import (
	"slices"
	"strconv"
	"strings"
)

type parsedInput struct {
	orderings [][2]int
	updates   [][]int
}

func parseInput(input string) parsedInput {
	sections := strings.Split(input, "\n\n")

	orderings := [][2]int{}
	for _, line := range strings.Split(sections[0], "\n") {
		numbers := strings.Split(line, "|")
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])
		orderings = append(orderings, [2]int{a, b})
	}

	pages := [][]int{}
	for _, line := range strings.Split(sections[1], "\n") {
		numbers := strings.Split(line, ",")
		page := []int{}
		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			page = append(page, num)
		}
		pages = append(pages, page)
	}

	return parsedInput{orderings, pages}
}

func SolvePartOne(input string) int {
	parsedInput := parseInput(input)

	sum := 0
	for _, update := range parsedInput.updates {
		if isCorrectOrdering(parsedInput.orderings, update) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func SolvePartTwo(input string) int {
	parsedInput := parseInput(input)

	sum := 0
	for _, update := range parsedInput.updates {
		if !isCorrectOrdering(parsedInput.orderings, update) {
			correctedUpdate := correctUpdate(parsedInput.orderings, update)
			sum += correctedUpdate[len(correctedUpdate)/2]
		}
	}
	return sum
}

func isCorrectOrdering(orderings [][2]int, update []int) bool {
	for i, page := range update {
		before := findOrderingsBefore(orderings, page)
		for _, b := range before {
			if slices.Index(update[i+1:], b) != -1 {
				return false
			}
		}

		after := findOrderingsAfter(orderings, page)
		for _, a := range after {
			if slices.Index(update[:i], a) != -1 {
				return false
			}
		}
	}

	return true
}

func findOrderingsBefore(orderings [][2]int, page int) []int {
	before := []int{}
	for _, ordering := range orderings {
		if ordering[1] == page {
			before = append(before, ordering[0])
		}
	}
	return before
}

func findOrderingsAfter(orderings [][2]int, page int) []int {
	after := []int{}
	for _, ordering := range orderings {
		if ordering[0] == page {
			after = append(after, ordering[1])
		}
	}
	return after
}

func correctUpdate(orderings [][2]int, update []int) []int {
	corrected := []int{}
	for index, page := range update {
		insertPosition := index

		dependentPages := findOrderingsAfter(orderings, page)
		for _, a := range dependentPages {
			existingPosition := slices.Index(corrected, a)
			if existingPosition != -1 && existingPosition < insertPosition {
				insertPosition = existingPosition
			}
		}

		corrected = append(
			corrected[:insertPosition],
			append([]int{page}, corrected[insertPosition:]...)...,
		)
	}
	return corrected
}
