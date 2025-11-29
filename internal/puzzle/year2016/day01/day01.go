package day01

import (
	"strconv"
	"strings"
)

type instruction struct {
	Direction byte // 'L' or 'R'
	Steps     int
}

type point struct {
	X, Y int
}

func SolveLevel1(input string) int {
	instructions := parseInput(input)

	x, y := takeSteps(instructions, false)
	return numOfBlocksAway(x, y)
}

func SolveLevel2(input string) int {
	instructions := parseInput(input)

	x, y := takeSteps(instructions, true)
	return numOfBlocksAway(x, y)
}

func parseInput(input string) []instruction {
	parts := strings.Split(
		strings.Trim(input, "\n"),
		", ",
	)

	parsed := make([]instruction, len(parts))
	for index, part := range parts {
		direction := part[0]
		if direction != 'L' && direction != 'R' {
			panic("Invalid direction")
		}

		steps, err := strconv.Atoi(part[1:])
		if err != nil {
			panic(err)
		}

		parsed[index] = instruction{direction, steps}
	}

	return parsed
}

func takeSteps(instructions []instruction, trackVisited bool) (int, int) {
	var x, y int = 0, 0
	var facing byte = 'N'

	var visited = map[point]bool{{0, 0}: true}

	for _, instruction := range instructions {
		if isTurningNorth(facing, instruction.Direction) {
			facing = 'N'
			for range instruction.Steps {
				y += 1
				if trackVisited && hasAlreadyVisited(&visited, x, y) {
					return x, y
				}
			}
			continue
		}

		if isTurningEast(facing, instruction.Direction) {
			facing = 'E'
			for range instruction.Steps {
				x += 1
				if trackVisited && hasAlreadyVisited(&visited, x, y) {
					return x, y
				}
			}
			continue
		}

		if isTurningSouth(facing, instruction.Direction) {
			facing = 'S'
			for range instruction.Steps {
				y -= 1
				if trackVisited && hasAlreadyVisited(&visited, x, y) {
					return x, y
				}
			}
			continue
		}

		if isTurningWest(facing, instruction.Direction) {
			facing = 'W'
			for range instruction.Steps {
				x -= 1
				if trackVisited && hasAlreadyVisited(&visited, x, y) {
					return x, y
				}
			}
			continue
		}

		panic("Invalid direction")
	}

	return x, y
}

func isTurningNorth(facing byte, direction byte) bool {
	return (facing == 'E' && direction == 'L') ||
		(facing == 'W' && direction == 'R')
}

func isTurningEast(facing byte, direction byte) bool {
	return (facing == 'N' && direction == 'R') ||
		(facing == 'S' && direction == 'L')
}

func isTurningSouth(facing byte, direction byte) bool {
	return (facing == 'E' && direction == 'R') ||
		(facing == 'W' && direction == 'L')
}

func isTurningWest(facing byte, direction byte) bool {
	return (facing == 'N' && direction == 'L') ||
		(facing == 'S' && direction == 'R')
}

func hasAlreadyVisited(visited *map[point]bool, x, y int) bool {
	if (*visited)[point{x, y}] {
		return true
	}
	(*visited)[point{x, y}] = true
	return false
}

func numOfBlocksAway(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}
