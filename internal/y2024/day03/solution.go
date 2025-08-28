package day03

import (
	"regexp"
	"strconv"
)

func parseInput(input string) []string {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	return re.FindAllString(input, -1)
}

func SolvePartOne(input string) int {
	instructions := parseInput(input)

	sum := 0
	for _, instruction := range instructions {
		if instruction[:3] == "mul" {
			re := regexp.MustCompile(`([0-9]+),([0-9]+)`)
			matches := re.FindStringSubmatch(instruction)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			sum += a * b
		}
	}
	return sum
}

func SolvePartTwo(input string) int {
	instructions := parseInput(input)

	enabled := true
	sum := 0
	for _, instruction := range instructions {
		if instruction == "do()" {
			enabled = true
		}
		if instruction == "don't()" {
			enabled = false
		}
		if instruction[:3] == "mul" {
			if !enabled {
				continue
			}

			re := regexp.MustCompile(`([0-9]+),([0-9]+)`)
			matches := re.FindStringSubmatch(instruction)
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[2])
			sum += a * b
		}
	}
	return sum
}
