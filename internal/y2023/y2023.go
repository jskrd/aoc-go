package y2023

import (
	"github.com/jskrd/aoc-go/internal/aoc"
	"github.com/jskrd/aoc-go/internal/y2023/day01"
	"github.com/jskrd/aoc-go/internal/y2023/day02"
)

var Days = [25]aoc.DayParts{
	{day01.SolvePartOne, day01.SolvePartTwo},
	{day02.SolvePartOne, day02.SolvePartTwo},
}
