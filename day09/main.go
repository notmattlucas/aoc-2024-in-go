package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	segments := parse(string(b))
	defrag := defragment(segments, func(from Segment, candidate Segment) bool {
		return candidate.id == -1
	})
	log.Printf("Part 1: %d", checksum(defrag))

	segments = parse(string(b))
	defrag = defragment(segments, func(from Segment, candidate Segment) bool {
		return candidate.id == -1 && candidate.count >= from.count
	})
	log.Printf("Part 2: %d", checksum(defrag))
}

type Segment struct {
	id, count int
}

func parse(input string) []Segment {
	segments := make([]Segment, 0)
	for idx, count := range input {
		count, _ := strconv.Atoi(string(count))
		if idx%2 == 0 {
			segments = append(segments, Segment{idx / 2, count})
		} else {
			segments = append(segments, Segment{-1, count})
		}
	}
	return segments
}

func checksum(disk []Segment) int {
	sum := 0
	idx := 0
	for _, segment := range disk {
		for j := 0; j < segment.count; j++ {
			if segment.id != -1 {
				sum += segment.id * idx
			}
			idx++
		}
	}
	return sum
}

func defragment(disk []Segment, condition func(Segment, Segment) bool) []Segment {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i].id == -1 {
			continue
		}
		empty := find(&disk, condition, disk[i])
		if empty != -1 && empty < i {
			i = move(&disk, i, empty)
		}
	}
	return disk
}

func split(segments *[]Segment, i int, count int) {
	add := []Segment{
		{(*segments)[i].id, count},
		{(*segments)[i].id, (*segments)[i].count - count},
	}
	*segments = append((*segments)[:i], append(add, (*segments)[i+1:]...)...)
}

func move(segments *[]Segment, from int, to int) int {
	// If more space in target, split target into suitable size
	if (*segments)[to].count > (*segments)[from].count {
		split(segments, to, (*segments)[from].count)
		// Bump from because array is bigger
		from++
	}
	// If more data in source, split source into suitable size for transfer
	if (*segments)[from].count > (*segments)[to].count {
		split(segments, from, (*segments)[from].count-(*segments)[to].count)
		// Bump from because array is bigger
		from++
	}
	tmp := (*segments)[to].id
	(*segments)[to].id = (*segments)[from].id
	(*segments)[from].id = tmp
	return from
}

func find(segments *[]Segment, condition func(Segment, Segment) bool, from Segment) int {
	for i, candidate := range *segments {
		if condition(from, candidate) {
			return i
		}
	}
	return -1
}
