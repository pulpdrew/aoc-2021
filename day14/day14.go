package day14

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputSections := strings.Split(string(puzzleInput), "\n\n")

	template := inputSections[0]
	rules := make(map[string]rune)
	for _, line := range strings.Split(inputSections[1], "\n") {
		ruleParts := strings.Split(line, " -> ")
		rules[ruleParts[0]] = rune(ruleParts[1][0])
	}

	currentPairCounts := make(map[string]uint64)
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		currentPairCounts[pair]++
	}

	currentRuneCounts := make(map[rune]uint64)
	for _, r := range template {
		currentRuneCounts[r]++
	}

	completedSteps := 0
	for ; completedSteps < 10; completedSteps++ {
		currentPairCounts, currentRuneCounts = step(currentPairCounts, currentRuneCounts, rules)
	}

	min, max := minAndMax(currentRuneCounts)
	fmt.Printf("Day 14 - Part 1: After step %v, max (%v) - min (%v) = %v\n", completedSteps, max, min, max-min)

	for ; completedSteps < 40; completedSteps++ {
		currentPairCounts, currentRuneCounts = step(currentPairCounts, currentRuneCounts, rules)
	}

	min, max = minAndMax(currentRuneCounts)
	fmt.Printf("Day 14 - Part 2: After step %v, max (%v) - min (%v) = %v\n", completedSteps, max, min, max-min)
}

func minAndMax(counts map[rune]uint64) (min, max uint64) {
	min, max = math.MaxInt64, 0
	for _, count := range counts {
		if count < min {
			min = count
		} else if count > max {
			max = count
		}
	}
	return min, max
}

func step(currentPairsCounts map[string]uint64, currentRuneCount map[rune]uint64, rules map[string]rune) (map[string]uint64, map[rune]uint64) {

	// Copy the existing rune counts
	newRunesCounts := make(map[rune]uint64)
	for r, count := range currentRuneCount {
		newRunesCounts[r] += count
	}

	// Calculate the new pairs, tracking the rune count
	newPairsCounts := make(map[string]uint64)
	for pair, inserted := range rules {
		firstNewPair := pair[0:1] + string(inserted)
		newPairsCounts[firstNewPair] += currentPairsCounts[pair]

		secondNewPair := string(inserted) + pair[1:2]
		newPairsCounts[secondNewPair] += currentPairsCounts[pair]

		newRunesCounts[inserted] += currentPairsCounts[pair]
	}

	return newPairsCounts, newRunesCounts
}
