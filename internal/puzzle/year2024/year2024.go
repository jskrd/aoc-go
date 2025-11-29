package year2024

import (
	"github.com/jskrd/aoc-go/internal/puzzle/year2024/day01"
	"github.com/jskrd/aoc-go/internal/puzzle/year2024/day02"
	"github.com/jskrd/aoc-go/internal/puzzle/year2024/day03"
	"github.com/jskrd/aoc-go/internal/puzzle/year2024/day04"
	"github.com/jskrd/aoc-go/internal/puzzle/year2024/day05"
	"github.com/jskrd/aoc-go/internal/puzzle/year2024/day06"
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
	3: {
		1: day03.SolveLevel1,
		2: day03.SolveLevel2,
	},
	4: {
		1: day04.SolveLevel1,
		2: day04.SolveLevel2,
	},
	5: {
		1: day05.SolveLevel1,
		2: day05.SolveLevel2,
	},
	6: {
		1: day06.SolveLevel1,
		2: day06.SolveLevel2,
	},
}
