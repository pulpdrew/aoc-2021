package day06

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(puzzleInput), ",")
	fish := []int{}
	for _, input := range inputs {
		n, _ := strconv.Atoi(input)
		fish = append(fish, n)
	}

	solve(fish)
}

func solve(fish []int) {

	fishCounts := make([]int, 7)
	for _, f := range fish {
		fishCounts[f]++
	}

	fishWithTimer7, fishWithTimer8 := 0, 0
	for currentDay := 0; currentDay < 256; currentDay++ {

		newCount := fishCounts[currentDay%7]
		fishCounts[currentDay%7] += fishWithTimer7
		fishWithTimer7 = fishWithTimer8
		fishWithTimer8 = newCount

		if currentDay+1 == 80 {
			total := fishWithTimer8 + fishWithTimer7
			for _, count := range fishCounts {
				total += count
			}
			fmt.Printf("Day 6 - Part 2: %v fish after 80 days\n", total)

		} else if currentDay+1 == 256 {
			total := fishWithTimer8 + fishWithTimer7
			for _, count := range fishCounts {
				total += count
			}
			fmt.Printf("Day 6 - Part 2: %v fish after 256 days\n", total)
		}
	}
}
