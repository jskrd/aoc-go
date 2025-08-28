package years

import (
	"github.com/jskrd/aoc-go/internal/aoc"
	y2023 "github.com/jskrd/aoc-go/internal/y2023"
	y2024 "github.com/jskrd/aoc-go/internal/y2024"
)

var Years = map[int][25]aoc.DayParts{
	2023: y2023.Days,
	2024: y2024.Days,
}
