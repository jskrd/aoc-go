package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		year     int
		day      int
		part     int
		input    string
		expected int
	}{
		{2023, 1, 1, "../../test/data/2023/day01/input.txt", 53651},
		{2023, 1, 2, "../../test/data/2023/day01/input.txt", 53894},
		{2023, 2, 1, "../../test/data/2023/day02/input.txt", 2377},
		{2023, 2, 2, "../../test/data/2023/day02/input.txt", 71220},
		{2024, 1, 1, "../../test/data/2024/day01/input.txt", 3508942},
		{2024, 1, 2, "../../test/data/2024/day01/input.txt", 26593248},
		{2024, 2, 1, "../../test/data/2024/day02/input.txt", 279},
		{2024, 2, 2, "../../test/data/2024/day02/input.txt", 343},
		{2024, 3, 1, "../../test/data/2024/day03/input.txt", 182780583},
		{2024, 3, 2, "../../test/data/2024/day03/input.txt", 90772405},
		{2024, 4, 1, "../../test/data/2024/day04/input.txt", 2642},
		{2024, 4, 2, "../../test/data/2024/day04/input.txt", 1974},
		{2024, 5, 1, "../../test/data/2024/day05/input.txt", 4637},
		{2024, 5, 2, "../../test/data/2024/day05/input.txt", 6370},
		{2024, 6, 1, "../../test/data/2024/day06/input.txt", 5145},
		{2024, 6, 2, "../../test/data/2024/day06/input.txt", 1523},
	}

	for _, test := range tests {
		actual := solve(test.year, test.day, test.part, test.input)
		if test.expected != actual {
			t.Errorf("Expected %d but got %d", test.expected, actual)
		}
	}
}
