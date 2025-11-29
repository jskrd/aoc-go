package day01

import (
	"slices"
	"testing"
)

func TestSolveLevel1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"R2, L3", 5},
		{"R2, R2, R2", 2},
		{"R5, L5, R5, R3", 12},
	}

	for _, test := range tests {
		actual := SolveLevel1(test.input)
		if test.expected != actual {
			t.Errorf("Expected %d but got %d with input `%s`", test.expected, actual, test.input)
		}
	}
}

func TestSolveLevel2(t *testing.T) {
	expected := 4
	actual := SolveLevel2("R8, R4, R4, R8")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []instruction
	}{
		{
			"R2, L3",
			[]instruction{{'R', 2}, {'L', 3}},
		},
		{
			"R2, R2, R2",
			[]instruction{{'R', 2}, {'R', 2}, {'R', 2}},
		},
		{
			"R5, L5, R5, R3",
			[]instruction{{'R', 5}, {'L', 5}, {'R', 5}, {'R', 3}},
		},
		{
			"L1, L12, L123, R3, R32, R321",
			[]instruction{
				{'L', 1},
				{'L', 12},
				{'L', 123},
				{'R', 3},
				{'R', 32},
				{'R', 321},
			},
		},
		{
			"\nR2, L3\n\n\n",
			[]instruction{{'R', 2}, {'L', 3}},
		},
	}

	for _, test := range tests {
		actual := parseInput(test.input)
		if !slices.Equal(test.expected, actual) {
			t.Errorf("Expected %v but got %v with input `%s`", test.expected, actual, test.input)
		}
	}
}
