package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part1_input.txt")
	xs, ys := load(string(b))
	dist := distance(xs, ys)
	log.Printf("Part 1: %d", dist)
	sim := similarity(xs, ys)
	log.Printf("Part 2: %d", sim)
}

func load(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	length := len(lines)
	xs := make([]int, length)
	ys := make([]int, length)
	for idx, line := range lines {
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		xs[idx], _ = strconv.Atoi(fields[0])
		ys[idx], _ = strconv.Atoi(fields[1])
	}
	return xs, ys
}

func distance(xs []int, ys []int) int {
	sort.Ints(xs)
	sort.Ints(ys)
	var acc float64 = 0
	for i := 0; i < len(xs); i++ {
		acc += math.Abs(float64(xs[i] - ys[i]))
	}
	return int(acc)
}

func similarity(xs []int, ys []int) int {
	ycount := make(map[int]int)
	for _, y := range ys {
		ycount[y] += 1
	}

	acc := 0
	for _, x := range xs {
		acc += x * ycount[x]
	}

	return acc
}
