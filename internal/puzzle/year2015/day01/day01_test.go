package day01

import "testing"

func TestSolveLevel1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, test := range tests {
		actual := SolveLevel1(test.input)
		if test.expected != actual {
			t.Errorf("Expected %d but got %d with input `%s`", test.expected, actual, test.input)
		}
	}
}

func TestSolveLevel2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, test := range tests {
		actual := SolveLevel2(test.input)
		if test.expected != actual {
			t.Errorf("Expected %d but got %d with input `%s`", test.expected, actual, test.input)
		}
	}
}
