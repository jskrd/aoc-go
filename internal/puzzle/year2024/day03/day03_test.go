package day03

import (
	"reflect"
	"testing"
)

const input1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const input2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input1,
			[]string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
		},
		{
			input2,
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

func TestSolveLevel1(t *testing.T) {
	expected := 161
	actual := SolveLevel1(input1)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolveLevel2(t *testing.T) {
	expected := 48
	actual := SolveLevel2(input2)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
