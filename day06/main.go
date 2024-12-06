package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	grid, guard := parse(string(b))
	visited, _ := walk(grid, guard, neverHalt)
	log.Printf("Part 1: %d", len(visited))
	loops := findLoops(grid, guard, visited)
	log.Printf("Part 2: %d", len(loops))
}

// Vector Represents direction of travel of the guard
type Vector struct {
	x int
	y int
}

// Cell Point in the grid
type Cell Vector

// Guard has a position and a direction of travel
type Guard struct {
	x      int
	y      int
	vector Vector
}

// Vectors Map of icon to direction
var Vectors = map[string]Vector{
	"^": {0, -1},
	"v": {0, 1},
	"<": {-1, 0},
	">": {1, 0},
}

// Given a grid and an initial guard position, walk the grid until a halt condition is met or the guard leaves the grid
// Returns the visited cells and a boolean indicating if the function was halted
func walk(grid [][]string, guard Guard, halt func(guard Guard, visited map[Cell][]Vector) bool) (map[Cell][]Vector, bool) {
	visited := map[Cell][]Vector{}
	for {
		if halt(guard, visited) {
			return visited, true
		}

		cellVisited, ok := visited[Cell{guard.x, guard.y}]
		if !ok {
			cellVisited = make([]Vector, 0)
		}

		visited[Cell{guard.x, guard.y}] = append(cellVisited, guard.vector)
		next := Cell{guard.x + guard.vector.x, guard.y + guard.vector.y}

		if !inRoom(grid, next) {
			break
		}

		if isObstacle(grid, next) {
			turn(&guard)
		} else {
			move(&guard, next)
		}
	}

	return visited, false
}

// Given a grid, guard and a path, find all the points where we could place an obstacle to put the guard in a loop
func findLoops(grid [][]string, guard Guard, visited map[Cell][]Vector) map[Cell]bool {
	loops := map[Cell]bool{}
	for cell, _ := range visited {
		blocked := gridWithObstacle(grid, cell)
		_, halted := walk(blocked, guard, haltOnRetread)
		if halted {
			loops[cell] = true
		}
	}
	return loops
}

// Create a new instance of the grid with an obstacle at the given cell
func gridWithObstacle(grid [][]string, cell Cell) [][]string {
	duplicate := make([][]string, len(grid))
	for i := range grid {
		duplicate[i] = make([]string, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	duplicate[cell.y][cell.x] = "#"
	return duplicate
}

func isObstacle(grid [][]string, cell Cell) bool {
	return grid[cell.y][cell.x] == "#"
}

// Turns the guard 90 degrees clockwise
func turn(guard *Guard) {
	guard.vector = Vector{-guard.vector.y, guard.vector.x}
}

// Moves the guard given its current position and vector of travel
func move(guard *Guard, cell Cell) {
	guard.x = cell.x
	guard.y = cell.y
}

// Returns true if the cell is within the grid
func inRoom(grid [][]string, cell Cell) bool {
	ymax := len(grid)
	xmax := len(grid[0])
	return cell.x >= 0 && cell.x < xmax && cell.y >= 0 && cell.y < ymax
}

// Halt function that never halts
func neverHalt(_ Guard, _ map[Cell][]Vector) bool {
	return false
}

// Halt function that halts if the guard has been in the cell before, travelling in the same direction.
// Since the guard's behaviour is deterministic, this means the guard is in a loop
func haltOnRetread(guard Guard, visited map[Cell][]Vector) bool {
	value, found := visited[Cell{guard.x, guard.y}]
	if !found {
		return false
	}
	for _, vector := range value {
		if vector == guard.vector {
			return true
		}
	}
	return false
}

func parse(input string) ([][]string, Guard) {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	var guard Guard
	for y, line := range lines {
		line = strings.TrimSpace(line)
		row := make([]string, len(line))
		for x, char := range line {
			if char == '#' || char == '.' {
				row[x] = string(char)
			} else {
				guard = Guard{x, y, Vectors[string(char)]}
				row[x] = "."
			}
		}
		grid[y] = row
	}
	return grid, guard
}
