package main

import (
	collections "github.com/notmattlucas/aoc-2024-in-go"
	"log"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	grid, _, _ := parse(string(b))
	log.Printf("Part 1: %d", wordSearch("XMAS", grid))
	log.Printf("Part 2: %d", masSearch(grid))
}

type Point struct {
	x, y int
}

func parse(input string) ([][]string, int, int) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	width, height := 0, len(lines)
	mapped := make([][]string, height)
	for y, line := range lines {
		line = strings.TrimSpace(line)
		width = len(line)
		mapped[y] = make([]string, width)
		for x, char := range line {
			mapped[y][x] = string(char)
		}
	}
	return mapped, width, height
}

func masSearch(grid [][]string) int {
	points := pos("A", grid)
	found := 0
	for _, point := range points {
		ws := centreOf(point, grid)
		mas := collections.Count(ws, func(word string) bool {
			return word == "MAS"
		})
		if mas == 2 {
			found++
		}
	}
	return found
}

func wordSearch(find string, grid [][]string) int {
	points := pos(find[0:1], grid)
	found := 0
	for _, point := range points {
		ws := startingWith(point, grid, len(find))
		found += collections.Count(ws, func(word string) bool {
			return word == find
		})
	}
	return found
}

func pos(char string, grid [][]string) []Point {
	points := make([]Point, 0)
	for y, row := range grid {
		for x, cell := range row {
			if cell == char {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func startingWith(point Point, grid [][]string, size int) []string {
	return words(point, grid, [][]Point{
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		{{0, 0}, {0, -1}, {0, -2}, {0, -3}},
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}},
		{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}},
		{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
		{{0, 0}, {1, -1}, {2, -2}, {3, -3}},
		{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}},
	}, size)
}

func centreOf(point Point, grid [][]string) []string {
	return words(point, grid, [][]Point{
		{{-1, -1}, {0, 0}, {1, 1}},
		{{1, 1}, {0, 0}, {-1, -1}},
		{{1, -1}, {0, 0}, {-1, 1}},
		{{-1, 1}, {0, 0}, {1, -1}},
	}, 3)
}

func words(point Point, grid [][]string, transformations [][]Point, size int) []string {
	words := make([]string, 0)

	for _, transformation := range transformations {
		var sb strings.Builder
		for _, tx := range transformation {
			x := point.x + tx.x
			y := point.y + tx.y
			if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
				continue
			}
			sb.WriteString(grid[y][x])
		}
		if sb.Len() == size {
			words = append(words, sb.String())
		}
	}

	return words
}
