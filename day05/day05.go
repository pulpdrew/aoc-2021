package day05

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func (line Line) IsHorizontal() bool {
	return line.start.y == line.end.y
}

func (line Line) IsVertical() bool {
	return line.start.x == line.end.x
}

func (line *Line) CoveredPoints() []Point {
	points := []Point{}

	startX, endX := line.start.x, line.end.x
	startY, endY := line.start.y, line.end.y

	xInc, yInc := 0, 0

	if startX < endX {
		xInc = 1
	} else if startX > endX {
		xInc = -1
	}

	if startY < endY {
		yInc = 1
	} else if startY > endY {
		yInc = -1
	}

	for x, y := startX, startY; x != endX+xInc || y != endY+yInc; x, y = x+xInc, y+yInc {
		points = append(points, Point{x, y})
	}

	return points
}

func Run() {
	puzzleInput, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputLines := strings.Split(string(puzzleInput), "\n")

	lines := []Line{}
	for _, line := range inputLines {
		points := strings.Split(line, " -> ")

		startCoords := strings.Split(points[0], ",")
		startX, _ := strconv.Atoi(startCoords[0])
		startY, _ := strconv.Atoi(startCoords[1])
		start := Point{
			x: startX,
			y: startY,
		}

		endCoords := strings.Split(points[1], ",")
		endX, _ := strconv.Atoi(endCoords[0])
		endY, _ := strconv.Atoi(endCoords[1])
		end := Point{
			x: endX,
			y: endY,
		}

		lines = append(lines, Line{start: start, end: end})
	}

	part1(lines)
	part2(lines)
}

func part1(lines []Line) {
	var coveredPoints = make(map[Point]int)
	for _, line := range lines {
		if line.IsHorizontal() || line.IsVertical() {
			for _, point := range line.CoveredPoints() {
				coveredPoints[point]++
			}
		}
	}

	numberOfDangerousAreas := 0
	for _, count := range coveredPoints {
		if count >= 2 {
			numberOfDangerousAreas++
		}
	}

	fmt.Printf("Day 5 - Part 1: Number of dangerous areas = %v\n", numberOfDangerousAreas)
}

func part2(lines []Line) {
	var coveredPoints = make(map[Point]int)
	for _, line := range lines {
		for _, point := range line.CoveredPoints() {
			coveredPoints[point]++
		}
	}

	numberOfDangerousAreas := 0
	for _, count := range coveredPoints {
		if count >= 2 {
			numberOfDangerousAreas++
		}
	}

	fmt.Printf("Day 5 - Part 2: Number of dangerous areas = %v\n", numberOfDangerousAreas)
}
