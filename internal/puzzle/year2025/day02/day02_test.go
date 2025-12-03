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

func TestSolveLevel2(t *testing.T) {
	expected := 4174379265
	actual := SolveLevel2("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")

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

func TestIsRepeatedAtLeastTwice(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{11, true},
		{12, false},
		{13, false},
		{14, false},
		{15, false},
		{16, false},
		{17, false},
		{18, false},
		{19, false},
		{20, false},
		{21, false},
		{22, true},
		{95, false},
		{96, false},
		{97, false},
		{98, false},
		{99, true},
		{100, false},
		{101, false},
		{102, false},
		{103, false},
		{104, false},
		{105, false},
		{106, false},
		{107, false},
		{108, false},
		{109, false},
		{110, false},
		{111, true},
		{112, false},
		{113, false},
		{114, false},
		{115, false},
		{998, false},
		{999, true},
		{1000, false},
		{1001, false},
		{1002, false},
		{1003, false},
		{1004, false},
		{1005, false},
		{1006, false},
		{1007, false},
		{1008, false},
		{1009, false},
		{1010, true},
		{1011, false},
		{1012, false},
	}

	for _, test := range tests {
		actual := isRepeatedAtLeastTwice(test.id)
		if actual != test.expected {
			t.Errorf("Expected %v but got %v with `%d`", test.expected, actual, test.id)
		}
	}
}
