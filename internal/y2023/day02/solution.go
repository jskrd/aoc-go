package day02

import (
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	data, err := os.ReadFile("internal/day02/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	part1Result := make(chan int)
	part2Result := make(chan int)

	go func() {
		part1Result <- SolvePart1(&lines)
	}()
	go func() {
		part2Result <- SolvePart2(&lines)
	}()

	return <-part1Result, <-part2Result
}

func SolvePart1(lines *[]string) int {
	sum := 0
	for _, line := range *lines {
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

func SolvePart2(lines *[]string) int {
	sum := 0
	for _, line := range *lines {
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

// Wrappers to match the codebase's Solver signature.
func SolvePartOne(input string) int {
	lines := strings.Split(input, "\n")
	return SolvePart1(&lines)
}

func SolvePartTwo(input string) int {
	lines := strings.Split(input, "\n")
	return SolvePart2(&lines)
}
