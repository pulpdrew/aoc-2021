package day03

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
	puzzleInput, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(puzzleInput), "\n")

	part1(lines)
	part2(lines)
}

func mostCommonBit(lines []string, bit int) int {
	count1 := 0
	for _, line := range lines {
		if line[bit] == '1' {
			count1++
		}
	}

	if count1 >= len(lines)-count1 {
		return 1
	} else {
		return 0
	}
}

func part1(lines []string) {
	gamma := 0
	epsilon := 0
	for bit := range lines[0] {
		digit := mostCommonBit(lines, bit)
		inverse := 1 - digit
		gamma = (gamma << 1) | digit
		epsilon = (epsilon << 1) | inverse
	}

	fmt.Printf("Day 3 - Part 1: (gamma: %v x epsilon: %v) = %v\n", gamma, epsilon, gamma*epsilon)
}

func part2(lines []string) {
	oxygenLines := lines
	for bit := 0; len(oxygenLines) > 1; bit++ {
		mostCommon := mostCommonBit(oxygenLines, bit)
		newOxygenLines := []string{}
		for _, line := range oxygenLines {
			if (line[bit] == '1') == (mostCommon == 1) {
				newOxygenLines = append(newOxygenLines, line)
			}
		}

		oxygenLines = newOxygenLines
	}

	co2Lines := lines
	for bit := 0; len(co2Lines) > 1; bit++ {
		leastCommon := 1 - mostCommonBit(co2Lines, bit)
		newCo2Lines := []string{}
		for _, line := range co2Lines {
			if (line[bit] == '1') == (leastCommon == 1) {
				newCo2Lines = append(newCo2Lines, line)
			}
		}

		co2Lines = newCo2Lines
	}

	oxygenInt, _ := strconv.ParseInt(oxygenLines[0], 2, 64)
	co2Int, _ := strconv.ParseInt(co2Lines[0], 2, 64)
	fmt.Printf("Day 3 - Part 2: (%v oxygen x %v c02) = %v\n", oxygenLines[0], co2Lines[0], oxygenInt*co2Int)
}
