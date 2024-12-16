package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		_, mvs, _, _, _ := parse(`########
				#..O.O.#
				##@.O..#
				#...O..#
				#.#.O..#
				#...O..#
				#......#
				########
				
				<^^>>>vv<v>>v<<`, 1)

		assert.Equal(t, []Move{
			{-1, 0},
			{0, -1},
			{0, -1},
			{1, 0},
			{1, 0},
			{1, 0},
			{0, 1},
			{0, 1},
			{-1, 0},
			{0, 1},
			{1, 0},
			{1, 0},
			{0, 1},
			{-1, 0},
			{-1, 0},
		}, mvs)

	})

	t.Run("Part 1 Small", func(t *testing.T) {
		grid, mvs, pos, xmax, ymax := parse(`########
				#..O.O.#
				##@.O..#
				#...O..#
				#.#.O..#
				#...O..#
				#......#
				########
				
				<^^>>>vv<v>>v<<`, 1)
		run(grid, mvs, pos)
		assert.Equal(t, 7, xmax)
		assert.Equal(t, 8, ymax)
		assert.Equal(t, 2028, score(grid))
	})

	t.Run("Part 1 Large", func(t *testing.T) {
		grid, mvs, pos, _, _ := parse(`##########
									#..O..O.O#
									#......O.#
									#.OO..O.O#
									#..O@..O.#
									#O#..O...#
									#O..O..O.#
									#.OO.O.OO#
									#....O...#
									##########
									
									<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
									vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
									><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
									<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
									^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
									^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
									>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
									<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
									^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
									v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`, 1)
		run(grid, mvs, pos)
		assert.Equal(t, 10092, score(grid))
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2 Small", func(t *testing.T) {
		grid, mvs, pos, xmax, ymax := parse(`#######
									#...#.#
									#.....#
									#..OO@#
									#..O..#
									#.....#
									#######
									
									<vv<<^^<<^^`, 2)
		fmt.Println(print(grid))
		run(grid, mvs, pos)
		fmt.Println(print(grid))
		assert.Equal(t, 14, xmax)
		assert.Equal(t, 7, ymax)
		//assert.Equal(t, 105, score2(grid, xmax, ymax))
	})

	t.Run("Part 2 Large", func(t *testing.T) {
		grid, mvs, pos, xmax, ymax := parse(`##########
									#..O..O.O#
									#......O.#
									#.OO..O.O#
									#..O@..O.#
									#O#..O...#
									#O..O..O.#
									#.OO.O.OO#
									#....O...#
									##########
									
									<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
									vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
									><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
									<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
									^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
									^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
									>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
									<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
									^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
									v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`, 2)
		fmt.Println(print(grid))
		run(grid, mvs, pos)
		fmt.Println(print(grid))
		assert.Equal(t, 9021, score2(grid, xmax, ymax))
	})

}
