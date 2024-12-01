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
	result := distance(load(string(b)))
	log.Printf("Part 1: %d", result)
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
