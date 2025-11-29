# 🎄 Advent of Code

[![Test](https://github.com/jskrd/aoc-go/actions/workflows/test.yml/badge.svg)](https://github.com/jskrd/aoc-go/actions/workflows/test.yml)

## Setup

Create a `session.txt` file in the project root with your Advent of Code session cookie value. To find your session value, log in to [adventofcode.com](https://adventofcode.com), open your browser's developer tools, and copy the value of the `session` cookie.

```sh
$ go run cmd/solver/main.go -h
Usage of ...:
	-day int
		The day of the puzzle to solve
	-level int
		The level of the puzzle to solve
	-year int
		The year of the puzzles to solve
```

```sh
$ go run cmd/solver/main.go -year=2015 -day=1 -level=1
74
Answer submitted successfully!
```
