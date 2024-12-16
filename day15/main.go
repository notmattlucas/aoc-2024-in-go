package main

import (
	collections "github.com/notmattlucas/aoc-2024-in-go"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	objects, mvs, robot, xmax, ymax := parse(string(b), 1)
	run(objects, mvs, robot)
	// 1487337
	log.Printf("Part 1: %d", score(objects))

	objects, mvs, robot, xmax, ymax = parse(string(b), 2)
	run(objects, mvs, robot)
	log.Printf("Part 2: %d", score2(objects, xmax, ymax))
}

type Move struct {
	x, y int
}
type Point = Move

var moves = map[string]Move{
	"^": {0, -1},
	"v": {0, 1},
	"<": {-1, 0},
	">": {1, 0},
}

type Object struct {
	points []Point
	icon   string
}

func search(objects []*Object, x, y int) *Object {
	for _, obj := range objects {
		if obj.at(x, y) {
			return obj
		}
	}
	return nil
}

func (o *Object) at(x, y int) bool {
	b, _ := collections.Any(o.points, func(p Point) bool {
		return p.x == x && p.y == y
	})
	return b
}

func (o *Object) closest(xmax int, ymax int) (int, int) {
	xs := make([]int, 0)
	ys := make([]int, 0)
	for _, p := range o.points {
		xs = append(xs, p.x)
		ys = append(ys, p.y)
	}
	xmn := slices.Min(xs)
	xmx := slices.Min(xs)
	ymn := slices.Min(ys)
	ymx := slices.Min(ys)
	yscore := math.Min(float64(ymax-ymx), float64(ymn))
	xscore := math.Min(float64(xmax-xmx), float64(xmn))
	return int(xscore), int(yscore)
}

func (o *Object) movable(objects []*Object, velocity Move) bool {
	next := o.findNext(objects, velocity)

	if o.icon == "#" {
		return false
	}

	if len(next) == 0 {
		return true
	}

	b, _ := collections.All(next, func(next *Object) bool {
		return next.movable(objects, velocity)
	})

	return b
}

func (o *Object) findNext(objects []*Object, velocity Move) []*Object {
	next := make([]*Object, 0)
	for _, p := range o.points {
		nx := p.x + velocity.x
		ny := p.y + velocity.y
		object := search(objects, nx, ny)
		if object != nil && object != o {
			next = append(next, object)
		}

	}
	return next
}

func (o *Object) move(objects []*Object, velocity Move) {
	if !o.movable(objects, velocity) {
		return
	}

	next := o.findNext(objects, velocity)
	for _, nx := range next {
		nx.move(objects, velocity)
	}

	npx := make([]Point, 0)
	for _, p := range o.points {
		npx = append(npx, Point{p.x + velocity.x, p.y + velocity.y})
	}
	o.points = npx
}

func run(objects []*Object, mvs []Move, robot *Object) {
	for _, mv := range mvs {
		(*robot).move(objects, mv)
	}
}

func score(objects []*Object) int {
	total := 0
	for _, obj := range objects {
		if obj.icon == "O" {
			for _, p := range obj.points {
				total += p.y*100 + p.x
			}
		}
	}
	return total
}

func score2(objects []*Object, xmax int, ymax int) int {
	total := 0
	for _, obj := range objects {
		if obj.icon == "O" {
			x, y := obj.closest(xmax, ymax)
			total += y*100 + x
		}
	}
	return total
}

func parse(input string, width int) ([]*Object, []Move, *Object, int, int) {
	lines := strings.Split(input, "\n")
	var robot Object
	var objects []*Object
	var mvs []Move
	ymax := 0
	xmax := 0
	for y, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "<") || strings.HasPrefix(line, ">") || strings.HasPrefix(line, "^") || strings.HasPrefix(line, "v") {
			for _, c := range line {
				mvs = append(mvs, moves[string(c)])
			}
		} else {
			x := 0
			ymax += 1
			for _, c := range strings.Split(line, "") {

				points := make([]Point, 0)
				for i := 0; i < width; i++ {
					points = append(points, Point{x + i, y})
				}

				switch c {
				case "#":
					objects = append(objects, &Object{points, c})
				case "O":
					objects = append(objects, &Object{points, c})
				case "@":
					{
						robot = Object{[]Point{{x, y}}, c}
						objects = append(objects, &robot)
					}
				}

				xmax = x
				x += width
			}
		}
	}
	return objects, mvs, &robot, xmax, ymax
}

func print(objects []*Object) string {
	xmax := 0
	ymax := 0
	for _, object := range objects {
		for _, point := range object.points {
			xmax = int(math.Max(float64(xmax), float64(point.x)))
			ymax = int(math.Max(float64(ymax), float64(point.y)))
		}
	}
	grid := make([][]string, ymax+1)
	for y := 0; y <= ymax; y++ {
		grid[y] = make([]string, xmax+1)
		for x := 0; x <= xmax; x++ {
			grid[y][x] = "."
			for _, object := range objects {
				if object.at(x, y) {
					grid[y][x] = object.icon
				}
			}
		}
	}

	s := ""
	for y := 0; y <= ymax; y++ {
		for x := 0; x <= xmax; x++ {
			s += grid[y][x]
		}
		s += "\n"
	}
	return s
}
