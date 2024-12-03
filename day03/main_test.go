package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		ops := parse("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
		assert.Equal(t, []Mul{
			{2, 4},
			{5, 5},
			{11, 8},
			{8, 5},
		}, ops)
	})

	t.Run("Part 1", func(t *testing.T) {
		ops := parse("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
		sum := calc(ops)
		assert.Equal(t, 161, sum)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Conditional Parse", func(t *testing.T) {
		ops := condParse("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
		assert.Equal(t, []Mul{
			{2, 4},
			{8, 5},
		}, ops)
	})

	t.Run("Part 2", func(t *testing.T) {
		ops := condParse("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
		sum := calc(ops)
		assert.Equal(t, 48, sum)
	})

}
