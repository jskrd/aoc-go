package day01

import (
	"runtime"
	"strings"
	"sync"
)

func SolveLevel1(input string) int {
	lines := strings.Split(input, "\n")

	jobs := make(chan string, len(lines))
	results := make(chan int, len(lines))
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go workerLevel1(jobs, results, &wg)
	}

	go func() {
		for _, line := range lines {
			jobs <- line
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var sum int
	for result := range results {
		sum += result
	}
	return sum
}

func workerLevel1(jobs <-chan string, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range jobs {
		results <- findCalibrationValueLevel1(line)
	}
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

	jobs := make(chan string, len(lines))
	results := make(chan int, len(lines))
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go workerLevel2(jobs, results, &wg)
	}

	go func() {
		for _, line := range lines {
			jobs <- line
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var sum int
	for result := range results {
		sum += result
	}
	return sum
}

func workerLevel2(lines <-chan string, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range lines {
		results <- findCalibrationValueLevel2(line)
	}
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
