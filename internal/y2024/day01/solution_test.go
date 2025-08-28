package day01

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day01/input_example.txt")

	expected := [][]int{
		{3, 4, 2, 1, 3, 3},
		{4, 3, 5, 3, 9, 3},
	}
	actual := parseInput(string(input))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolvePartOne(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day01/input_example.txt")

	expected := 11
	actual := SolvePartOne(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day01/input_example.txt")

	expected := 31
	actual := SolvePartTwo(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
