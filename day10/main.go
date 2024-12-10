package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	grid := parse(string(b))
	log.Printf("Part 1: %d", grid.findRoutes().score())
	log.Printf("Part 2: %d", grid.findRoutes().rating())
}

type Point struct {
	x, y int
}

type Step struct {
	point Point
	route []Point
}

type Path []Step

type Grid [][]int

type Routes map[Point][]Step

func (routes Routes) score() int {
	acc := 0
	for _, paths := range routes {
		peaks := map[Point]bool{}
		for _, path := range paths {
			peaks[path.point] = true
		}
		acc += len(peaks)
	}
	return acc
}

func (routes Routes) rating() int {
	acc := 0
	for _, paths := range routes {
		acc += len(paths)
	}
	return acc
}

func (grid Grid) findRoutes() Routes {
	routes := make(Routes)
	heads := grid.findTrailHeads()
	for _, head := range heads {
		peaks := grid.explore(head)
		routes[head] = peaks
	}
	return routes
}

func (grid Grid) explore(root Point) Path {
	peaks := make([]Step, 0)
	front := []Step{{root, []Point{root}}}
	for len(front) > 0 {
		head, rest := front[0], front[1:]
		ns := grid.neighbors(head)
		front = append(rest, ns...)
		for _, n := range ns {
			if grid[n.point.y][n.point.x] == 9 {
				peaks = append(peaks, n)
			}
		}
	}
	return peaks
}

func (grid Grid) neighbors(p Step) Path {
	moves := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	neighbors := make([]Step, 0)
	value := grid[p.point.y][p.point.x]
	for _, m := range moves {
		x := p.point.x + m.x
		y := p.point.y + m.y
		if x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid) && grid[y][x] == value+1 {
			neighbors = append(neighbors, Step{Point{x, y}, append(p.route, Point{x, y})})
		}
	}
	return neighbors
}

func (grid Grid) findTrailHeads() []Point {
	points := make([]Point, 0)
	for y, row := range grid {
		for x, cell := range row {
			if cell == 0 {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func parse(input string) Grid {
	grid := make(Grid, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, cell := range strings.TrimSpace(line) {
			if cell == '.' {
				row = append(row, -1)
				continue
			}
			num, _ := strconv.Atoi(string(cell))
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return grid
}
