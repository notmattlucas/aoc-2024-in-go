package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("If register C contains 9, the program 2,6 would set register B to 1.", func(t *testing.T) {
		computer := NewComputer([]int{2, 6})
		computer.SetRegister("c", 9)
		computer.Run()
		assert.Equal(t, 1, computer.GetRegister("b"))
	})

	t.Run("If register A contains 10, the program 5,0,5,1,5,4 would output 0,1,2.", func(t *testing.T) {
		computer := NewComputer([]int{5, 0, 5, 1, 5, 4})
		computer.SetRegister("a", 10)
		output := computer.Run()
		assert.Equal(t, []int{0, 1, 2}, output)
	})

	t.Run("If register A contains 2024, the program 0,1,5,4,3,0 would output 4,2,5,6,7,7,7,7,3,1,0 and leave 0 in register A.", func(t *testing.T) {
		computer := NewComputer([]int{0, 1, 5, 4, 3, 0})
		computer.SetRegister("a", 2024)
		output := computer.Run()
		assert.Equal(t, 0, computer.GetRegister("a"))
		assert.Equal(t, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}, output)
	})

	t.Run("If register B contains 29, the program 1,7 would set register B to 26.", func(t *testing.T) {
		computer := NewComputer([]int{1, 7})
		computer.SetRegister("b", 29)
		computer.Run()
		assert.Equal(t, 26, computer.GetRegister("b"))
	})

	t.Run("If register B contains 2024 and register C contains 43690, the program 4,0 would set register B to 44354.", func(t *testing.T) {
		computer := NewComputer([]int{4, 0})
		computer.SetRegister("b", 2024)
		computer.SetRegister("c", 43690)
		computer.Run()
		assert.Equal(t, 44354, computer.GetRegister("b"))
	})

	t.Run("Parse", func(t *testing.T) {
		cmp := parse(`Register A: 729
					Register B: 0
					Register C: 0
					
					Program: 0,1,5,4,3,0`)
		assert.Equal(t, 729, cmp.GetRegister("a"))
		assert.Equal(t, 0, cmp.GetRegister("b"))
		assert.Equal(t, 0, cmp.GetRegister("c"))
		assert.Equal(t, []int{0, 1, 5, 4, 3, 0}, cmp.instructs)
	})

	t.Run("Part 1", func(t *testing.T) {
		cmp := parse(`Register A: 729
					Register B: 0
					Register C: 0
					
					Program: 0,1,5,4,3,0`)
		assert.Equal(t, []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}, cmp.Run())
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {

	})

}
