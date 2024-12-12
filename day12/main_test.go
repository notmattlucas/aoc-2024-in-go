package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		garden := parse(`AAAA
							   BBCD
							   BBCC
							   EEEC`)
		assert.Equal(t, garden, Garden{
			{"A", "A", "A", "A"},
			{"B", "B", "C", "D"},
			{"B", "B", "C", "C"},
			{"E", "E", "E", "C"},
		})
	})

	t.Run("Find Plots", func(t *testing.T) {
		garden := parse(`AAAA
							   BBCD
							   BBCC
							   EEEC`)
		plots := findPlots(garden)
		assert.Equal(t, []Plot{
			{"A", []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
			{"B", []Point{{0, 1}, {0, 2}, {1, 1}, {1, 2}}},
			{"C", []Point{{2, 1}, {2, 2}, {3, 2}, {3, 3}}},
			{"D", []Point{{3, 1}}},
			{"E", []Point{{0, 3}, {1, 3}, {2, 3}}},
		}, plots)
	})

	t.Run("Area", func(t *testing.T) {
		garden := parse(`AAAA
							   BBCD
							   BBCC
							   EEEC`)
		a := findPlots(garden)[0]
		b := findPlots(garden)[1]
		c := findPlots(garden)[2]
		d := findPlots(garden)[3]
		e := findPlots(garden)[4]
		assert.Equal(t, a.area(), 4)
		assert.Equal(t, b.area(), 4)
		assert.Equal(t, c.area(), 4)
		assert.Equal(t, d.area(), 1)
		assert.Equal(t, e.area(), 3)
	})

	t.Run("Perimeter", func(t *testing.T) {
		garden := parse(`AAAA
							   BBCD
							   BBCC
							   EEEC`)
		a := findPlots(garden)[0]
		b := findPlots(garden)[1]
		c := findPlots(garden)[2]
		d := findPlots(garden)[3]
		e := findPlots(garden)[4]
		assert.Equal(t, a.perimeter(garden), 10)
		assert.Equal(t, b.perimeter(garden), 8)
		assert.Equal(t, c.perimeter(garden), 10)
		assert.Equal(t, d.perimeter(garden), 4)
		assert.Equal(t, e.perimeter(garden), 8)
	})

	t.Run("Part 1", func(t *testing.T) {
		garden := parse(`AAAA
							   BBCD
							   BBCC
							   EEEC`)
		plots := findPlots(garden)
		assert.Equal(t, 140, fenceCost(garden, plots))

		garden = parse(`OOOOO
								OXOXO
								OOOOO
								OXOXO
								OOOOO`)
		plots = findPlots(garden)
		assert.Equal(t, 772, fenceCost(garden, plots))

		garden = parse(`RRRRIICCFF
								RRRRIICCCF
								VVRRRCCFFF
								VVRCCCJFFF
								VVVVCJJCFE
								VVIVCCJJEE
								VVIIICJJEE
								MIIIIIJJEE
								MIIISIJEEE
								MMMISSJEEE`)
		plots = findPlots(garden)
		assert.Equal(t, 1930, fenceCost(garden, plots))
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		garden := parse(`AAAA
							   BBCD
							   BBCC
							   EEEC`)
		plots := findPlots(garden)
		assert.Equal(t, 80, fenceCostTwo(garden, plots))

		garden = parse(`OOOOO
								OXOXO
								OOOOO
								OXOXO
								OOOOO`)
		plots = findPlots(garden)
		assert.Equal(t, 436, fenceCostTwo(garden, plots))

		garden = parse(`EEEEE
								EXXXX
								EEEEE
								EXXXX
								EEEEE`)
		plots = findPlots(garden)
		assert.Equal(t, 236, fenceCostTwo(garden, plots))

		garden = parse(`AAAAAA
								AAABBA
								AAABBA
								ABBAAA
								ABBAAA
								AAAAAA`)
		plots = findPlots(garden)
		assert.Equal(t, 368, fenceCostTwo(garden, plots))

		garden = parse(`RRRRIICCFF
								RRRRIICCCF
								VVRRRCCFFF
								VVRCCCJFFF
								VVVVCJJCFE
								VVIVCCJJEE
								VVIIICJJEE
								MIIIIIJJEE
								MIIISIJEEE
								MMMISSJEEE`)
		plots = findPlots(garden)
		assert.Equal(t, 1206, fenceCostTwo(garden, plots))
	})

}
