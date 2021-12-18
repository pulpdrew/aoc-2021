package day17

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	xStart := strings.Index(string(input), "x=") + 2
	xEnd := strings.Index(string(input), ", y=")

	xs := strings.Split(string(input[xStart:xEnd]), "..")
	ys := strings.Split(string(input[xEnd+4:]), "..")

	yMin, _ := strconv.Atoi(ys[0])
	yMax, _ := strconv.Atoi(ys[1])
	xMin, _ := strconv.Atoi(xs[0])
	xMax, _ := strconv.Atoi(xs[1])

	maxValidYVelocity := math.MinInt
	countOfVelocities := 0
	for yVel := 500; yVel >= -500; yVel-- {
		count := countValidVelocities(yVel, xMin, xMax, yMin, yMax)
		if count > 0 && yVel > maxValidYVelocity {
			maxValidYVelocity = yVel
		}
		countOfVelocities += count
	}

	maxYPosition := maxValidYVelocity * (maxValidYVelocity + 1) / 2
	fmt.Printf("Day 17 - Part 1: Max Y Position = %v\n", maxYPosition)
	fmt.Printf("Day 17 - Part 2: Number of Distinct Velocities = %v\n", countOfVelocities)

}

func countValidVelocities(yVel, xMin, xMax, yMin, yMax int) int {
	var a, b, cMin, cMax float64 = -1, float64(2*yVel + 1), float64(-2 * yMin), float64(-2 * yMax)

	validSet := make(map[int]bool)

	s1Min, s2Min := solveQuadratic(a, b, cMin)
	s1Max, s2Max := solveQuadratic(a, b, cMax)
	sMin, sMax := math.Max(s1Min, s2Min), math.Max(s1Max, s2Max)

	for step := int(math.Ceil(sMax)); step <= int(math.Floor(sMin)); step++ {
		for xVel := range validXVels(step, xMin, xMax) {
			validSet[xVel] = true
		}
	}

	return len(validSet)
}

// Returns the number of valid X velocities that result in an x position
// in the range [xMin, xMax] after `step` steps
func validXVels(step, xMin, xMax int) map[int]bool {

	validSet := make(map[int]bool)

	velMin := float64(step*step-step+2*xMin) / float64(2*step)
	velMax := float64(step*step-step+2*xMax) / float64(2*step)
	for vel := int(math.Ceil(velMin)); vel <= int(math.Floor(velMax)) && vel > step; vel++ {
		validSet[vel] = true
	}

	v1Min, v2Min := solveQuadratic(1, 1, float64(-2*xMin))
	v1Max, v2Max := solveQuadratic(1, 1, float64(-2*xMax))
	vMin, vMax := math.Max(v1Min, v2Min), math.Max(v1Max, v2Max)
	for vel := int(math.Ceil(vMin)); vel <= int(math.Floor(vMax)) && vel <= step; vel++ {
		validSet[vel] = true
	}

	return validSet
}

func solveQuadratic(a, b, c float64) (x1, x2 float64) {
	x1 = ((-1 * b) + math.Sqrt((b*b)-(4*a*c))) / (2 * a)
	x2 = ((-1 * b) - math.Sqrt((b*b)-(4*a*c))) / (2 * a)

	return x1, x2
}
