package day02_test

import (
	"strings"
	"testing"

	"github.com/jskrd/aoc-go/internal/y2023/day02"
)

const input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"

func TestSolvePart1(t *testing.T) {
	lines := strings.Split(input, "\n")

	expected := 8
	actual := day02.SolvePart1(&lines)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolvePart2(t *testing.T) {
	lines := strings.Split(input, "\n")

	expected := 2286
	actual := day02.SolvePart2(&lines)

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	lines := strings.Split(input, "\n")

	for i := 0; i < b.N; i++ {
		day02.SolvePart1(&lines)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	lines := strings.Split(input, "\n")

	for i := 0; i < b.N; i++ {
		day02.SolvePart2(&lines)
	}
}
