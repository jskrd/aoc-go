package day02

import (
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	reports := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		report := []int{}
		for _, num := range strings.Split(line, " ") {
			level, _ := strconv.Atoi(num)
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	return reports
}

func SolvePartOne(input string) int {
	reports := parseInput(input)

	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func SolvePartTwo(input string) int {
	reports := parseInput(input)

	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
			continue
		}

		possibleReports := getAllPossibleReportsRemovingOne(report)
		for _, possibleReport := range possibleReports {
			if isSafe(possibleReport) {
				count++
				break
			}
		}
	}
	return count
}

func isSafe(report []int) bool {
	first := report[0]
	last := report[len(report)-1]

	if first < last {
		for i, level := range report {
			if i == 0 {
				continue
			}
			if level < report[i-1]+1 || level > report[i-1]+3 {
				return false
			}
		}
	} else {
		for i, level := range report {
			if i == 0 {
				continue
			}
			if level < report[i-1]-3 || level > report[i-1]-1 {
				return false
			}
		}
	}

	return true
}

func getAllPossibleReportsRemovingOne(report []int) [][]int {
	if len(report) <= 1 {
		return [][]int{}
	}

	reports := make([][]int, len(report))
	for i := 0; i < len(report); i++ {
		levels := make([]int, 0, len(report)-1)
		for j := 0; j < len(report); j++ {
			if j != i {
				levels = append(levels, report[j])
			}
		}
		reports[i] = levels
	}

	return reports
}
