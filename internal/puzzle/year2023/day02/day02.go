package day02

import (
	"strconv"
	"strings"
)

func SolveLevel1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		info := strings.Split(line, ": ")
		game, rounds := info[0], info[1]

		id, _ := strconv.Atoi(strings.Split(game, " ")[1])

		isPossible := true
		for _, round := range strings.Split(rounds, "; ") {
			if isPossible == false {
				break
			}

			for _, group := range strings.Split(round, ", ") {
				cube := strings.Split(group, " ")
				num, color := cube[0], cube[1]

				numI, _ := strconv.Atoi(num)

				if color == "red" && numI > 12 {
					isPossible = false
				} else if color == "green" && numI > 13 {
					isPossible = false
				} else if color == "blue" && numI > 14 {
					isPossible = false
				}
			}
		}

		if isPossible {
			sum += id
		}
	}

	return sum
}

func SolveLevel2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		info := strings.Split(line, ": ")

		var red, green, blue int
		for _, round := range strings.Split(info[1], "; ") {
			for _, group := range strings.Split(round, ", ") {
				cube := strings.Split(group, " ")
				num, color := cube[0], cube[1]

				numI, _ := strconv.Atoi(num)

				if color == "red" && numI > red {
					red = numI
				} else if color == "green" && numI > green {
					green = numI
				} else if color == "blue" && numI > blue {
					blue = numI
				}
			}
		}

		sum += red * green * blue
	}

	return sum
}
