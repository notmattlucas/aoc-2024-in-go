package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		robots, room := parse(`p=0,4 v=3,-3
					p=6,3 v=-1,-3
					p=10,3 v=-1,2
					p=2,0 v=2,-1
					p=0,0 v=1,3
					p=3,0 v=-2,-2
					p=7,6 v=-1,-3
					p=3,0 v=-1,-2
					p=9,3 v=2,3
					p=7,3 v=-1,2
					p=2,4 v=2,-3
					p=9,5 v=-3,-3`, 11, 7)
		assert.Equal(t, []Robot{
			{Position{0, 4}, Velocity{3, -3}},
			{Position{6, 3}, Velocity{-1, -3}},
			{Position{10, 3}, Velocity{-1, 2}},
			{Position{2, 0}, Velocity{2, -1}},
			{Position{0, 0}, Velocity{1, 3}},
			{Position{3, 0}, Velocity{-2, -2}},
			{Position{7, 6}, Velocity{-1, -3}},
			{Position{3, 0}, Velocity{-1, -2}},
			{Position{9, 3}, Velocity{2, 3}},
			{Position{7, 3}, Velocity{-1, 2}},
			{Position{2, 4}, Velocity{2, -3}},
			{Position{9, 5}, Velocity{-3, -3}},
		}, robots)
		assert.Equal(t, Room{11, 7}, room)
	})

	t.Run("Move", func(t *testing.T) {
		room := Room{11, 7}
		robot := Robot{Position{2, 4}, Velocity{2, -3}}

		assert.Equal(t, Robot{Position{4, 1}, Velocity{2, -3}}, robot.move(room, 1))

		assert.Equal(t, Robot{Position{6, 5}, Velocity{2, -3}}, robot.move(room, 2))

		assert.Equal(t, Robot{Position{8, 2}, Velocity{2, -3}}, robot.move(room, 3))

		assert.Equal(t, Robot{Position{10, 6}, Velocity{2, -3}}, robot.move(room, 4))

		assert.Equal(t, Robot{Position{1, 3}, Velocity{2, -3}}, robot.move(room, 5))
	})

	t.Run("Part 01", func(t *testing.T) {
		robots, room := parse(`p=0,4 v=3,-3
					p=6,3 v=-1,-3
					p=10,3 v=-1,2
					p=2,0 v=2,-1
					p=0,0 v=1,3
					p=3,0 v=-2,-2
					p=7,6 v=-1,-3
					p=3,0 v=-1,-2
					p=9,3 v=2,3
					p=7,3 v=-1,2
					p=2,4 v=2,-3
					p=9,5 v=-3,-3`, 11, 7)
		robots = move(robots, room, 100)
		assert.Equal(t, 12, quadCount(robots, room))
	})

}
