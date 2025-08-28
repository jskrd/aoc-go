package day01_test

import (
	"strings"
	"testing"

	"github.com/jskrd/aoc-go/internal/y2023/day01"
)

const inputPart1 = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"

const inputPart2 = "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"

func TestSolvePart1(t *testing.T) {
	lines := strings.Split(inputPart1, "\n")

	expected := 142
	actual := day01.SolvePart1(&lines)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	lines := strings.Split(inputPart2, "\n")

	expected := 281
	actual := day01.SolvePart2(&lines)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	lines := strings.Split(inputPart1, "\n")

	for i := 0; i < b.N; i++ {
		day01.SolvePart1(&lines)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	lines := strings.Split(inputPart2, "\n")

	for i := 0; i < b.N; i++ {
		day01.SolvePart2(&lines)
	}
}
