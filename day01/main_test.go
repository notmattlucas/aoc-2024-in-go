package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Load", func(t *testing.T) {
		input := `3   4
		 4   3
		 2   5
		 1   3
		 3   9
		 3   3`
		x, y := load(input)
		assert.Equal(t, []int{3, 4, 2, 1, 3, 3}, x)
		assert.Equal(t, []int{4, 3, 5, 3, 9, 3}, y)
	})

	t.Run("Part 1", func(t *testing.T) {
		assert.Equal(t, 11, distance(
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
		))
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {

	})

}
