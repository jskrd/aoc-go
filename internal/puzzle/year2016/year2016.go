package year2016

import (
	"github.com/jskrd/aoc-go/internal/puzzle/year2016/day01"
)

var Solutions = map[int]map[int]func(input string) int{
	1: {
		1: day01.SolveLevel1,
		2: day01.SolveLevel2,
	},
}
