package day01_test

import (
	"testing"

	"github.com/jskrd/aoc-go/internal/puzzle/year2023/day01"
)

func TestSolveLevel1(t *testing.T) {
	expected := 142
	actual := day01.SolveLevel1("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestSolveLevel2(t *testing.T) {
	expected := 281
	actual := day01.SolveLevel2("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")

	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
