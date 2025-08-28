package day04

import (
	"strings"
)

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func SolvePartOne(input string) int {
	grid := parseInput(input)

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != 'X' {
				continue
			}
			if searchUp(&grid, row, col) {
				count++
			}
			if searchUpRight(&grid, row, col) {
				count++
			}
			if searchRight(&grid, row, col) {
				count++
			}
			if searchDownRight(&grid, row, col) {
				count++
			}
			if searchDown(&grid, row, col) {
				count++
			}
			if searchDownLeft(&grid, row, col) {
				count++
			}
			if searchLeft(&grid, row, col) {
				count++
			}
			if searchUpLeft(&grid, row, col) {
				count++
			}
		}
	}

	return count
}

func SolvePartTwo(input string) int {
	grid := parseInput(input)

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != 'A' {
				continue
			}
			if searchX(&grid, row, col) {
				count++
			}
		}
	}

	return count
}

func searchUp(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchUp called on non-X cell")
	}

	if row < 3 {
		return false
	}

	if (*grid)[row-1][col] != 'M' ||
		(*grid)[row-2][col] != 'A' ||
		(*grid)[row-3][col] != 'S' {
		return false
	}

	return true
}

func searchUpRight(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchUpRight called on non-X cell")
	}

	if row-3 < 0 || col+3 > len((*grid)[row])-1 {
		return false
	}

	if (*grid)[row-1][col+1] != 'M' ||
		(*grid)[row-2][col+2] != 'A' ||
		(*grid)[row-3][col+3] != 'S' {
		return false
	}

	return true
}

func searchRight(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchRight called on non-X cell")
	}

	if col+3 > len((*grid)[row])-1 {
		return false
	}

	if (*grid)[row][col+1] != 'M' ||
		(*grid)[row][col+2] != 'A' ||
		(*grid)[row][col+3] != 'S' {
		return false
	}

	return true
}

func searchDownRight(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchDownRight called on non-X cell")
	}

	if row+3 > len(*grid)-1 || col+3 > len((*grid)[row])-1 {
		return false
	}

	if (*grid)[row+1][col+1] != 'M' ||
		(*grid)[row+2][col+2] != 'A' ||
		(*grid)[row+3][col+3] != 'S' {
		return false
	}

	return true
}

func searchDown(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchDown called on non-X cell")
	}

	if row+3 > len(*grid)-1 {
		return false
	}

	if (*grid)[row+1][col] != 'M' ||
		(*grid)[row+2][col] != 'A' ||
		(*grid)[row+3][col] != 'S' {
		return false
	}

	return true
}

func searchDownLeft(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchDownLeft called on non-X cell")
	}

	if row+3 > len(*grid)-1 || col-3 < 0 {
		return false
	}

	if (*grid)[row+1][col-1] != 'M' ||
		(*grid)[row+2][col-2] != 'A' ||
		(*grid)[row+3][col-3] != 'S' {
		return false
	}

	return true
}

func searchLeft(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchLeft called on non-X cell")
	}

	if col-3 < 0 {
		return false
	}

	if (*grid)[row][col-1] != 'M' ||
		(*grid)[row][col-2] != 'A' ||
		(*grid)[row][col-3] != 'S' {
		return false
	}

	return true
}

func searchUpLeft(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'X' {
		panic("searchUpLeft called on non-X cell")
	}

	if row-3 < 0 || col-3 < 0 {
		return false
	}

	if (*grid)[row-1][col-1] != 'M' ||
		(*grid)[row-2][col-2] != 'A' ||
		(*grid)[row-3][col-3] != 'S' {
		return false
	}

	return true
}

func searchX(grid *[][]rune, row, col int) bool {
	if (*grid)[row][col] != 'A' {
		panic("searchXmas called on non-A cell")
	}

	if row-1 < 0 ||
		row+1 > len(*grid)-1 ||
		col-1 < 0 ||
		col+1 > len((*grid)[row])-1 {
		return false
	}

	tl := (*grid)[row-1][col-1]
	tr := (*grid)[row-1][col+1]
	br := (*grid)[row+1][col+1]
	bl := (*grid)[row+1][col-1]
	if !(tl == 'M' && tr == 'M' && br == 'S' && bl == 'S') &&
		!(tl == 'S' && tr == 'M' && br == 'M' && bl == 'S') &&
		!(tl == 'S' && tr == 'S' && br == 'M' && bl == 'M') &&
		!(tl == 'M' && tr == 'S' && br == 'S' && bl == 'M') {
		return false
	}

	return true
}
