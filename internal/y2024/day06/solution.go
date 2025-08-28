package day06

import (
	"strings"
	"sync"
)

func parseInput(input string) Lab {
	lab := Lab{}
	lab.guardPositions = make(map[Guard]bool)
	lab.obstacles = make(map[[2]int]bool)
	for row, line := range strings.Split(input, "\n") {
		for col, cell := range line {
			if cell == '^' || cell == '>' || cell == 'v' || cell == '<' {
				direction := Up
				if cell == '>' {
					direction = Right
				} else if cell == 'v' {
					direction = Down
				} else if cell == '<' {
					direction = Left
				}

				lab.guard = Guard{
					direction: direction,
					position:  [2]int{row, col},
				}
				lab.guardStartingPosition = [2]int{row, col}
			} else if cell == '#' {
				lab.obstacles[[2]int{row, col}] = true
			}
		}
	}
	lab.size = [2]int{
		len(strings.Split(input, "\n")),
		len(strings.Split(input, "\n")[0]),
	}
	return lab
}

func SolvePartOne(input string) int {
	lab := parseInput(input)

	lab.play()

	return lab.distinctGuardPositionsCount()
}

func SolvePartTwo(input string) int {
	lab := parseInput(input)

	newObstacles := map[[2]int]bool{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	for !lab.hasGuardLeftMap() && !lab.hasGuardLooped() {
		lab.guardPositions[lab.guard] = true

		if lab.hasGuardHitObstacle() {
			lab.guard.turn()
			if lab.hasGuardHitObstacle() {
				lab.guard.turn()
			}
		}

		if !lab.hasGuardHitObstacle() && lab.hasGuardObstacleOnRight() {
			newObstacle := lab.guard.ahead()

			if lab.guardStartingPosition == newObstacle ||
				newObstacle[0] < 0 ||
				newObstacle[1] < 0 ||
				newObstacle[0] >= lab.size[0] ||
				newObstacle[1] >= lab.size[1] {
				break
			}

			labCopy := parseInput(input)

			wg.Add(1)
			go func(lab Lab, obstacle [2]int) {
				defer wg.Done()

				lab.obstacles[obstacle] = true
				lab.play()

				if lab.hasGuardLooped() {
					mu.Lock()
					newObstacles[obstacle] = true
					mu.Unlock()
				}
			}(labCopy, newObstacle)
		}

		lab.guard.step()
	}

	wg.Wait()
	return len(newObstacles)
}
