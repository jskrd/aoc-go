package day01

import (
	"slices"
	"testing"
)

func TestSolveLevel1(t *testing.T) {
	expected := 3
	actual := SolveLevel1("L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolveLevel2(t *testing.T) {
	expected := 6
	actual := SolveLevel2("L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []rotation
	}{
		{
			"L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			[]rotation{
				{'L', 68},
				{'L', 30},
				{'R', 48},
				{'L', 5},
				{'R', 60},
				{'L', 55},
				{'L', 1},
				{'L', 99},
				{'R', 14},
				{'L', 82},
			},
		},
		{
			"\nL68\nL82\n\n",
			[]rotation{{'L', 68}, {'L', 82}},
		},
	}

	for _, test := range tests {
		actual := parseInput(test.input)
		if !slices.Equal(test.expected, actual) {
			t.Errorf("Expected %v but got %v with input `%s`", test.expected, actual, test.input)
		}
	}
}
