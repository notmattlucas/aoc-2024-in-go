package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	total  int
	opands []int
}

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	equations := parse(string(b))
	log.Printf("Part 1: %d", sumSolvable(equations, false))
	log.Printf("Part 2: %d", sumSolvable(equations, true))
}

func sumSolvable(equation []Equation, concat bool) int {
	total := 0
	for _, eq := range equation {
		if reduce(eq.opands, eq.total, concat) {
			total += eq.total
		}
	}
	return total
}

func reduce(opands []int, limit int, concat bool) bool {
	if len(opands) == 1 {
		return opands[0] == limit
	}
	if opands[0] > limit {
		return false
	}
	cat := false
	if concat && len(opands) > 1 {
		x := strconv.Itoa(opands[0])
		y := strconv.Itoa(opands[1])
		xy, _ := strconv.Atoi(x + y)
		cat = reduce(append([]int{xy}, opands[2:]...), limit, concat)
	}
	add := reduce(append([]int{opands[0] + opands[1]}, opands[2:]...), limit, concat)
	mul := reduce(append([]int{opands[0] * opands[1]}, opands[2:]...), limit, concat)
	return add || mul || cat
}

func parse(input string) []Equation {
	lines := strings.Split(input, "\n")
	equations := make([]Equation, 0)
	for _, line := range lines {
		line := strings.TrimSpace(line)
		pair := strings.Split(line, ":")
		total, _ := strconv.Atoi(strings.TrimSpace(pair[0]))
		opands := make([]int, 0)
		for _, s := range strings.Split(strings.TrimSpace(pair[1]), " ") {
			opand, _ := strconv.Atoi(strings.TrimSpace(s))
			opands = append(opands, opand)
		}
		equations = append(equations, Equation{total, opands})
	}
	return equations
}
