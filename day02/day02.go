package day02

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	magnitude int
}

func Run() {
	puzzleInput, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(puzzleInput), "\n")

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		magnitude, _ := strconv.Atoi(splitLine[1])
		instructions[i] = Instruction{direction: splitLine[0], magnitude: magnitude}
	}

	part1(instructions)
	part2(instructions)
}

func part1(instructions []Instruction) {

	depth := 0
	position := 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case "forward":
			position += instruction.magnitude
		case "up":
			depth -= instruction.magnitude
		case "down":
			depth += instruction.magnitude
		}
	}

	fmt.Printf("Day 2 - Part 1:\n (%v depth x %v position) = %v\n", depth, position, depth*position)
}

func part2(instructions []Instruction) {

	aim := 0
	depth := 0
	position := 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case "forward":
			position += instruction.magnitude
			depth += aim * instruction.magnitude
		case "up":
			aim -= instruction.magnitude
		case "down":
			aim += instruction.magnitude
		}
	}

	fmt.Printf("Day 2 - Part 2:\n (%v depth x %v position) = %v\n", depth, position, depth*position)
}
