package day06

import (
	"reflect"
	"testing"
)

const input = "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#..."

func TestParseInput(t *testing.T) {
	expected := Lab{
		guard:                 Guard{direction: Up, position: [2]int{6, 4}},
		guardPositions:        map[Guard]bool{},
		guardStartingPosition: [2]int{6, 4},
		obstacles: map[[2]int]bool{
			{0, 4}: true,
			{1, 9}: true,
			{3, 2}: true,
			{4, 7}: true,
			{6, 1}: true,
			{7, 8}: true,
			{8, 0}: true,
			{9, 6}: true,
		},
		size: [2]int{10, 10},
	}
	actual := parseInput(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolveLevel1(t *testing.T) {
	expected := 41
	actual := SolveLevel1(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolveLevel2(t *testing.T) {
	expected := 6
	actual := SolveLevel2(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
