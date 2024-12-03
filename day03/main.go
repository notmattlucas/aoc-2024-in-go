package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	ops := parse(string(b))
	fmt.Printf("Part 1: %d\n", calc(ops))
	ops = condParse(string(b))
	fmt.Printf("Part 2: %d\n", calc(ops))
}

type Mul struct {
	x int
	y int
}

func condParse(input string) []Mul {
	input = "do()" + input
	re := regexp.MustCompile(`don't\(\)`)
	segments := re.Split(input, -1)
	re = regexp.MustCompile(`do\(\)`)
	builder := ""
	for _, segment := range segments {
		dos := re.Split(segment, -1)[1:]
		for _, dos := range dos {
			builder += dos
		}
	}
	return parse(builder)
}

func calc(ops []Mul) int {
	sum := 0
	for _, op := range ops {
		sum += op.x * op.y
	}
	return sum
}

func parse(input string) []Mul {
	output := make([]Mul, 0)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	segments := re.FindAllString(input, -1)
	re = regexp.MustCompile(`(\d+)`)
	for _, segment := range segments {
		ops := re.FindAllString(segment, -1)
		x, _ := strconv.Atoi(ops[0])
		y, _ := strconv.Atoi(ops[1])
		output = append(output, Mul{x, y})
	}
	return output
}
