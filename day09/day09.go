package day09

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	heights := [][]int{}
	for _, line := range strings.Split(string(puzzleInput), "\n") {
		row := []int{}
		for _, height := range line {
			row = append(row, int(height-'0'))
		}
		heights = append(heights, row)
	}

	part1(heights)
	part2(heights)
}

func getAdjacent(heights [][]int, row int, col int) []int {
	adjacents := []int{}

	if row < len(heights)-1 {
		adjacents = append(adjacents, heights[row+1][col])
	}

	if row > 0 {
		adjacents = append(adjacents, heights[row-1][col])
	}

	if col < len(heights[row])-1 {
		adjacents = append(adjacents, heights[row][col+1])
	}

	if col > 0 {
		adjacents = append(adjacents, heights[row][col-1])
	}

	return adjacents
}

func flood(heights [][]int, row int, col int) int {
	floodCount := 0

	if heights[row][col] != -1 {
		heights[row][col] = -1
		floodCount++
	}

	if row < len(heights)-1 && heights[row+1][col] != 9 && heights[row+1][col] != -1 {
		heights[row+1][col] = -1
		floodCount += 1 + flood(heights, row+1, col)
	}

	if row > 0 && heights[row-1][col] != 9 && heights[row-1][col] != -1 {
		heights[row-1][col] = -1
		floodCount += 1 + flood(heights, row-1, col)
	}

	if col < len(heights[row])-1 && heights[row][col+1] != 9 && heights[row][col+1] != -1 {
		heights[row][col+1] = -1
		floodCount += 1 + flood(heights, row, col+1)
	}

	if col > 0 && heights[row][col-1] != 9 && heights[row][col-1] != -1 {
		heights[row][col-1] = -1
		floodCount += 1 + flood(heights, row, col-1)
	}

	return floodCount
}

func part1(heights [][]int) {

	sum := 0
	for row := range heights {
		for col := range heights[row] {
			adjacents := getAdjacent(heights, row, col)
			isLowPoint := true
			for _, height := range adjacents {
				if heights[row][col] >= height {
					isLowPoint = false
				}
			}

			if isLowPoint {
				sum += 1 + heights[row][col]
			}
		}
	}

	fmt.Printf("Day 9 - Part 1: Sum of the risk levels: %v\n", sum)
}

func part2(heights [][]int) {

	basinSizes := []int{}
	for row := range heights {
		for col := range heights[row] {
			adjacents := getAdjacent(heights, row, col)
			isLowPoint := true
			for _, height := range adjacents {
				if heights[row][col] >= height {
					isLowPoint = false
				}
			}

			if isLowPoint {
				basinSizes = append(basinSizes, flood(heights, row, col))
			}
		}
	}

	sort.Ints(basinSizes)
	product := 1
	for _, size := range basinSizes[len(basinSizes)-3:] {
		product *= size
	}

	fmt.Printf("Day 9 - Part 2: Basin Sizes: %v. Product = %v\n", basinSizes[len(basinSizes)-3:], product)
}
