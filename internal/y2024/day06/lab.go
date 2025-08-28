package day06

type Lab struct {
	guard                 Guard
	guardPositions        map[Guard]bool
	guardStartingPosition [2]int
	obstacles             map[[2]int]bool
	size                  [2]int
}

func (lab *Lab) distinctGuardPositionsCount() int {
	distinct := make(map[[2]int]bool)
	for guard := range lab.guardPositions {
		distinct[guard.position] = true
	}
	return len(distinct)
}

func (lab *Lab) hasGuardHitObstacle() bool {
	return lab.obstacles[lab.guard.ahead()]
}

func (lab *Lab) hasGuardLeftMap() bool {
	return lab.guard.position[0] < 0 ||
		lab.guard.position[1] < 0 ||
		lab.guard.position[0] >= lab.size[0] ||
		lab.guard.position[1] >= lab.size[1]
}

func (lab *Lab) hasGuardLooped() bool {
	return lab.guardPositions[lab.guard]
}

func (lab *Lab) hasGuardObstacleOnRight() bool {
	coord := lab.guard.position
	for !lab.obstacles[coord] {
		if lab.guard.direction == Up {
			coord[1]++
		} else if lab.guard.direction == Right {
			coord[0]++
		} else if lab.guard.direction == Down {
			coord[1]--
		} else if lab.guard.direction == Left {
			coord[0]--
		} else {
			panic("Guard is not facing a valid direction")
		}

		if coord[0] < 0 || coord[1] < 0 {
			return false
		}

		if coord[0] >= lab.size[0] || coord[1] >= lab.size[1] {
			return false
		}
	}
	return true
}

func (lab *Lab) moveGuard() {
	lab.guardPositions[lab.guard] = true

	if lab.hasGuardHitObstacle() {
		lab.guard.turn()
		if lab.hasGuardHitObstacle() {
			lab.guard.turn()
		}
	}
	lab.guard.step()
}

func (lab *Lab) play() {
	for !lab.hasGuardLooped() && !lab.hasGuardLeftMap() {
		lab.moveGuard()
	}
}
