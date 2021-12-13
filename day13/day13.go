package day13

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputSections := strings.Split(string(puzzleInput), "\n\n")

	points := [][2]int{}
	for _, line := range strings.Split(inputSections[0], "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points = append(points, [2]int{x, y})
	}

	folds := []fold{}
	for _, line := range strings.Split(inputSections[1], "\n") {
		foldOverY := strings.ContainsRune(line, 'y')
		value, _ := strconv.Atoi(strings.Split(line, "=")[1])
		folds = append(folds, fold{foldOverY, value})
	}

	// Part 1
	foldedPoints, _, _ := plot(points, folds[0:1])
	fmt.Printf("Day 13 Part 1: %v plotted points\n", len(foldedPoints))

	// Part 2
	foldedPoints2, maxX2, maxY2 := plot(points, folds)
	fmt.Printf("Day 13 Part 2: The Code is\n")
	print(foldedPoints2, maxX2, maxY2)
}

type fold struct {
	foldOverY bool
	value     int
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func print(points map[[2]int]bool, maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if points[[2]int{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func plot(points [][2]int, folds []fold) (map[[2]int]bool, int, int) {

	foldedPoints := make(map[[2]int]bool)
	maxX, maxY := 0, 0
	for _, point := range points {
		for _, fold := range folds {
			point = foldPoint(point, fold)
		}
		foldedPoints[point] = true

		maxX = max(maxX, point[0])
		maxY = max(maxY, point[1])
	}

	return foldedPoints, maxX, maxY
}

func foldPoint(point [2]int, f fold) [2]int {
	if f.foldOverY && point[1] > f.value {
		return [2]int{point[0], 2*f.value - point[1]}
	} else if !f.foldOverY && point[0] > f.value {
		return [2]int{2*f.value - point[0], point[1]}
	} else {
		return point
	}
}
