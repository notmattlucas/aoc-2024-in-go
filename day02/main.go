package main

import (
	collections "github.com/notmattlucas/aoc-2024-in-go"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Record = []int

type Direction = int

const (
	Increasing Direction = iota
	Decreasing
)

type Adjacent = [2]int

type Movement struct {
	direction Direction
	amount    int
	allowed   bool
}

func main() {
	b, err := os.ReadFile("./part01_input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	records := load(string(b))
	log.Printf("Part 1: %d", countSafe(records, 0))
	log.Printf("Part 2: %d", countSafe(records, 1))
}

func load(input string) []Record {
	var records = make([]Record, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		record := make(Record, 0)
		for _, field := range fields {
			val, _ := strconv.Atoi(field)
			record = append(record, val)
		}
		records = append(records, record)
	}
	return records
}

func countSafe(records []Record, dampen int) int {
	acc := 0
	for _, record := range records {
		if anySafe(record, dampen) {
			acc++
		}
	}
	return acc
}

func anySafe(record Record, dampen int) bool {
	possibilities := []Record{record}
	for remove := 1; remove <= dampen; remove++ {
		for i := 0; i < len(record); i++ {
			clone := make(Record, len(record))
			copy(clone, record)
			clone = append(clone[:i], clone[i+1:]...)
			possibilities = append(possibilities, clone)
		}
	}

	safe, _ := collections.Any(possibilities, func(attempt Record) bool { return isSafe(attempt) })
	return safe
}

func isSafe(record Record) bool {
	pairs := make([]Adjacent, 0)
	for i := 0; i < len(record)-1; i++ {
		pairs = append(pairs, Adjacent{record[i], record[i+1]})
	}

	movements := make([]Movement, 0)
	for _, pair := range pairs {
		movements = append(movements, movement(pair))
	}

	increasing, _ := collections.All(movements, func(move Movement) bool {
		return move.direction == Increasing && move.allowed
	})
	decreasing, _ := collections.All(movements, func(move Movement) bool {
		return move.direction == Decreasing && move.allowed
	})
	return increasing || decreasing
}

func movement(pair Adjacent) Movement {
	step := int(math.Abs(float64(pair[0] - pair[1])))
	allowed := step >= 1 && step <= 3
	if pair[0] < pair[1] {
		return Movement{Increasing, pair[1] - pair[0], allowed}
	}
	return Movement{Decreasing, pair[0] - pair[1], allowed}
}
