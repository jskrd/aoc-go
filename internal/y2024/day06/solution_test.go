package day06

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day06/input_example.txt")

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
	actual := parseInput(string(input))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolvePartOne(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day06/input_example.txt")

	expected := 41
	actual := SolvePartOne(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day06/input_example.txt")

	expected := 6
	actual := SolvePartTwo(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
