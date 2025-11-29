package day01

import (
	"strings"
)

func SolveLevel1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		sum += findCalibrationValueLevel1(line)
	}
	return sum
}

func findCalibrationValueLevel1(line string) int {
	var first, last rune

	for i, j := 0, len(line)-1; i <= j; {
		if first == 0 && line[i] >= '1' && line[i] <= '9' {
			first = rune(line[i])
		}
		if line[j] >= '0' && line[j] <= '9' {
			last = rune(line[j])
		}

		if first == 0 {
			i++
		}
		if last == 0 {
			j--
		}
		if first != 0 && last != 0 {
			break
		}
	}

	return (int(first-'0') * 10) + int(last-'0')
}

func SolveLevel2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		sum += findCalibrationValueLevel2(line)
	}
	return sum
}

var digits = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
	"1":     '1',
	"2":     '2',
	"3":     '3',
	"4":     '4',
	"5":     '5',
	"6":     '6',
	"7":     '7',
	"8":     '8',
	"9":     '9',
}

func findCalibrationValueLevel2(line string) int {
	var first, last rune
	for i := 0; i < len(line); i++ {
		for word, digit := range digits {
			high := i + len(word)

			if high > len(line) {
				continue
			}

			if line[i:high] == word {
				if first == 0 {
					first = digit
				}

				last = digit
			}
		}
	}

	return (int(first-'0') * 10) + int(last-'0')
}
