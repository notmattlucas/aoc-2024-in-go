package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	maze, start, end, velo := parse(string(b))
	score, path := solve(maze, start, end, velo)
	log.Printf("Part 1: %d", score)
	log.Printf("Part 2: %d", len(uniq(path)))
}

type Maze = [][]string

type Velocity struct {
	x, y int
}

type Point struct {
	x, y int
}

type State struct {
	pos  Point
	velo Velocity
	cost int
	path []Point
}

type Visit struct {
	pos  Point
	velo Velocity
}

func ok(maze Maze, p Point) bool {
	return maze[p.y][p.x] != "#"
}

func (state State) move(maze Maze) []State {
	states := make([]State, 0)
	fwd := Point{state.pos.x + state.velo.x, state.pos.y + state.velo.y}
	if ok(maze, fwd) {
		newpath := make([]Point, len(state.path))
		copy(newpath, state.path)
		states = append(states, State{fwd, state.velo, state.cost + 1, append(newpath, state.pos)})
	}
	states = append(states, State{state.pos, Velocity{-state.velo.y, state.velo.x}, state.cost + 1000, state.path})
	states = append(states, State{state.pos, Velocity{state.velo.y, -state.velo.x}, state.cost + 1000, state.path})
	return states
}

func (state State) visit() Visit {
	return Visit{state.pos, state.velo}
}

func uniq(path []Point) map[Point]bool {
	seen := map[Point]bool{}
	for _, p := range path {
		seen[p] = true
	}
	return seen
}

func solve(maze Maze, start Point, end Point, initial Velocity) (int, []Point) {

	progress := map[Visit]int{}
	score := math.MaxInt32
	visited := map[int][]Point{}

	pending := []State{
		{start, initial, 0, make([]Point, 0)},
	}

	for len(pending) > 0 {

		head := pending[0]
		pending = pending[1:]

		// if at end, update score if it was better than previous
		if head.pos == end {
			if head.cost != 0 && head.cost <= score {
				score = head.cost
				if _, ok := visited[score]; !ok {
					visited[score] = head.path
				}
				visited[score] = slices.Concat(visited[score], head.path, []Point{head.pos})
			}
			continue
		}

		// if we've already had a better score going the same direction, skip
		prior := progress[head.visit()]
		if prior != 0 && prior < head.cost {
			continue
		}

		// update progress and add new states to pending
		progress[head.visit()] = head.cost
		for _, state := range head.move(maze) {
			pending = append(pending, state)
		}

	}

	return score, visited[score]
}

func parse(input string) (Maze, Point, Point, Velocity) {
	maze := make(Maze, 0)
	var start, end Point
	for y, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		row := make([]string, 0)
		for x, char := range strings.Split(line, "") {
			row = append(row, string(char))
			if char == "S" {
				start = Point{x, y}
			}
			if char == "E" {
				end = Point{x, y}
			}
		}
		maze = append(maze, row)
	}
	return maze, start, end, Velocity{1, 0}
}
