package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Printf("Part 1: %d", sum("8793800 1629 65 5 960 0 138983 85629", 25, &Cache{}))
	log.Printf("Part 1: %d", sum("8793800 1629 65 5 960 0 138983 85629", 75, &Cache{}))
}

type Entry struct {
	pebble    int
	remaining int
}
type Cache map[Entry]int

func sum(pebbles string, remaining int, cache *Cache) int {
	total := 0
	for _, pebble := range strings.Split(pebbles, " ") {
		i, _ := strconv.Atoi(pebble)
		total += blink(i, remaining, cache)
	}
	return total
}

func blink(pebble int, remaining int, cache *Cache) int {
	if remaining == 0 {
		return 1
	}

	k := Entry{pebble, remaining}
	if value, ok := (*cache)[k]; ok {
		return value
	}

	// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
	if pebble == 0 {
		result := blink(1, remaining-1, cache)
		(*cache)[k] = result
		return result
	}

	// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
	// The left half of the digits are engraved on the new left stone, and the right half of the digits are
	// engraved on the new right stone.
	// (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
	s := strconv.Itoa(pebble)
	l := len(s)
	if l%2 == 0 {
		first, _ := strconv.Atoi(s[:l/2])
		second, _ := strconv.Atoi(s[l/2:])
		result := blink(first, remaining-1, cache) + blink(second, remaining-1, cache)
		(*cache)[k] = result
		return result
	}

	// If none of the other rules apply, the stone is replaced by a new stone;
	// the old stone's number multiplied by 2024 is engraved on the new stone.
	result := blink(pebble*2024, remaining-1, cache)

	(*cache)[k] = result
	return result
}
