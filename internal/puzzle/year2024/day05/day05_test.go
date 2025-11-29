package day05

import (
	"reflect"
	"testing"
)

const input = "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47"

func TestParseInput(t *testing.T) {
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
	actual := parseInput(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestSolveLevel1(t *testing.T) {
	expected := 143
	actual := SolveLevel1(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolveLevel2(t *testing.T) {
	expected := 123
	actual := SolveLevel2(input)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestIsCorrectOrdering(t *testing.T) {
	example := parseInput(input)

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
	example := parseInput(input)

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
	example := parseInput(input)

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
	example := parseInput(input)

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
