package day25

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day25/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	floor := SeaFloor{
		cucumbers: make(map[Coords]SeaCucumber),
		yLen:      len(strings.Split(string(input), "\n")),
		xLen:      len(strings.Split(string(input), "\n")[0]),
	}

	for y, line := range strings.Split(string(input), "\n") {
		for x, r := range line {
			floor.cucumbers[coords(x, y)] = SeaCucumber{
				exists:       r != '.',
				isEastFacing: r == '>',
			}
		}
	}

	moved := true
	steps := 0
	for ; moved; steps++ {
		floor, moved = floor.step()
	}

	fmt.Printf("Day 25 - Part 1: The cucumbers stop moving after %v steps\n", steps)
}

type Coords [2]int

func coords(x, y int) Coords {
	return [2]int{x, y}
}

type SeaCucumber struct {
	exists       bool
	isEastFacing bool
}

type SeaFloor struct {
	cucumbers  map[Coords]SeaCucumber
	xLen, yLen int
}

func (floor SeaFloor) String() string {
	builder := strings.Builder{}

	for y := 0; y < floor.yLen; y++ {
		for x := 0; x < floor.xLen; x++ {
			cucumber := floor.get(coords(x, y))
			if cucumber.exists && cucumber.isEastFacing {
				builder.WriteRune('>')
			} else if cucumber.exists {
				builder.WriteRune('v')
			} else {
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}

func (floor SeaFloor) get(c Coords) SeaCucumber {
	wrapped := coords(c[0]%floor.xLen, c[1]%floor.yLen)
	return floor.cucumbers[wrapped]
}

func (floor *SeaFloor) set(c Coords, cucumber SeaCucumber) {
	wrapped := coords(c[0]%floor.xLen, c[1]%floor.yLen)
	floor.cucumbers[wrapped] = cucumber
}

func (floor SeaFloor) step() (next SeaFloor, moved bool) {
	temp := SeaFloor{
		cucumbers: make(map[Coords]SeaCucumber),
		xLen:      floor.xLen,
		yLen:      floor.yLen,
	}

	// Copy South Facing, Move East Facing if possible
	for x := 0; x < floor.xLen; x++ {
		for y := 0; y < floor.yLen; y++ {
			location := coords(x, y)
			cucumber := floor.get(location)

			if cucumber.exists && cucumber.isEastFacing && !floor.get(coords(x+1, y)).exists {
				temp.set(coords(x+1, y), cucumber) // Move East Facing Cucumbers if possible
				moved = true
			} else if cucumber.exists {
				temp.set(location, cucumber) // Copy South Facing and immovable Cucumbers
			}
		}
	}

	next = SeaFloor{
		cucumbers: make(map[Coords]SeaCucumber),
		xLen:      floor.xLen,
		yLen:      floor.yLen,
	}

	// Move South Facing
	for x := 0; x < temp.xLen; x++ {
		for y := 0; y < temp.yLen; y++ {
			location := coords(x, y)
			cucumber := temp.get(location)

			if cucumber.exists && !cucumber.isEastFacing && !temp.get(coords(x, y+1)).exists {
				next.set(coords(x, y+1), cucumber) // Move South Facing Cucumbers if possible
				moved = true
			} else if cucumber.exists {
				next.set(location, cucumber) // Copy East Facing and immovable cucumbers
			}
		}
	}

	return next, moved
}
