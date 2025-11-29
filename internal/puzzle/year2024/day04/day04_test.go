package day04

import (
	"reflect"
	"testing"
)

const input = "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

func TestParseInput(t *testing.T) {
	expected := [][]rune{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}
	actual := parseInput(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolveLevel1(t *testing.T) {
	expected := 18
	actual := SolveLevel1(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolveLevel2(t *testing.T) {
	expected := 9
	actual := SolveLevel2(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

// The remaining tests validate the directional search helpers
// using the same inline grids; they don't read files.
func TestSearchUp(t *testing.T) {
	grid := [][]rune{
		{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
		{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
		{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
		{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
		{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
		{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 4, false},
		{0, 5, false},
		{1, 4, false},
		{3, 9, false},
		{4, 0, false},
		{4, 6, true},
		{5, 0, false},
		{5, 6, false},
		{9, 1, false},
		{9, 3, false},
		{9, 5, false},
		{9, 9, true},
	}

	for _, test := range tests {
		actual := searchUp(&grid, test.row, test.col)
		if test.expected != actual {
			t.Errorf("Expected %t but got %t", test.expected, actual)
		}
	}
}
