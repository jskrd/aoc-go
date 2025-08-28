package day06

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Guard struct {
	direction Direction
	position  [2]int
}

func (guard *Guard) ahead() [2]int {
	coord := guard.position
	if guard.direction == Up {
		coord[0]--
	} else if guard.direction == Right {
		coord[1]++
	} else if guard.direction == Down {
		coord[0]++
	} else if guard.direction == Left {
		coord[1]--
	} else {
		panic("Guard is not facing a valid direction")
	}
	return coord
}

func (guard *Guard) step() {
	guard.position = guard.ahead()
}

func (guard *Guard) turn() {
	if guard.direction == Up {
		guard.direction = Right
	} else if guard.direction == Right {
		guard.direction = Down
	} else if guard.direction == Down {
		guard.direction = Left
	} else if guard.direction == Left {
		guard.direction = Up
	} else {
		panic("Guard is not facing a valid direction")
	}
}
