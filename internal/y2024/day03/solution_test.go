package day03

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input1, _ := os.ReadFile("../../../test/data/2024/day03/input_example1.txt")
	input2, _ := os.ReadFile("../../../test/data/2024/day03/input_example2.txt")

	tests := []struct {
		input    string
		expected []string
	}{
		{
			string(input1),
			[]string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
		},
		{
			string(input2),
			[]string{"mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"},
		},
	}

	for _, test := range tests {
		actual := parseInput(test.input)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestSolvePartOne(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day03/input_example1.txt")

	expected := 161
	actual := SolvePartOne(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day03/input_example2.txt")

	expected := 48
	actual := SolvePartTwo(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
