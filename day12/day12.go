package day12

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputLines := strings.Split(string(puzzleInput), "\n")

	caves := make(map[string][]string)
	for _, line := range inputLines {
		first := strings.Split(line, "-")[0]
		second := strings.Split(line, "-")[1]

		caves[first] = append(caves[first], second)
		caves[second] = append(caves[second], first)
	}

	part1 := countPaths("start", caves, make(map[string]int), canVisitPart1)
	fmt.Printf("Day 12 - Part 1: %v paths\n", part1)
	part2 := countPaths("start", caves, make(map[string]int), canVisitPart2)
	fmt.Printf("Day 12 - Part 2: %v paths\n", part2)
}

func copy(original *map[string]int) map[string]int {
	copy := make(map[string]int)
	for k, v := range *original {
		copy[k] = v
	}
	return copy
}

func canVisitPart1(cave string, visitCounts map[string]int) bool {
	return visitCounts[cave] == 0 || strings.ToUpper(cave) == cave
}

func isBigCave(cave string) bool {
	return strings.ToUpper(cave) == cave
}

func canVisitPart2(cave string, visitCounts map[string]int) bool {
	if visitCounts[cave] == 0 || isBigCave(cave) {
		return true
	} else if (cave == "start" || cave == "end") && visitCounts[cave] == 1 {
		return false
	} else {
		for name, visits := range visitCounts {
			if !isBigCave(name) && visits > 1 {
				return false // If any small cave has been visited more than once
			}
		}
		return true
	}
}

func countPaths(start string, caves map[string][]string, visited map[string]int, canVisit func(string, map[string]int) bool) int {

	if start == "end" {
		return 1
	}

	newVisited := copy(&visited)
	newVisited[start]++

	paths := 0
	for _, next := range caves[start] {
		if canVisit(next, newVisited) {
			paths += countPaths(next, caves, newVisited, canVisit)
		}
	}

	return paths
}
