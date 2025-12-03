package day02

import (
	"strconv"
	"strings"
)

type idRange struct {
	Start int
	End   int
}

func SolveLevel1(input string) int {
	idRanges := parseInput(input)

	var invalidIds []int
	for _, idRange := range idRanges {
		for id := idRange.Start; id <= idRange.End; id++ {
			if isRepeatedOnlyTwice(id) {
				invalidIds = append(invalidIds, id)
			}
		}
	}

	var sum int = 0
	for _, id := range invalidIds {
		sum += id
	}

	return sum
}

func SolveLevel2(input string) int {
	idRanges := parseInput(input)

	var invalidIds []int
	for _, idRange := range idRanges {
		for id := idRange.Start; id <= idRange.End; id++ {
			if isRepeatedAtLeastTwice(id) {
				invalidIds = append(invalidIds, id)
			}
		}
	}

	var sum int = 0
	for _, id := range invalidIds {
		sum += id
	}

	return sum
}

func parseInput(input string) []idRange {
	parts := strings.Split(
		strings.Trim(input, "\n"),
		",",
	)

	parsed := make([]idRange, len(parts))
	for index, part := range parts {
		split := strings.Split(part, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		parsed[index] = idRange{Start: start, End: end}
	}

	return parsed
}

func isRepeatedOnlyTwice(id int) bool {
	digits := strconv.Itoa(id)
	digitsCount := len(digits)

	if digitsCount%2 != 0 {
		return false
	}

	firstHalf := digits[:digitsCount/2]
	secondHalf := digits[digitsCount/2:]

	return firstHalf == secondHalf
}

func isRepeatedAtLeastTwice(id int) bool {
	digits := strconv.Itoa(id)
	digitsCount := len(digits)

	for division := 1; division <= digitsCount/2; division++ {
		if digitsCount%division != 0 {
			continue
		}

		part := digits[:division]
		if strings.Count(digits, part) == digitsCount/division {
			return true
		}
	}

	return false
}
