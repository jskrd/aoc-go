package day02

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day02/input_example.txt")

	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	actual := parseInput(string(input))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolvePartOne(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day02/input_example.txt")

	expected := 2
	actual := SolvePartOne(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day02/input_example.txt")

	expected := 4
	actual := SolvePartTwo(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestIsSafe(t *testing.T) {
	tests := []struct {
		report   []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		actual := isSafe(test.report)
		if test.expected != actual {
			t.Errorf(
				"Expected %t but got %t for %v",
				test.expected,
				actual,
				test.report,
			)
		}
	}
}

func TestGetAllPossibleReportsRemovingOne(t *testing.T) {
	tests := []struct {
		report   []int
		expected [][]int
	}{
		{
			[]int{1},
			[][]int{},
		},
		{
			[]int{1, 2},
			[][]int{{2}, {1}},
		},
		{
			[]int{1, 2, 3},
			[][]int{{2, 3}, {1, 3}, {1, 2}},
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{{2, 3, 4}, {1, 3, 4}, {1, 2, 4}, {1, 2, 3}},
		},
	}

	for _, test := range tests {
		actual := getAllPossibleReportsRemovingOne(test.report)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf(
				"Expected %v but got %v for %v",
				test.expected,
				actual,
				test.report,
			)
		}
	}
}
