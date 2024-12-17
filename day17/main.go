package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	cmp := parse(`Register A: 51571418
						Register B: 0
						Register C: 0
						
						Program: 2,4,1,1,7,5,0,3,1,4,4,5,5,5,3,0
						`)
	output := cmp.Run()
	log.Printf("Part 1: %s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output)), ","), "[]"))
}

type Computer struct {
	instructs []int
	instruct  int
	registers map[string]int
}

func NewComputer(instructions []int) Computer {
	return Computer{
		instructs: instructions,
		instruct:  0,
		registers: map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
		},
	}
}

func (c *Computer) Run() []int {
	output := make([]int, 0)
	for c.instruct < len(c.instructs) {
		op := c.instructs[c.instruct]
		opand := c.instructs[c.instruct+1]

		var progress bool
		var result []int

		switch op {
		case 0:
			progress, result = c.adv(opand)
		case 1:
			progress, result = c.bxl(opand)
		case 2:
			progress, result = c.bst(opand)
		case 3:
			progress, result = c.jnz(opand)
		case 4:
			progress, result = c.bxc(opand)
		case 5:
			progress, result = c.out(opand)
		case 6:
			progress, result = c.bdv(opand)
		case 7:
			progress, result = c.cdv(opand)
		}

		output = append(output, result...)
		if progress {
			c.instruct += 2
		}
	}
	return output
}

func (c *Computer) GetRegister(name string) int {
	return c.registers[name]
}

func (c *Computer) SetRegister(name string, val int) {
	c.registers[name] = val
}

func (c *Computer) combo(opand int) int {
	if opand < 4 {
		return opand
	}
	return c.GetRegister([]string{"a", "b", "c"}[opand-4])
}

func (c *Computer) adv(opand int) (bool, []int) {
	denominator := math.Pow(2, float64(c.combo(opand)))
	val := float64(c.GetRegister("a")) / denominator
	c.SetRegister("a", int(val))
	return true, []int{}
}

func (c *Computer) bxl(opand int) (bool, []int) {
	c.SetRegister("b", c.GetRegister("b")^opand)
	return true, []int{}
}

func (c *Computer) bst(opand int) (bool, []int) {
	c.SetRegister("b", c.combo(opand)%8)
	return true, []int{}
}

func (c *Computer) jnz(opand int) (bool, []int) {
	if c.GetRegister("a") != 0 {
		c.instruct = opand
		return false, []int{}
	}
	return true, []int{}
}

func (c *Computer) bxc(_ int) (bool, []int) {
	c.SetRegister("b", c.GetRegister("b")^c.GetRegister("c"))
	return true, []int{}
}

func (c *Computer) out(opand int) (bool, []int) {
	return true, []int{c.combo(opand) % 8}
}

func (c *Computer) bdv(opand int) (bool, []int) {
	val := c.GetRegister("a") / int(math.Pow(2, float64(c.combo(opand))))
	c.SetRegister("b", val)
	return true, []int{}
}

func (c *Computer) cdv(opand int) (bool, []int) {
	val := c.GetRegister("a") / int(math.Pow(2, float64(c.combo(opand))))
	c.SetRegister("c", val)
	return true, []int{}
}

func parse(input string) Computer {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`(\d+)`)
	a, _ := strconv.Atoi(re.FindAllString(lines[0], -1)[0])
	b, _ := strconv.Atoi(re.FindAllString(lines[1], -1)[0])
	c, _ := strconv.Atoi(re.FindAllString(lines[2], -1)[0])
	instructs := make([]int, 0)
	for _, num := range re.FindAllString(lines[4], -1) {
		n, _ := strconv.Atoi(num)
		instructs = append(instructs, n)
	}
	return Computer{
		instructs: instructs,
		instruct:  0,
		registers: map[string]int{
			"a": a,
			"b": b,
			"c": c,
		},
	}
}
