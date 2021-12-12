package day11

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var offsets = [8][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	{-1, -1}, {1, 1}, {1, -1}, {-1, 1},
}

type octo struct {
	value      int
	hasFlashed bool
}

type cavern struct {
	octos [10][10]octo
}

func (c cavern) step() cavern {

	// Reset hasFlashed and increment value for each octo
	for row := range c.octos {
		for col := range c.octos[row] {
			octo := &c.octos[row][col]

			octo.hasFlashed = false
			octo.value++
		}
	}

	// Iterate until there are no more changes to be made
	hasChanged := true
	for hasChanged {
		hasChanged = false

		for row := range c.octos {
			for col := range c.octos[row] {
				octo := &c.octos[row][col]

				// Check each octopus that hasn't yet flashed
				if !octo.hasFlashed && octo.value > 9 {

					// Increment all of the adjacent octos if this octo flashes
					for _, offset := range offsets {
						adjRow, adjCol := row+offset[0], col+offset[1]
						if adjRow >= 0 && adjRow < len(c.octos) && adjCol >= 0 && adjCol < len(c.octos[row]) {
							c.octos[adjRow][adjCol].value++
						}
					}
					octo.hasFlashed = true
					hasChanged = true
				}

			}
		}

	}

	// Set all the octos who flashed to 0
	for row := range c.octos {
		for col := range c.octos[row] {
			octo := &c.octos[row][col]
			if octo.hasFlashed {
				octo.value = 0
			}
		}
	}

	return c
}

func (c *cavern) countFlashes() int {
	count := 0
	for row := range c.octos {
		for col := range c.octos[row] {
			octo := &c.octos[row][col]
			if octo.hasFlashed {
				count++
			}
		}
	}
	return count
}

func (c *cavern) print() {
	for row := range c.octos {
		for col := range c.octos[row] {
			octo := &c.octos[row][col]
			fmt.Print(octo.value)
		}
		fmt.Println()
	}
}

func Run() {
	puzzleInput, err := os.ReadFile("./day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputLines := strings.Split(string(puzzleInput), "\n")

	cavern := cavern{}
	for row, line := range inputLines {
		for col, value := range line {
			cavern.octos[row][col] = octo{value: int(value - '0'), hasFlashed: false}
		}
	}

	part1(cavern)
	part2(cavern)
}

func part1(c cavern) {
	flashes := 0
	for step := 0; step < 100; step++ {
		c = c.step()
		flashes += c.countFlashes()
	}
	fmt.Printf("Day 11 - Part 1: After step 100, there have been a total of %v flashes\n", flashes)
}

func part2(c cavern) {
	steps := 0
	for c.countFlashes() != 100 {
		c = c.step()
		steps++
	}
	fmt.Printf("Day 11 - Part 2: All octos flashed after step %v\n", steps)
}
