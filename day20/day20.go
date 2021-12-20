package day20

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	algoDescription := strings.Split(string(input), "\n")[0]
	algo := [512]bool{}
	for i, r := range algoDescription {
		algo[i] = r == '#'
	}

	imageData := make(map[[2]int]bool)
	for row, line := range strings.Split(string(input), "\n")[2:] {
		for col, r := range line {
			if r == '#' {
				imageData[[2]int{row, col}] = true
			}
		}
	}

	xMin, xMax, yMin, yMax := getBounds(imageData)
	image := image{
		data:             imageData,
		xMin:             xMin,
		xMax:             xMax,
		yMin:             yMin,
		yMax:             yMax,
		outOfBoundsValue: false,
	}

	// Part 1

	enhanced := image.enhance(algo)
	doubleEnhanced := enhanced.enhance(algo)
	fmt.Printf("Day 20 - Part 1: After enhancing twice, there are %v lit pixels\n", len(doubleEnhanced.data))

	// Part 2

	for i := 0; i < 50; i++ {
		image = image.enhance(algo)
	}

	fmt.Printf("Day 20 - Part 2: After enhancing 50 times, there are %v lit pixels\n", len(image.data))
}

type image struct {
	data                   map[[2]int]bool
	xMin, xMax, yMin, yMax int
	outOfBoundsValue       bool
}

func (i *image) get(x, y int) bool {
	if x > i.xMax || x < i.xMin || y > i.yMax || y < i.yMin {
		return i.outOfBoundsValue
	} else {
		return i.data[[2]int{x, y}]
	}
}

func (i *image) set(x, y int) {
	i.data[[2]int{x, y}] = true
	i.xMin = min(i.xMin, x)
	i.xMax = max(i.xMax, x)
	i.yMin = min(i.yMin, y)
	i.yMax = max(i.yMax, y)
}

func (i *image) getPixelSignature(r, c int) (signature int) {
	for x := r - 1; x <= r+1; x++ {
		for y := c - 1; y <= c+1; y++ {
			if i.get(x, y) {
				signature++
			}
			signature *= 2
		}
	}

	return signature / 2
}

func (i *image) print() {
	for x := i.xMin - 1; x <= i.xMax+1; x++ {
		for y := i.yMin - 1; y <= i.yMax+1; y++ {
			if i.get(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (i *image) enhance(algo [512]bool) image {
	enhancedData := make(map[[2]int]bool)

	for x := i.xMin - 1; x <= i.xMax+1; x++ {
		for y := i.yMin - 1; y <= i.yMax+1; y++ {
			sig := i.getPixelSignature(x, y)
			if algo[sig] {
				enhancedData[[2]int{x, y}] = true
			}
		}
	}

	enhanced := image{
		data: enhancedData,
		xMin: i.xMin - 1,
		xMax: i.xMax + 1,
		yMin: i.yMin - 1,
		yMax: i.yMax + 1,
	}

	if i.outOfBoundsValue {
		enhanced.outOfBoundsValue = algo[511]
	} else {
		enhanced.outOfBoundsValue = algo[0]
	}

	return enhanced
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func getBounds(image map[[2]int]bool) (xMin, xMax, yMin, yMax int) {
	xMax, yMax = math.MinInt, math.MinInt
	xMin, yMin = math.MaxInt, math.MaxInt

	for coords, lit := range image {
		if !lit {
			continue
		}

		xMin = min(xMin, coords[0])
		yMin = min(yMin, coords[1])
		xMax = max(xMax, coords[0])
		yMax = max(yMax, coords[1])
	}

	return xMin, xMax, yMin, yMax
}
