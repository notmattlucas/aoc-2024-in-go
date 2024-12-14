package main

import (
	collections "github.com/notmattlucas/aoc-2024-in-go"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	robots, room := parse(string(b), 101, 103)
	robots = move(robots, room, 100)
	log.Printf("Part 1: %d", quadCount(robots, room))

	robots, room = parse(string(b), 101, 103)
	log.Printf("Part 2: %d", findTree(robots, room, 10000))

}

type Velocity struct{ x, y int }

type Position struct{ x, y int }

type Room struct{ width, height int }

type Positions = [][]bool

func (room Room) luq(pos Position) bool {
	return pos.x < room.width/2 && pos.y < room.height/2
}

func (room Room) ruq(pos Position) bool {
	return pos.x >= (room.width-room.width/2) && pos.y < room.height/2
}

func (room Room) llq(pos Position) bool {
	return pos.x < room.width/2 && pos.y >= (room.height-room.height/2)
}

func (room Room) rlq(pos Position) bool {
	return pos.x >= (room.width-room.width/2) && pos.y >= (room.height-room.height/2)
}

func (room Room) positions(robots []Robot) Positions {
	positions := make(map[Position]bool)
	for _, robot := range robots {
		positions[robot.position] = true
	}

	grid := make([][]bool, room.height)
	for y := 0; y < room.height; y++ {
		grid[y] = make([]bool, room.width)
		for x := 0; x < room.width; x++ {
			if positions[Position{x, y}] {
				grid[y][x] = true
			}
		}
	}

	return grid
}

type Robot struct {
	position Position
	velocity Velocity
}

func (robot Robot) move(room Room, times int) Robot {
	px := (((robot.position.x + (robot.velocity.x * times)) % room.width) + room.width) % room.width
	py := (((robot.position.y + (robot.velocity.y * times)) % room.height) + room.height) % room.height
	return Robot{
		Position{px, py},
		robot.velocity,
	}
}

func get(positions Positions, x, y int) bool {
	if x < 0 || y < 0 || y >= len(positions) || x >= len(positions[y]) {
		return false
	}
	return positions[y][x]
}

func quadCount(robots []Robot, room Room) int {
	luq, ruq, llq, rlq := 0, 0, 0, 0
	for _, robot := range robots {
		pos := robot.position
		if room.luq(pos) {
			luq++
		} else if room.ruq(pos) {
			ruq++
		} else if room.llq(pos) {
			llq++
		} else if room.rlq(pos) {
			rlq++
		}
	}
	return luq * ruq * llq * rlq
}

func move(robots []Robot, room Room, times int) []Robot {
	moved := make([]Robot, 0)
	for _, robot := range robots {
		moved = append(moved, robot.move(room, times))
	}
	return moved
}

func findTree(robots []Robot, room Room, times int) int {
	for i := 1; i < times; i++ {
		moved := move(robots, room, i)
		if hasTreeTop(moved, room) {
			return i
		}
	}
	return -1
}

func hasTreeTop(robots []Robot, room Room) bool {
	positions := room.positions(robots)
	for _, robot := range robots {
		pos := robot.position
		shift := []Velocity{{0, 0},
			{-1, 1}, {0, 1}, {1, 1},
			{-2, 2}, {-1, 2}, {0, 2}, {1, 2}, {2, 2}}
		found, _ := collections.All(shift, func(v Velocity) bool {
			return get(positions, pos.x+v.x, pos.y+v.y)
		})
		if found {
			return true
		}
	}
	return false
}

func parse(input string, xmax int, ymax int) ([]Robot, Room) {
	re := regexp.MustCompile(`(-*\d+)`)
	robots := make([]Robot, 0)
	for _, row := range strings.Split(input, "\n") {
		parts := re.FindAllString(row, 4)
		px, _ := strconv.Atoi(parts[0])
		py, _ := strconv.Atoi(parts[1])
		vx, _ := strconv.Atoi(parts[2])
		vy, _ := strconv.Atoi(parts[3])
		robots = append(robots, Robot{
			Position{px, py},
			Velocity{vx, vy},
		})
	}
	return robots, Room{xmax, ymax}
}
