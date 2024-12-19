package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	patterns, designs := parse(string(b))
	count, sum := nsatisfied(designs, patterns)
	log.Printf("Part 1: %d", count)
	log.Printf("Part 2: %d", sum)
}

func nsatisfied(designs []string, patterns []string) (int, int) {
	count := 0
	sum := 0
	for _, design := range designs {
		sat := satisfied(&map[string]int{}, design, patterns)
		if sat > 0 {
			count++
		}
		sum += sat
	}
	return count, sum
}

func satisfied(cache *map[string]int, design string, patterns []string) int {
	if value, ok := (*cache)[design]; ok {
		return value
	}

	if len(design) == 0 {
		return 1
	}

	count := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			count += satisfied(cache, design[len(pattern):], patterns)
		}
	}

	(*cache)[design] = count

	return count
}

func parse(input string) ([]string, []string) {
	lines := strings.Split(input, "\n")
	patterns := strings.Split(lines[0], ", ")
	designs := make([]string, 0)
	for _, line := range lines[2:] {
		designs = append(designs, strings.TrimSpace(line))
	}
	return patterns, designs
}
