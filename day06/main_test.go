package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		grid, guard := parse(`....#.....
									.........#
									..........
									..#.......
									.......#..
									..........
									.#..^.....
									........#.
									#.........
									......#...`)
		assert.Equal(t, [][]string{
			{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
			{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		}, grid)
		assert.Equal(t, Guard{4, 6, Vector{0, -1}}, guard)
	})

	t.Run("Walk", func(t *testing.T) {
		grid, guard := parse(`....#.....
									.........#
									..........
									..#.......
									.......#..
									..........
									.#..^.....
									........#.
									#.........
									......#...`)

		visited, _ := walk(grid, guard, neverHalt)
		assert.Equal(t, 41, len(visited))
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		grid, guard := parse(`....#.....
									.........#
									..........
									..#.......
									.......#..
									..........
									.#..^.....
									........#.
									#.........
									......#...`)

		visited, _ := walk(grid, guard, neverHalt)
		loops := findLoops(grid, guard, visited)
		assert.Equal(t, 6, len(loops))
	})

}
