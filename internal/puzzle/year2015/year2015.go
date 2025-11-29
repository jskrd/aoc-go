package year2015

import (
	"github.com/jskrd/aoc-go/internal/puzzle/year2015/day01"
)

var Solutions = map[int]map[int]func(input string) int{
	1: {
		1: day01.SolveLevel1,
		2: day01.SolveLevel2,
	},
}
