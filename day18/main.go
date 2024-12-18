package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	maze, corruptions := parse(string(b), 71, 71)
	maze = corrupt(maze, corruptions, 1024)
	path := solve(maze)
	log.Printf("Part 1: %d", len(path))
	pt := search(maze, corruptions)
	log.Printf("Part 2: %d", pt)
}

type Maze = [][]string

type Point struct {
	x, y int
}

type State struct {
	pos  Point
	path []Point
}

func ok(maze Maze, p Point) bool {
	return p.y >= 0 && p.y < len(maze) && p.x >= 0 && p.x < len(maze[0]) && maze[p.y][p.x] != "#"
}

func (state State) move(maze Maze) []State {
	states := make([]State, 0)
	for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		fwd := Point{state.pos.x + dir.x, state.pos.y + dir.y}
		if ok(maze, fwd) {
			newpath := make([]Point, len(state.path))
			copy(newpath, state.path)
			states = append(states, State{fwd, append(newpath, state.pos)})
		}
	}
	return states
}

// Could've done binary search, but problem space is small enough to just brute force it, and I can't be bothered!
func search(maze Maze, corruptions []Point) Point {
	var prevent Point
	for i := 0; i < len(corruptions); i++ {
		subset := corruptions[:i]
		if solve(corrupt(maze, subset, i)) == nil {
			prevent = corruptions[i-1]
			break
		}
	}
	return prevent
}

func solve(maze Maze) []Point {

	end := Point{len(maze[0]) - 1, len(maze) - 1}
	pending := []State{
		{Point{0, 0}, make([]Point, 0)},
	}

	visited := map[Point]bool{}

	for len(pending) > 0 {

		head := pending[0]
		pending = pending[1:]

		if visited[head.pos] {
			continue
		}
		visited[head.pos] = true

		if head.pos == end {
			return head.path
		}

		for _, state := range head.move(maze) {
			pending = append(pending, state)
		}

	}

	// no solution
	return nil
}

func corrupt(maze Maze, corruptions []Point, n int) Maze {
	// reset maze
	for y, _ := range maze {
		for x, _ := range maze[y] {
			maze[y][x] = "."
		}
	}
	// apply corruptions
	for corrupted, c := range corruptions {
		if corrupted >= n {
			break
		}
		maze[c.y][c.x] = "#"
	}
	return maze
}

func parse(input string, xsize int, ysize int) (Maze, []Point) {
	maze := make(Maze, ysize)
	for y, _ := range maze {
		maze[y] = make([]string, xsize)
		for x, _ := range maze[y] {
			maze[y][x] = "."
		}
	}

	locations := make([]Point, 0)
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		spl := strings.Split(line, ",")
		x, _ := strconv.Atoi(spl[0])
		y, _ := strconv.Atoi(spl[1])
		locations = append(locations, Point{x, y})
	}

	return maze, locations
}
