package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		eqs := parse(`190: 10 19
							3267: 81 40 27
							83: 17 5
							156: 15 6
							7290: 6 8 6 15
							161011: 16 10 13
							192: 17 8 14
							21037: 9 7 18 13
							292: 11 6 16 20`)
		assert.Equal(t, []Equation{
			{190, []int{10, 19}},
			{3267, []int{81, 40, 27}},
			{83, []int{17, 5}},
			{156, []int{15, 6}},
			{7290, []int{6, 8, 6, 15}},
			{161011, []int{16, 10, 13}},
			{192, []int{17, 8, 14}},
			{21037, []int{9, 7, 18, 13}},
			{292, []int{11, 6, 16, 20}},
		}, eqs)
	})

	t.Run("Part 1", func(t *testing.T) {
		eqs := parse(`190: 10 19
							3267: 81 40 27
							83: 17 5
							156: 15 6
							7290: 6 8 6 15
							161011: 16 10 13
							192: 17 8 14
							21037: 9 7 18 13
							292: 11 6 16 20`)
		sum := sumSolvable(eqs, false)
		assert.Equal(t, 3749, sum)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		eqs := parse(`190: 10 19
							3267: 81 40 27
							83: 17 5
							156: 15 6
							7290: 6 8 6 15
							161011: 16 10 13
							192: 17 8 14
							21037: 9 7 18 13
							292: 11 6 16 20`)
		sum := sumSolvable(eqs, true)
		assert.Equal(t, 11387, sum)
	})

}
