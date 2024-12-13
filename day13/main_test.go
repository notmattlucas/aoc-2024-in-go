package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		game := parse(`Button A: X+94, Y+34
							Button B: X+22, Y+67
							Prize: X=8400, Y=5400`, 0)
		assert.Equal(t, Game{
			Vector{94, 34},
			Vector{22, 67},
			Location{8400, 5400},
		}, game[0])
	})

	t.Run("Solve", func(t *testing.T) {
		game := parse(`Button A: X+94, Y+34
							Button B: X+22, Y+67
							Prize: X=8400, Y=5400`, 0)
		assert.Equal(t, int64(280), game[0].cost())

		game = parse(`Button A: X+26, Y+66
							Button B: X+67, Y+21
							Prize: X=12748, Y=12176`, 0)
		assert.Equal(t, int64(0), game[0].cost())
		
		game = parse(`Button A: X+17, Y+86
							Button B: X+84, Y+37
							Prize: X=7870, Y=6450`, 0)
		assert.Equal(t, int64(200), game[0].cost())

		//game = parse(`Button A: X+69, Y+23
		//					Button B: X+27, Y+71
		//					Prize: X=18641, Y=10279`, 0)
		//assert.Equal(t, int64(0), game[0].cost())
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {

	})

}
