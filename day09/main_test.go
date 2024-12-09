package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		segments := parse("2333133121414131402")
		defrag := defragment(segments, func(from Segment, candidate Segment) bool {
			return candidate.id == -1
		})
		assert.Equal(t, 1928, checksum(defrag))
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		segments := parse("2333133121414131402")
		defrag := defragment(segments, func(from Segment, candidate Segment) bool {
			return candidate.id == -1 && candidate.count >= from.count
		})
		assert.Equal(t, 2858, checksum(defrag))
	})

}
