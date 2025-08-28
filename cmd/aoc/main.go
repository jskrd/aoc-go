package main

import (
	"flag"
	"os"

	"github.com/jskrd/aoc-go/internal/day01"
	"github.com/jskrd/aoc-go/internal/day02"
	"github.com/jskrd/aoc-go/internal/day03"
	"github.com/jskrd/aoc-go/internal/day04"
	"github.com/jskrd/aoc-go/internal/day05"
	"github.com/jskrd/aoc-go/internal/day06"
)

func main() {
	day := flag.Int("day", 0, "The day of the puzzle to solve")
	part := flag.Int("part", 0, "The part of the puzzle to solve")
	input := flag.String("input", "", "The path to the input file")

	flag.Parse()

	answer := solve(*day, *part, *input)
	println(answer)
}

func solve(day int, part int, input string) int {
	puzzles := [25][2]func(string) int{
		{day01.SolvePartOne, day01.SolvePartTwo},
		{day02.SolvePartOne, day02.SolvePartTwo},
		{day03.SolvePartOne, day03.SolvePartTwo},
		{day04.SolvePartOne, day04.SolvePartTwo},
		{day05.SolvePartOne, day05.SolvePartTwo},
		{day06.SolvePartOne, day06.SolvePartTwo},
	}

	if day < 1 || day > len(puzzles) {
		panic("A solution for this day does not exist")
	}

	if part < 1 || part > len(puzzles[day]) {
		panic("A solution for this part does not exist")
	}

	contents, error := os.ReadFile(input)
	if error != nil {
		panic(error)
	}

	return puzzles[day-1][part-1](string(contents))
}
