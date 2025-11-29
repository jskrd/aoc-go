package puzzle

import (
	"fmt"

	"github.com/jskrd/aoc-go/internal/puzzle/year2015"
	"github.com/jskrd/aoc-go/internal/puzzle/year2023"
	"github.com/jskrd/aoc-go/internal/puzzle/year2024"
)

var Solutions = map[int]map[int]map[int]func(input string) int{
	2015: year2015.Solutions,
	2023: year2023.Solutions,
	2024: year2024.Solutions,
}

func GetSolution(year int, day int, level int) (func(input string) int, error) {
	days, ok := Solutions[year]
	if !ok {
		return nil, fmt.Errorf("year %d not found", year)
	}

	levels, ok := days[day]
	if !ok {
		return nil, fmt.Errorf("day %d not found", day)
	}

	_, ok = levels[level]
	if !ok {
		return nil, fmt.Errorf("level %d not found", level)
	}

	return levels[level], nil
}
