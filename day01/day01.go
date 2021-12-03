package day01

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(puzzleInput), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	count := 0
	for i := 1; i < len(lines); i++ {
		prev, _ := strconv.Atoi(lines[i-1])
		cur, _ := strconv.Atoi(lines[i])
		if prev < cur {
			count++
		}
	}

	fmt.Printf("Day 1 - Part 1: %v\n", count)
}

func part2(lines []string) {
	first, _ := strconv.Atoi(lines[0])
	second, _ := strconv.Atoi(lines[1])
	third, _ := strconv.Atoi(lines[2])
	previousSum := first + second + third

	count := 0
	for i := 3; i < len(lines); i++ {
		previous, _ := strconv.Atoi(lines[i-3])
		next, _ := strconv.Atoi(lines[i])
		sum := previousSum + next - previous

		if sum > previousSum {
			count++
		}

		previousSum = sum
	}

	fmt.Printf("Day 1 - Part 2: %v\n", count)
}
