package main

import (
	"log"
	"math"
	"os"
	"strings"
)

type Cell struct {
	x, y int
}
type Vector Cell

type Span struct {
	from, to int
}

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	grid := parse(string(b))
	dim := int(math.Max(float64(len(grid)), float64(len(grid[0]))))
	log.Printf("Part 1: %d", len(antinodes(grid, Span{1, 1})))
	log.Printf("Part 2: %d", len(antinodes(grid, Span{0, dim})))
}

func antinodes(grid [][]string, span Span) map[Cell]string {
	antinodes := make(map[Cell]string)

	outOfBounds := func(x int, y int) bool {
		return x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0])
	}

	antenna := findAntenna(grid)

	for name, cells := range antenna {
		for _, source := range cells {
			for _, target := range cells {
				if source == target {
					continue
				}

				dx := target.x - source.x
				dy := target.y - source.y
				candidates := hops(source, target, Vector{dx, dy}, span, outOfBounds)

				for _, candidate := range candidates {
					antinodes[candidate] = name
				}
			}
		}
	}

	return antinodes
}

func hops(source Cell, target Cell, hop Vector, span Span, outOfBounds func(x int, y int) bool) []Cell {
	hops := make([]Cell, 0)
	for i := span.from; i <= span.to; i++ {
		hop := Cell{target.x + (i * hop.x), target.y + (i * hop.y)}
		if outOfBounds(hop.x, hop.y) {
			break
		}
		hops = append(hops, hop)
	}
	for i := span.from; i <= span.to; i++ {
		hop := Cell{source.x - (i * hop.x), source.y - (i * hop.y)}
		if outOfBounds(hop.x, hop.y) {
			break
		}
		hops = append(hops, hop)
	}
	return hops
}

func findAntenna(grid [][]string) map[string][]Cell {
	antenna := make(map[string][]Cell)
	for y, row := range grid {
		for x, cell := range row {
			if cell == "." {
				continue
			}
			if _, ok := antenna[cell]; !ok {
				antenna[cell] = []Cell{}
			}
			antenna[cell] = append(antenna[cell], Cell{x, y})
		}
	}
	return antenna
}

func parse(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for y, line := range lines {
		line = strings.TrimSpace(line)
		row := make([]string, len(line))
		for x, char := range line {
			row[x] = string(char)
		}
		grid[y] = row
	}
	return grid
}
