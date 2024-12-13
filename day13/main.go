package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	games := parse(string(b), 0)
	log.Printf("Part 1: %d", all(games))

	games = parse(string(b), 10000000000000)
	log.Printf("Part 2: %d", all(games))
}

type Vector struct {
	x, y int64
}
type Location Vector

type Game struct {
	a, b  Vector
	prize Location
}

func (game Game) cost() int64 {
	// Cramer's Rule - https://www.mathcentre.ac.uk/resources/Engineering%20maths%20first%20aid%20kit/latexsource%20and%20diagrams/5_2.pdf
	ax, ay, bx, by, px, py := game.a.x, game.a.y, game.b.x, game.b.y, game.prize.x, game.prize.y

	a := (px*by - bx*py) / (ax*by - bx*ay)
	b := (ax*py - px*ay) / (ax*by - bx*ay)

	// Check if the solution is valid
	if (a*ax+b*bx) != px || (a*ay+b*by != py) {
		return 0
	}

	// Calculate cost
	return 3*a + b
}

func all(games []Game) int64 {
	var acc int64 = 0
	for _, game := range games {
		acc += game.cost()
	}
	return acc
}

func parse(input string, offset int64) []Game {
	lines := strings.Split(input, "\n")
	games := make([]Game, 0)
	for i := 0; i < len(lines); i += 4 {
		a1, a2 := pair(lines[i])
		b1, b2 := pair(lines[i+1])
		p1, p2 := pair(lines[i+2])
		games = append(games, Game{
			Vector{a1, a2},
			Vector{b1, b2},
			Location{p1 + offset, p2 + offset},
		})
	}
	return games
}

func pair(input string) (int64, int64) {
	re := regexp.MustCompile(`(\d+)`)
	parts := re.FindAllString(input, 2)
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return int64(x), int64(y)
}
