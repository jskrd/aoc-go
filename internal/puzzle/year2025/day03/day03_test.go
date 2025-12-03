package day03

import (
	"slices"
	"testing"
)

func TestSolveLevel1(t *testing.T) {
	expected := 357
	actual := SolveLevel1("987654321111111\n811111111111119\n234234234234278\n818181911112111")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			"987654321111111\n811111111111119\n234234234234278\n818181911112111",
			[]string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
		},
		{
			"\n987654321111111\n811111111111119\n\n",
			[]string{"987654321111111", "811111111111119"},
		},
	}

	for _, test := range tests {
		actual := parseInput(test.input)
		if !slices.Equal(test.expected, actual) {
			t.Errorf("Expected %v but got %v with input `%s`", test.expected, actual, test.input)
		}
	}
}

func TestFindLargestJoltage(t *testing.T) {
	tests := []struct {
		bank     string
		expected int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, test := range tests {
		actual := findLargestJoltage(test.bank)
		if actual != test.expected {
			t.Errorf("Expected %d but got %d with input `%s`", test.expected, actual, test.bank)
		}
	}
}
