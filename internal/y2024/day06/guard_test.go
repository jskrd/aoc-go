package day06

import "testing"

func TestAhead(t *testing.T) {
	tests := []struct {
		guard    Guard
		expected [2]int
	}{
		{
			Guard{direction: Up, position: [2]int{0, 0}},
			[2]int{-1, 0},
		},
		{
			Guard{direction: Right, position: [2]int{0, 0}},
			[2]int{0, 1},
		},
		{
			Guard{direction: Down, position: [2]int{0, 0}},
			[2]int{1, 0},
		},
		{
			Guard{direction: Left, position: [2]int{0, 0}},
			[2]int{0, -1},
		},
	}

	for _, test := range tests {
		actual := test.guard.ahead()
		if test.expected != actual {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		guard    Guard
		expected [2]int
	}{
		{
			Guard{direction: Up, position: [2]int{0, 0}},
			[2]int{-1, 0},
		},
		{
			Guard{direction: Right, position: [2]int{0, 0}},
			[2]int{0, 1},
		},
		{
			Guard{direction: Down, position: [2]int{0, 0}},
			[2]int{1, 0},
		},
		{
			Guard{direction: Left, position: [2]int{0, 0}},
			[2]int{0, -1},
		},
	}

	for _, test := range tests {
		test.guard.step()
		actual := test.guard.position
		if test.expected != actual {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestTurn(t *testing.T) {
	tests := []struct {
		guard    Guard
		expected Direction
	}{
		{
			Guard{direction: Up, position: [2]int{0, 0}},
			Right,
		},
		{
			Guard{direction: Right, position: [2]int{0, 0}},
			Down,
		},
		{
			Guard{direction: Down, position: [2]int{0, 0}},
			Left,
		},
		{
			Guard{direction: Left, position: [2]int{0, 0}},
			Up,
		},
	}

	for _, test := range tests {
		test.guard.turn()
		actual := test.guard.direction
		if test.expected != actual {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}
