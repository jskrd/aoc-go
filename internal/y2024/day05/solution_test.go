package day05

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day05/input_example.txt")

	expected := parsedInput{
		[][2]int{
			{47, 53},
			{97, 13},
			{97, 61},
			{97, 47},
			{75, 29},
			{61, 13},
			{75, 53},
			{29, 13},
			{97, 29},
			{53, 29},
			{61, 53},
			{97, 53},
			{61, 29},
			{47, 13},
			{75, 47},
			{97, 75},
			{47, 61},
			{75, 61},
			{47, 29},
			{75, 13},
			{53, 13},
		},
		[][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		},
	}
	actual := parseInput(string(input))

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolvePartOne(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day05/input_example.txt")

	expected := 143
	actual := SolvePartOne(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day05/input_example.txt")

	expected := 123
	actual := SolvePartTwo(string(input))

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestIsCorrectOrdering(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day05/input_example.txt")
	example := parseInput(string(input))

	tests := []bool{true, true, true, false, false, false}

	for i, test := range tests {
		expected := test
		actual := isCorrectOrdering(example.orderings, example.updates[i])

		if expected != actual {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}

func TestFindOrderingsBefore(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day05/input_example.txt")
	example := parseInput(string(input))

	tests := []struct {
		page     int
		expected []int
	}{
		{75, []int{97}},
		{47, []int{97, 75}},
		{61, []int{97, 47, 75}},
		{53, []int{47, 75, 61, 97}},
		{29, []int{75, 97, 53, 61, 47}},
	}

	for _, test := range tests {
		expected := test.expected
		actual := findOrderingsBefore(example.orderings, test.page)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}

func TestFindOrderingsAfter(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day05/input_example.txt")
	example := parseInput(string(input))

	tests := []struct {
		page     int
		expected []int
	}{
		{75, []int{29, 53, 47, 61, 13}},
		{47, []int{53, 13, 61, 29}},
		{61, []int{13, 53, 29}},
		{53, []int{29, 13}},
		{29, []int{13}},
	}

	for _, test := range tests {
		expected := test.expected
		actual := findOrderingsAfter(example.orderings, test.page)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}

func TestCorrectUpdate(t *testing.T) {
	input, _ := os.ReadFile("../../../test/data/2024/day05/input_example.txt")
	example := parseInput(string(input))

	tests := []struct {
		update   []int
		expected []int
	}{
		{[]int{75, 97, 47, 61, 53}, []int{97, 75, 47, 61, 53}},
		{[]int{61, 13, 29}, []int{61, 29, 13}},
		{[]int{97, 13, 75, 29, 47}, []int{97, 75, 47, 29, 13}},
	}

	for _, test := range tests {
		expected := test.expected
		actual := correctUpdate(example.orderings, test.update)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}
