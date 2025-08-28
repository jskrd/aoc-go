package y2024

import (
	"github.com/jskrd/aoc-go/internal/aoc"
	"github.com/jskrd/aoc-go/internal/y2024/day01"
	"github.com/jskrd/aoc-go/internal/y2024/day02"
	"github.com/jskrd/aoc-go/internal/y2024/day03"
	"github.com/jskrd/aoc-go/internal/y2024/day04"
	"github.com/jskrd/aoc-go/internal/y2024/day05"
	"github.com/jskrd/aoc-go/internal/y2024/day06"
)

var Days = [25]aoc.DayParts{
	{day01.SolvePartOne, day01.SolvePartTwo},
	{day02.SolvePartOne, day02.SolvePartTwo},
	{day03.SolvePartOne, day03.SolvePartTwo},
	{day04.SolvePartOne, day04.SolvePartTwo},
	{day05.SolvePartOne, day05.SolvePartTwo},
	{day06.SolvePartOne, day06.SolvePartTwo},
}
