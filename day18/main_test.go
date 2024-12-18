package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		maze, corruptions := parse(`5,4
					4,2
					4,5
					3,0
					2,1
					6,3
					2,4
					1,5
					0,6
					3,3
					2,6
					5,1
					1,2
					5,5
					2,5
					6,5
					1,4
					0,4
					6,4
					1,1
					6,1
					1,0
					0,5
					1,6
					2,0`, 7, 7)
		assert.Equal(t, Maze{
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
		}, maze)
		assert.Equal(t, []Point{
			{5, 4},
			{4, 2},
			{4, 5},
			{3, 0},
			{2, 1},
			{6, 3},
			{2, 4},
			{1, 5},
			{0, 6},
			{3, 3},
			{2, 6},
			{5, 1},
			{1, 2},
			{5, 5},
			{2, 5},
			{6, 5},
			{1, 4},
			{0, 4},
			{6, 4},
			{1, 1},
			{6, 1},
			{1, 0},
			{0, 5},
			{1, 6},
			{2, 0},
		}, corruptions)
	})

	t.Run("Corrupt", func(t *testing.T) {
		maze, corruptions := parse(`5,4
					4,2
					4,5
					3,0
					2,1
					6,3
					2,4
					1,5
					0,6
					3,3
					2,6
					5,1
					1,2
					5,5
					2,5
					6,5
					1,4
					0,4
					6,4
					1,1
					6,1
					1,0
					0,5
					1,6
					2,0`, 7, 7)
		maze = corrupt(maze, corruptions, 12)
		assert.Equal(t, Maze{
			{".", ".", ".", "#", ".", ".", "."},
			{".", ".", "#", ".", ".", "#", "."},
			{".", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", "#", ".", ".", "#"},
			{".", ".", "#", ".", ".", "#", "."},
			{".", "#", ".", ".", "#", ".", "."},
			{"#", ".", "#", ".", ".", ".", "."},
		}, maze)
	})

	t.Run("Solve", func(t *testing.T) {
		maze, corruptions := parse(`5,4
					4,2
					4,5
					3,0
					2,1
					6,3
					2,4
					1,5
					0,6
					3,3
					2,6
					5,1
					1,2
					5,5
					2,5
					6,5
					1,4
					0,4
					6,4
					1,1
					6,1
					1,0
					0,5
					1,6
					2,0`, 7, 7)
		maze = corrupt(maze, corruptions, 12)
		path := solve(maze)
		assert.Equal(t, 22, len(path))
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		maze, corruptions := parse(`5,4
					4,2
					4,5
					3,0
					2,1
					6,3
					2,4
					1,5
					0,6
					3,3
					2,6
					5,1
					1,2
					5,5
					2,5
					6,5
					1,4
					0,4
					6,4
					1,1
					6,1
					1,0
					0,5
					1,6
					2,0`, 7, 7)
		pt := search(maze, corruptions)
		assert.Equal(t, Point{6, 1}, pt)
	})

}
