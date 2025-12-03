package year2025

import (
	"github.com/jskrd/aoc-go/internal/puzzle/year2025/day01"
	"github.com/jskrd/aoc-go/internal/puzzle/year2025/day02"
)

var Solutions = map[int]map[int]func(input string) int{
	1: {
		1: day01.SolveLevel1,
		2: day01.SolveLevel2,
	},
	2: {
		1: day02.SolveLevel1,
		2: day02.SolveLevel2,
	},
}
