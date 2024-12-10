package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		grid := parse(`...0...
						 	 ...1...
					 		 ...2...
							 6543456
							 7.....7
							 8.....8
							 9.....9`)
		assert.Equal(t, Grid{
			{-1, -1, -1, 0, -1, -1, -1},
			{-1, -1, -1, 1, -1, -1, -1},
			{-1, -1, -1, 2, -1, -1, -1},
			{6, 5, 4, 3, 4, 5, 6},
			{7, -1, -1, -1, -1, -1, 7},
			{8, -1, -1, -1, -1, -1, 8},
			{9, -1, -1, -1, -1, -1, 9},
		}, grid)
	})

	t.Run("Part 1", func(t *testing.T) {
		grid := parse(`89010123
						 	 78121874
							 87430965
							 96549874
							 45678903
							 32019012
							 01329801
							 10456732`)
		routes := grid.findRoutes().score()
		assert.Equal(t, 36, routes)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		grid := parse(`89010123
						 	 78121874
							 87430965
							 96549874
							 45678903
							 32019012
							 01329801
							 10456732`)
		routes := grid.findRoutes().rating()
		assert.Equal(t, 81, routes)
	})

}
