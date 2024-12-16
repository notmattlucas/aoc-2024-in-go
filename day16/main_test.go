package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParts(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		maze, start, end, velo := parse(`###############
												#.......#....E#
												#.#.###.#.###.#
												#.....#.#...#.#
												#.###.#####.#.#
												#.#.#.......#.#
												#.#.#####.###.#
												#...........#.#
												###.#.#####.#.#
												#...#.....#.#.#
												#.#.#.###.#.#.#
												#.....#...#.#.#
												#.###.#.#.#.#.#
												#S..#.....#...#
												###############`)
		assert.Equal(t, 15, len(maze))
		assert.Equal(t, 15, len(maze[0]))
		assert.Equal(t, Point{1, 13}, start)
		assert.Equal(t, Point{13, 1}, end)
		assert.Equal(t, Velocity{1, 0}, velo)
	})

	t.Run("Sample 1", func(t *testing.T) {
		maze, start, end, velo := parse(`###############
												#.......#....E#
												#.#.###.#.###.#
												#.....#.#...#.#
												#.###.#####.#.#
												#.#.#.......#.#
												#.#.#####.###.#
												#...........#.#
												###.#.#####.#.#
												#...#.....#.#.#
												#.#.#.###.#.#.#
												#.....#...#.#.#
												#.###.#.#.#.#.#
												#S..#.....#...#
												###############`)
		score, path := solve(maze, start, end, velo)
		assert.Equal(t, 7036, score)
		u := uniq(path)
		assert.Equal(t, 45, len(u))
	})

	t.Run("Sample 2", func(t *testing.T) {
		maze, start, end, velo := parse(`#################
												#...#...#...#..E#
												#.#.#.#.#.#.#.#.#
												#.#.#.#...#...#.#
												#.#.#.#.###.#.#.#
												#...#.#.#.....#.#
												#.#.#.#.#.#####.#
												#.#...#.#.#.....#
												#.#.#####.#.###.#
												#.#.#.......#...#
												#.#.###.#####.###
												#.#.#...#.....#.#
												#.#.#.#####.###.#
												#.#.#.........#.#
												#.#.#.#########.#
												#S#.............#
												#################`)
		score, path := solve(maze, start, end, velo)
		assert.Equal(t, 11048, score)
		u := uniq(path)
		assert.Equal(t, 64, len(u))
	})

}
