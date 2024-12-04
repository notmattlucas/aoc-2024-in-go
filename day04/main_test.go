package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Char Position", func(t *testing.T) {
		grid, width, height := parse(`MMMSXXMASM
					MSAMXMSMSA
					AMXSXMAAMM
					MSAMASMSMX
					XMASAMXAMM
					XXAMMXXAMA
					SMSMSASXSS
					SAXAMASAAA
					MAMMMXMMMM
					MXMXAXMASX`)
		pos := pos("X", grid)
		assert.Equal(t, []Point{{4, 0}, {5, 0}, {4, 1}, {2, 2}, {4, 2}, {9, 3}, {0, 4}, {6, 4}, {0, 5}, {1, 5}, {5, 5}, {6, 5}, {7, 6}, {2, 7}, {5, 8}, {1, 9}, {3, 9}, {5, 9}, {9, 9}}, pos)
		assert.Equal(t, 10, width)
		assert.Equal(t, 10, height)
	})

	t.Run("Words", func(t *testing.T) {
		grid, _, _ := parse(`MMMSXXMASM
					MSAMXMSMSA
					AMXSXMAAMM
					MSAMASMSMX
					XMASAMXAMM
					XXAMMXXAMA
					SMSMSASXSS
					SAXAMASAAA
					MAMMMXMMMM
					MXMXAXMASX`)
		ws := startingWith(Point{6, 4}, grid, 4)
		assert.Equal(t, []string{"XXSS", "XMAS", "XAMM", "XMAS", "XSXM", "XASA", "XSMA", "XXSA"}, ws)
	})

	t.Run("Word Search", func(t *testing.T) {
		grid, _, _ := parse(`MMMSXXMASM
					MSAMXMSMSA
					AMXSXMAAMM
					MSAMASMSMX
					XMASAMXAMM
					XXAMMXXAMA
					SMSMSASXSS
					SAXAMASAAA
					MAMMMXMMMM
					MXMXAXMASX`)
		assert.Equal(t, 18, wordSearch("XMAS", grid))
	})
}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		t.Run("MAS Search", func(t *testing.T) {
			grid, _, _ := parse(`MMMSXXMASM
					MSAMXMSMSA
					AMXSXMAAMM
					MSAMASMSMX
					XMASAMXAMM
					XXAMMXXAMA
					SMSMSASXSS
					SAXAMASAAA
					MAMMMXMMMM
					MXMXAXMASX`)
			assert.Equal(t, 9, masSearch(grid))
		})
	})

}
