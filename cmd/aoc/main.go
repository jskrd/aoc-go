package main

import (
	"flag"
	"os"

	"fmt"
	"path/filepath"

	"github.com/jskrd/aoc-go/internal/years"
)

func main() {
	year := flag.Int("year", 0, "The year of the puzzles to solve")
	day := flag.Int("day", 0, "The day of the puzzle to solve")
	part := flag.Int("part", 0, "The part of the puzzle to solve")
	input := flag.String("input", "", "The path to the input file")

	flag.Parse()

	answer := solve(*year, *day, *part, *input)
	println(answer)
}

func solve(year int, day int, part int, input string) int {
	yr, ok := years.Years[year]
	if !ok {
		panic("A solution for this year does not exist")
	}

	if day < 1 || day > len(yr) {
		panic("A solution for this day does not exist")
	}

	if part < 1 || part > len(yr[day-1]) {
		panic("A solution for this part does not exist")
	}

	inputPath := input
	if inputPath == "" {
		inputPath = filepath.Join("test", "data", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day), "input.txt")
	}

	contents, error := os.ReadFile(inputPath)
	if error != nil {
		panic(error)
	}

	return yr[day-1][part-1](string(contents))
}
