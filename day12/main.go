package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	garden := parse(string(b))
	plots := findPlots(garden)
	log.Printf("Part 1: %d", fenceCost(garden, plots))

	garden = parse(string(b))
	plots = findPlots(garden)
	log.Printf("Part 1: %d", fenceCostTwo(garden, plots))
}

type Garden [][]string

func (garden Garden) at(point Point) string {
	if garden.inBounds(point) {
		return garden[point.y][point.x]
	}
	return ""
}

func (garden Garden) inBounds(point Point) bool {
	return point.x >= 0 && point.x < len(garden[0]) && point.y >= 0 && point.y < len(garden)
}

func (garden Garden) neighbors(point Point) []Point {
	neighbors := make([]Point, 0)
	for _, neighbor := range point.neighbours() {
		if garden.inBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (garden Garden) plot(points []Point) [][]bool {
	shape := make([][]bool, len(garden))
	for y, row := range garden {
		shape[y] = make([]bool, len(row))
	}
	for _, point := range points {
		shape[point.y][point.x] = true
	}
	return shape
}

type Point struct {
	x, y int
}

func (point Point) neighbours() []Point {
	return []Point{
		{point.x, point.y - 1},
		{point.x, point.y + 1},
		{point.x - 1, point.y},
		{point.x + 1, point.y},
	}
}

type Plot struct {
	plant  string
	points []Point
}

func (plot Plot) edges(garden Garden) []Point {
	edges := make([]Point, 0)
	for _, point := range plot.points {
		for _, neighbor := range point.neighbours() {
			if garden.at(neighbor) != plot.plant {
				edges = append(edges, point)
			}
		}
	}
	return edges
}

func (plot Plot) area() int {
	return len(plot.points)
}

func (plot Plot) perimeter(garden Garden) int {
	edges := plot.edges(garden)
	return len(edges)
}

func (plot Plot) corners(garden Garden) int {
	up, down, left, right := Point{0, -1}, Point{0, 1}, Point{-1, 0}, Point{1, 0}
	corners := [][]Point{{up, left}, {up, right}, {down, left}, {down, right}}
	count := 0
	edges := plot.points
	for _, edge := range edges {
		for _, corner := range corners {
			first := garden.at(Point{edge.x + corner[0].x, edge.y + corner[0].y})
			second := garden.at(Point{edge.x + corner[1].x, edge.y + corner[1].y})
			third := garden.at(Point{edge.x + corner[0].x + corner[1].x, edge.y + corner[0].y + corner[1].y})

			// ■■
			// □■
			outerCorner := first != plot.plant && second != plot.plant

			// ■□
			// ■■
			innerCorner := first == plot.plant && second == plot.plant && third != plot.plant

			if outerCorner || innerCorner {
				count++
			}
		}
	}
	return count
}

func (plot Plot) fenceCost(garden Garden) int {
	return plot.area() * plot.perimeter(garden)
}

func (plot Plot) fenceCostTwo(garden Garden) int {
	return plot.area() * plot.corners(garden)
}

func fenceCost(garden Garden, plots []Plot) int {
	cost := 0
	for _, plot := range plots {
		cost += plot.fenceCost(garden)
	}
	return cost
}

func fenceCostTwo(garden Garden, plots []Plot) int {
	cost := 0
	for _, plot := range plots {
		cost += plot.fenceCostTwo(garden)
	}
	return cost
}

func findPlots(garden Garden) []Plot {
	plots := make([]Plot, 0)
	visited := make(map[Point]bool)
	for y, row := range garden {
		for x, _ := range row {
			start := Point{x, y}
			if visited[start] {
				continue
			}
			plot := findPlot(start, garden)
			for _, point := range plot.points {
				visited[point] = true
			}
			plots = append(plots, plot)
		}
	}
	return plots
}

func findPlot(root Point, garden Garden) Plot {
	todo := []Point{root}
	visited := make(map[Point]bool)
	name := garden.at(root)
	points := make([]Point, 0)
	for len(todo) > 0 {
		head, rest := todo[0], todo[1:]
		todo = rest
		if (visited)[head] {
			continue
		}
		(visited)[head] = true
		if garden.at(head) == name {
			points = append(points, head)
			todo = append(rest, garden.neighbors(head)...)
		}
	}
	return Plot{name, points}
}

func parse(input string) Garden {
	garden := make(Garden, 0)
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		row := make([]string, len(line))
		for x, plant := range line {
			row[x] = string(plant)
		}
		garden = append(garden, row)
	}
	return garden
}
