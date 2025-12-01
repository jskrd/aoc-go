package day01

import (
	"strconv"
	"strings"
)

type rotation struct {
	Direction byte // 'L' or 'R'
	Count     int
}

func SolveLevel1(input string) int {
	rotations := parseInput(input)

	var pointsAtZeroCount = 0
	var dial = 50

	for _, rotation := range rotations {
		if rotation.Direction == 'L' {
			dial -= rotation.Count
		} else if rotation.Direction == 'R' {
			dial += rotation.Count
		} else {
			panic("Invalid direction")
		}

		if dial%100 == 0 {
			pointsAtZeroCount++
		}
	}

	return pointsAtZeroCount
}

func SolveLevel2(input string) int {
	rotations := parseInput(input)

	var pointsAtZeroCount = 0
	var dial = 50

	for _, rotation := range rotations {
		var increment int

		switch rotation.Direction {
		case 'L':
			increment = -1
		case 'R':
			increment = 1
		default:
			panic("Invalid direction")
		}

		for range rotation.Count {
			dial += increment

			if dial%100 == 0 {
				pointsAtZeroCount++
			}
		}
	}

	return pointsAtZeroCount
}

func parseInput(input string) []rotation {
	parts := strings.Split(
		strings.Trim(input, "\n"),
		"\n",
	)

	parsed := make([]rotation, len(parts))
	for index, part := range parts {
		direction := part[0]
		if direction != 'L' && direction != 'R' {
			panic("Invalid direction")
		}

		steps, err := strconv.Atoi(part[1:])
		if err != nil {
			panic(err)
		}

		parsed[index] = rotation{direction, steps}
	}

	return parsed
}
