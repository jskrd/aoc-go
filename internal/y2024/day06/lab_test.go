package day06

import (
	"reflect"
	"testing"
)

func TestDistinctGuardPositionsCount(t *testing.T) {
	tests := []struct {
		guardPositions map[Guard]bool
		expected       int
	}{
		{
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}: true,
			},
			1,
		},
		{
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
			},
			2,
		},
		{
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
			},
			3,
		},
		{
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
				{direction: Left, position: [2]int{2, 1}}:  true,
				{direction: Left, position: [2]int{2, 0}}:  true,
				{direction: Up, position: [2]int{1, 0}}:    true,
				{direction: Right, position: [2]int{1, 1}}: true,
			},
			6,
		},
	}

	for _, test := range tests {
		lab := Lab{
			guardPositions: test.guardPositions,
		}
		actual := lab.distinctGuardPositionsCount()
		if test.expected != actual {
			t.Errorf("Expected %d but got %d", test.expected, actual)
		}
	}
}

func TestHasGuardHitObstacle(t *testing.T) {
	tests := []struct {
		direction Direction
		obstacles map[[2]int]bool
		expected  bool
	}{
		// . . .
		// # ^ #
		// . # .
		{Up, map[[2]int]bool{
			{1, 0}: true,
			{1, 2}: true,
			{2, 1}: true,
		}, false},

		// . # .
		// # > .
		// . # .
		{Right, map[[2]int]bool{
			{0, 1}: true,
			{1, 0}: true,
			{2, 1}: true,
		}, false},

		// . # .
		// # v #
		// . . .
		{Down, map[[2]int]bool{
			{0, 1}: true,
			{1, 0}: true,
			{1, 2}: true,
		}, false},

		// . # .
		// . < #
		// . # .
		{Left, map[[2]int]bool{
			{0, 1}: true,
			{1, 2}: true,
			{2, 1}: true,
		}, false},

		// . # .
		// . ^ .
		// . . .
		{Up, map[[2]int]bool{{0, 1}: true}, true},

		// . . .
		// . > #
		// . . .
		{Right, map[[2]int]bool{{1, 2}: true}, true},

		// . . .
		// . v .
		// . # .
		{Down, map[[2]int]bool{{2, 1}: true}, true},

		// . . .
		// # < .
		// . . .
		{Left, map[[2]int]bool{{1, 0}: true}, true},
	}

	for _, test := range tests {
		lab := Lab{
			guard: Guard{
				direction: test.direction,
				position:  [2]int{1, 1},
			},
			obstacles: test.obstacles,
			size:      [2]int{3, 3},
		}
		actual := lab.hasGuardHitObstacle()
		if test.expected != actual {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestHasGuardLeftMap(t *testing.T) {
	tests := []struct {
		position [2]int
		expected bool
	}{
		{[2]int{-1, -1}, true},
		{[2]int{-1, 0}, true},
		{[2]int{-1, 1}, true},
		{[2]int{-1, 2}, true},
		{[2]int{0, -1}, true},
		{[2]int{0, 0}, false},
		{[2]int{0, 1}, false},
		{[2]int{0, 2}, true},
		{[2]int{1, -1}, true},
		{[2]int{1, 0}, false},
		{[2]int{1, 1}, false},
		{[2]int{1, 2}, true},
		{[2]int{2, -1}, true},
		{[2]int{2, 0}, true},
		{[2]int{2, 1}, true},
		{[2]int{2, 2}, true},
	}

	for _, test := range tests {
		lab := Lab{
			guard: Guard{
				direction: Up,
				position:  test.position,
			},
			size: [2]int{2, 2},
		}
		actual := lab.hasGuardLeftMap()
		if test.expected != actual {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestHasGuardLooped(t *testing.T) {
	tests := []struct {
		guard          Guard
		guardPositions map[Guard]bool
		expected       bool
	}{
		// . . .
		// . ^ .
		// . . .
		{
			Guard{direction: Up, position: [2]int{1, 1}},
			map[Guard]bool{},
			false,
		},

		// . . .
		// . X >
		// . . .
		{
			Guard{direction: Right, position: [2]int{1, 2}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}: true,
			},
			false,
		},

		// . . .
		// . X X
		// . . v
		{
			Guard{direction: Down, position: [2]int{2, 2}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
			},
			false,
		},

		// . . .
		// . X X
		// . < X
		{
			Guard{direction: Left, position: [2]int{2, 1}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
			},
			false,
		},

		// . . .
		// . X X
		// < X X
		{
			Guard{direction: Left, position: [2]int{2, 0}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
				{direction: Left, position: [2]int{2, 1}}:  true,
			},
			false,
		},

		// . . .
		// ^ X X
		// X X X
		{
			Guard{direction: Up, position: [2]int{1, 0}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
				{direction: Left, position: [2]int{2, 1}}:  true,
				{direction: Left, position: [2]int{2, 0}}:  true,
			},
			false,
		},

		// . . .
		// X > X
		// X X X
		{
			Guard{direction: Right, position: [2]int{1, 1}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
				{direction: Left, position: [2]int{2, 1}}:  true,
				{direction: Left, position: [2]int{2, 0}}:  true,
				{direction: Up, position: [2]int{1, 0}}:    true,
			},
			false,
		},

		// . . .
		// X X >
		// X X X
		{
			Guard{direction: Right, position: [2]int{1, 2}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
				{direction: Left, position: [2]int{2, 1}}:  true,
				{direction: Left, position: [2]int{2, 0}}:  true,
				{direction: Up, position: [2]int{1, 0}}:    true,
				{direction: Right, position: [2]int{1, 1}}: true,
			},
			true,
		},

		// . . .
		// X X X
		// X X v
		{
			Guard{direction: Down, position: [2]int{2, 2}},
			map[Guard]bool{
				{direction: Up, position: [2]int{1, 1}}:    true,
				{direction: Right, position: [2]int{1, 2}}: true,
				{direction: Down, position: [2]int{2, 2}}:  true,
				{direction: Left, position: [2]int{2, 1}}:  true,
				{direction: Left, position: [2]int{2, 0}}:  true,
				{direction: Up, position: [2]int{1, 0}}:    true,
				{direction: Right, position: [2]int{1, 1}}: true,
				{direction: Right, position: [2]int{1, 2}}: true,
			},
			true,
		},
	}

	for _, test := range tests {
		lab := Lab{
			guard:          test.guard,
			guardPositions: test.guardPositions,
		}
		actual := lab.hasGuardLooped()
		if test.expected != actual {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestHasGuardObstacleOnRight(t *testing.T) {
	tests := []struct {
		guard     Guard
		obstacles map[[2]int]bool
		expected  bool
	}{
		// . . .
		// ^ . .
		// . . .
		{
			Guard{direction: Up, position: [2]int{1, 0}},
			map[[2]int]bool{},
			false,
		},

		// . . .
		// ^ . #
		// . . .
		{
			Guard{direction: Up, position: [2]int{1, 0}},
			map[[2]int]bool{{1, 2}: true},
			true,
		},

		// . > .
		// . . .
		// . . .
		{
			Guard{direction: Right, position: [2]int{0, 1}},
			map[[2]int]bool{},
			false,
		},

		// . > .
		// . . .
		// . # .
		{
			Guard{direction: Right, position: [2]int{0, 1}},
			map[[2]int]bool{{2, 1}: true},
			true,
		},

		// . . .
		// . . v
		// . . .
		{
			Guard{direction: Down, position: [2]int{1, 2}},
			map[[2]int]bool{},
			false,
		},

		// . . .
		// # . v
		// . . .
		{
			Guard{direction: Down, position: [2]int{1, 2}},
			map[[2]int]bool{{1, 0}: true},
			true,
		},

		// . . .
		// . . .
		// . < .
		{
			Guard{direction: Left, position: [2]int{2, 1}},
			map[[2]int]bool{},
			false,
		},

		// . # .
		// . . .
		// . < .
		{
			Guard{direction: Left, position: [2]int{2, 1}},
			map[[2]int]bool{{0, 1}: true},
			true,
		},
	}

	for _, test := range tests {
		lab := Lab{
			guard:     test.guard,
			obstacles: test.obstacles,
			size:      [2]int{3, 3},
		}
		actual := lab.hasGuardObstacleOnRight()
		if test.expected != actual {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestMoveGuard(t *testing.T) {
	tests := []struct {
		lab      Lab
		expected Lab
	}{
		// . . . | . ^ .
		// . ^ . | . X .
		// . . . | . . .
		{
			Lab{
				guard:          Guard{direction: Up, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Up, position: [2]int{0, 1}},
				guardPositions: map[Guard]bool{
					{direction: Up, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{},
				size:      [2]int{3, 3},
			},
		},

		// . . . | . . .
		// . > . | . X >
		// . . . | . . .
		{
			Lab{
				guard:          Guard{direction: Right, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Right, position: [2]int{1, 2}},
				guardPositions: map[Guard]bool{
					{direction: Right, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{},
				size:      [2]int{3, 3},
			},
		},

		// . . . | . . .
		// . v . | . X .
		// . . . | . v .
		{
			Lab{
				guard:          Guard{direction: Down, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Down, position: [2]int{2, 1}},
				guardPositions: map[Guard]bool{
					{direction: Down, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{},
				size:      [2]int{3, 3},
			},
		},

		// . . . | . . .
		// . < . | < X .
		// . . . | . . .
		{
			Lab{
				guard:          Guard{direction: Left, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Left, position: [2]int{1, 0}},
				guardPositions: map[Guard]bool{
					{direction: Left, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{},
				size:      [2]int{3, 3},
			},
		},

		// . # . | . # .
		// . ^ . | . X >
		// . . . | . . .
		{
			Lab{
				guard:          Guard{direction: Up, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{{0, 1}: true},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Right, position: [2]int{1, 2}},
				guardPositions: map[Guard]bool{
					{direction: Up, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{{0, 1}: true},
				size:      [2]int{3, 3},
			},
		},

		// . . . | . . .
		// . > # | . X #
		// . . . | . v .
		{
			Lab{
				guard:          Guard{direction: Right, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{{1, 2}: true},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Down, position: [2]int{2, 1}},
				guardPositions: map[Guard]bool{
					{direction: Right, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{{1, 2}: true},
				size:      [2]int{3, 3},
			},
		},

		// . . . | . . .
		// . v . | < X .
		// . # . | . # .
		{
			Lab{
				guard:          Guard{direction: Down, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{{2, 1}: true},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Left, position: [2]int{1, 0}},
				guardPositions: map[Guard]bool{
					{direction: Down, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{{2, 1}: true},
				size:      [2]int{3, 3},
			},
		},

		// . . . | . ^ .
		// # < . | # X .
		// . . . | . . .
		{
			Lab{
				guard:          Guard{direction: Left, position: [2]int{1, 1}},
				guardPositions: map[Guard]bool{},
				obstacles:      map[[2]int]bool{{1, 0}: true},
				size:           [2]int{3, 3},
			},
			Lab{
				guard: Guard{direction: Up, position: [2]int{0, 1}},
				guardPositions: map[Guard]bool{
					{direction: Left, position: [2]int{1, 1}}: true,
				},
				obstacles: map[[2]int]bool{{1, 0}: true},
				size:      [2]int{3, 3},
			},
		},
	}

	for _, test := range tests {
		test.lab.moveGuard()

		actual := test.lab
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Expected %v but got %v", test.expected, actual)
		}
	}
}

func TestPlay(t *testing.T) {
	lab := Lab{
		guard:                 Guard{direction: Up, position: [2]int{6, 4}},
		guardPositions:        map[Guard]bool{},
		guardStartingPosition: [2]int{6, 4},
		obstacles: map[[2]int]bool{
			{0, 4}: true,
			{1, 9}: true,
			{3, 2}: true,
			{4, 7}: true,
			{6, 1}: true,
			{7, 8}: true,
			{8, 0}: true,
			{9, 6}: true,
		},
		size: [2]int{10, 10},
	}

	lab.play()

	expected := Lab{
		guard: Guard{direction: Down, position: [2]int{10, 7}},
		guardPositions: map[Guard]bool{
			{direction: Up, position: [2]int{6, 4}}:    true,
			{direction: Up, position: [2]int{5, 4}}:    true,
			{direction: Up, position: [2]int{4, 4}}:    true,
			{direction: Up, position: [2]int{3, 4}}:    true,
			{direction: Up, position: [2]int{2, 4}}:    true,
			{direction: Up, position: [2]int{1, 4}}:    true,
			{direction: Right, position: [2]int{1, 5}}: true,
			{direction: Right, position: [2]int{1, 6}}: true,
			{direction: Right, position: [2]int{1, 7}}: true,
			{direction: Right, position: [2]int{1, 8}}: true,
			{direction: Down, position: [2]int{2, 8}}:  true,
			{direction: Down, position: [2]int{3, 8}}:  true,
			{direction: Down, position: [2]int{4, 8}}:  true,
			{direction: Down, position: [2]int{5, 8}}:  true,
			{direction: Down, position: [2]int{6, 8}}:  true,
			{direction: Left, position: [2]int{6, 7}}:  true,
			{direction: Left, position: [2]int{6, 6}}:  true,
			{direction: Left, position: [2]int{6, 5}}:  true,
			{direction: Left, position: [2]int{6, 4}}:  true,
			{direction: Left, position: [2]int{6, 3}}:  true,
			{direction: Left, position: [2]int{6, 2}}:  true,
			{direction: Up, position: [2]int{5, 2}}:    true,
			{direction: Up, position: [2]int{4, 2}}:    true,
			{direction: Right, position: [2]int{4, 3}}: true,
			{direction: Right, position: [2]int{4, 4}}: true,
			{direction: Right, position: [2]int{4, 5}}: true,
			{direction: Right, position: [2]int{4, 6}}: true,
			{direction: Down, position: [2]int{5, 6}}:  true,
			{direction: Down, position: [2]int{6, 6}}:  true,
			{direction: Down, position: [2]int{7, 6}}:  true,
			{direction: Down, position: [2]int{8, 6}}:  true,
			{direction: Left, position: [2]int{8, 5}}:  true,
			{direction: Left, position: [2]int{8, 4}}:  true,
			{direction: Left, position: [2]int{8, 3}}:  true,
			{direction: Left, position: [2]int{8, 2}}:  true,
			{direction: Left, position: [2]int{8, 1}}:  true,
			{direction: Up, position: [2]int{7, 1}}:    true,
			{direction: Right, position: [2]int{7, 2}}: true,
			{direction: Right, position: [2]int{7, 3}}: true,
			{direction: Right, position: [2]int{7, 4}}: true,
			{direction: Right, position: [2]int{7, 5}}: true,
			{direction: Right, position: [2]int{7, 6}}: true,
			{direction: Right, position: [2]int{7, 7}}: true,
			{direction: Down, position: [2]int{8, 7}}:  true,
			{direction: Down, position: [2]int{9, 7}}:  true,
		},
		guardStartingPosition: [2]int{6, 4},
		obstacles: map[[2]int]bool{
			{0, 4}: true,
			{1, 9}: true,
			{3, 2}: true,
			{4, 7}: true,
			{6, 1}: true,
			{7, 8}: true,
			{8, 0}: true,
			{9, 6}: true,
		},
		size: [2]int{10, 10},
	}

	actual := lab
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
