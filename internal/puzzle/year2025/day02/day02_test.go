package day02

import (
	"slices"
	"testing"
)

func TestSolveLevel1(t *testing.T) {
	expected := 1227775554
	actual := SolveLevel1("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []idRange
	}{
		{
			"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			[]idRange{
				{11, 22},
				{95, 115},
				{998, 1012},
				{1188511880, 1188511890},
				{222220, 222224},
				{1698522, 1698528},
				{446443, 446449},
				{38593856, 38593862},
				{565653, 565659},
				{824824821, 824824827},
				{2121212118, 2121212124},
			},
		},
		{
			"\n11-22,95-115\n\n",
			[]idRange{{11, 22}, {95, 115}},
		},
	}

	for _, test := range tests {
		actual := parseInput(test.input)
		if !slices.Equal(test.expected, actual) {
			t.Errorf("Expected %v but got %v with input `%s`", test.expected, actual, test.input)
		}
	}
}
