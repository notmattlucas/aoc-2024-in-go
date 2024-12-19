package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {
		patterns, designs := parse(`r, wr, b, g, bwu, rb, gb, br

					brwrr
					bggr
					gbbr
					rrbgbr
					ubwu
					bwurrg
					brgr
					bbrgwb`)

		assert.Equal(t, []string{
			"r", "wr", "b", "g", "bwu", "rb", "gb", "br",
		}, patterns)
		assert.Equal(t, []string{
			"brwrr",
			"bggr",
			"gbbr",
			"rrbgbr",
			"ubwu",
			"bwurrg",
			"brgr",
			"bbrgwb",
		}, designs)
	})

	t.Run("Part 1", func(t *testing.T) {
		patterns, designs := parse(`r, wr, b, g, bwu, rb, gb, br

					brwrr
					bggr
					gbbr
					rrbgbr
					ubwu
					bwurrg
					brgr
					bbrgwb`)

		count, _ := nsatisfied(designs, patterns)
		assert.Equal(t, 6, count)
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		patterns, designs := parse(`r, wr, b, g, bwu, rb, gb, br

					brwrr
					bggr
					gbbr
					rrbgbr
					ubwu
					bwurrg
					brgr
					bbrgwb`)

		_, sum := nsatisfied(designs, patterns)
		assert.Equal(t, 16, sum)
	})

}
