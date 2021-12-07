package day07

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(puzzleInput), ",")
	positions := []int{}
	for _, input := range inputs {
		n, _ := strconv.Atoi(input)
		positions = append(positions, n)
	}

	sort.Ints(positions)
	part1(positions)
	part2(positions)
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func part1(positions []int) {
	median := positions[len(positions)/2]
	fuel := 0
	for _, position := range positions {
		fuel += abs(position - median)
	}
	fmt.Printf("Day 7 - Part 1: position %v results in a minimal fuel consumption of %v\n", median, fuel)
}

func p2Cost(distance int) int {
	cost := 0
	for i := 1; i <= distance; i++ {
		cost += i
	}
	return cost
}

func p2TotalCost(alignment int, positions []int) int {
	fuel := 0
	for _, position := range positions {
		fuel += p2Cost(abs(position - alignment))
	}
	return fuel
}

func average(positions []int) int {
	sum := 0
	for _, position := range positions {
		sum += position
	}
	return int(math.Round(float64(sum) / float64(len(positions))))
}

func part2(positions []int) {

	alignment := average(positions)
	bestCost := p2TotalCost(alignment, positions)
	for {
		costBelow := p2TotalCost(alignment-1, positions)
		if costBelow < bestCost {
			bestCost = costBelow
			alignment--
			continue
		}

		costAbove := p2TotalCost(alignment+1, positions)
		if costAbove < bestCost {
			bestCost = costAbove
			alignment++
			continue
		}

		break
	}

	fmt.Printf("Day 7 - Part 2: position %v results in a minimal fuel consumption of %v\n", alignment, bestCost)
}
