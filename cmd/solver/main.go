package main

import (
	"flag"
	"fmt"

	"github.com/jskrd/aoc-go/internal/aoc"
	"github.com/jskrd/aoc-go/internal/puzzle"
)

func main() {
	year := flag.Int("year", 0, "The year of the puzzles to solve")
	day := flag.Int("day", 0, "The day of the puzzle to solve")
	level := flag.Int("level", 0, "The level of the puzzle to solve")
	flag.Parse()

	session, err := aoc.GetSession()
	if err != nil {
		fmt.Println("Failed to get session:", err)
		return
	}

	solution, err := puzzle.GetSolution(*year, *day, *level)
	if err != nil {
		fmt.Println("Failed to get solution:", err)
		return
	}

	input, err := aoc.FetchInput(session, *year, *day, *level)
	if err != nil {
		fmt.Println("Failed to fetch input:", err)
		return
	}

	answer := solution(input)
	fmt.Println(answer)

	err = aoc.SubmitAnswer(session, *year, *day, *level, answer)
	if err != nil {
		fmt.Println("Failed to submit answer:", err)
		return
	}

	fmt.Println("Answer submitted successfully!")
}
