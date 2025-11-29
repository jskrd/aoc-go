package day01

func SolveLevel1(input string) int {
	var floor = 0

	for _, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}

	return floor
}

func SolveLevel2(input string) int {
	var floor = 0

	for index, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}

		if floor == -1 {
			return index + 1
		}
	}

	panic("Unreachable")
}
