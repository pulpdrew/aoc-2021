package day19

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day19/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Input Parsing
	scans := list.New()
	for _, scanInput := range strings.Split(string(input), "\n\n") {

		points := [][3]int{}
		for _, pointInput := range strings.Split(scanInput, "\n")[1:] {
			coordinateInput := strings.Split(pointInput, ",")
			x, _ := strconv.Atoi(coordinateInput[0])
			y, _ := strconv.Atoi(coordinateInput[1])
			z, _ := strconv.Atoi(coordinateInput[2])
			points = append(points, [3]int{x, y, z})
		}

		scans.PushBack(newScan(points))
	}

	// Keep track of the translations, for part 2
	translations := [][3]int{}

	// Construct the full scan
	base := scans.Remove(scans.Front()).(*scan)
	base.rotations()

	for scans.Len() > 0 {
		other := scans.Remove(scans.Front()).(*scan)

		combined := false
		exists, commonDistances := base.getCorrespondingPoints(other)
		if exists {
			for _, rotatedScan := range other.rotations() {
				basePoint := base.distances[commonDistances]
				otherPoint := rotatedScan.distances[commonDistances]
				translation := difference(otherPoint, basePoint)

				overlap := base.countOverlap(rotatedScan, translation)
				if overlap >= 12 {
					base.addAll(rotatedScan, translation)
					translations = append(translations, translation)
					combined = true
				}
			}
		}

		if !combined {
			scans.PushBack(other)
		}
	}

	fmt.Printf("Day 19 - Part 1: there are %v total beacons\n", len(base.points))

	// Part 2

	largestDistance := 0
	for _, first := range translations {
		for _, second := range translations {
			dist := manhattan(first, second)
			if dist > largestDistance {
				largestDistance = dist
			}
		}
	}

	fmt.Printf("Day 19 - Part 2: The largest distance between scanners is %v\n", largestDistance)
}

////////////// Scan Stuff ///////////////////

type scan struct {
	points    map[[3]int]bool
	distances map[[2]int][3]int
}

func newScan(points [][3]int) *scan {
	s := new(scan)
	s.points = make(map[[3]int]bool)
	s.distances = make(map[[2]int][3]int)

	for _, point := range points {
		s.points[point] = true
	}

	for _, point := range points {
		distances := []int{}
		for _, other := range points {
			distances = append(distances, distanceSquared(point, other))
		}
		sort.Ints(distances)
		s.distances[[2]int{distances[1], distances[2]}] = point
	}

	return s
}

func (s *scan) roll() *scan {
	new := new(scan)
	new.points = make(map[[3]int]bool)
	new.distances = make(map[[2]int][3]int)

	for point := range s.points {
		new.points[roll(point)] = true
	}

	for distances, point := range s.distances {
		new.distances[distances] = roll(point)
	}

	return new
}

func (s *scan) turn() *scan {
	new := new(scan)
	new.points = make(map[[3]int]bool)
	new.distances = make(map[[2]int][3]int)

	for point := range s.points {
		new.points[turn(point)] = true
	}

	for distances, point := range s.distances {
		new.distances[distances] = turn(point)
	}

	return new
}

func (s *scan) rotations() []*scan {
	scans := []*scan{}

	// https://stackoverflow.com/questions/16452383/how-to-get-all-24-rotations-of-a-3-dimensional-array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			s = s.roll()
			scans = append(scans, s)
			for k := 0; k < 3; k++ {
				s = s.turn()
				scans = append(scans, s)
			}
		}
		s = s.roll().turn().roll()
	}

	return scans
}

func (base *scan) countOverlap(other *scan, translation [3]int) (count int) {

	for p := range other.points {
		transformed := translate(p, translation)
		_, overlaps := base.points[transformed]
		if overlaps {
			count++
		}
	}

	return count
}

func (base *scan) addAll(source *scan, translation [3]int) {
	for p := range source.points {
		transformed := translate(p, translation)
		base.points[transformed] = true
	}

	base.distances = make(map[[2]int][3]int)
	for point := range base.points {
		distances := []int{}
		for other := range base.points {
			distances = append(distances, distanceSquared(point, other))
		}
		sort.Ints(distances)
		base.distances[[2]int{distances[1], distances[2]}] = point
	}
}

func (base *scan) getCorrespondingPoints(other *scan) (exists bool, differences [2]int) {

	for distances := range base.distances {
		_, exists = other.distances[distances]
		if exists {
			return true, distances
		}
	}

	return false, [2]int{}
}

////////////// Vector Stuff /////////////////

func distanceSquared(p1, p2 [3]int) int {
	sum := 0
	for c := range p1 {
		sum += (p2[c] - p1[c]) * (p2[c] - p1[c])
	}
	return sum
}

func manhattan(p1, p2 [3]int) int {
	sum := 0
	for c := range p1 {
		sum += abs(p2[c] - p1[c])
	}
	return sum
}

func turn(p [3]int) [3]int {
	return [3]int{p[0], p[2], -p[1]}
}

func roll(p [3]int) [3]int {
	return [3]int{-p[1], p[0], p[2]}
}

func difference(p1, p2 [3]int) [3]int {
	return [3]int{int(p2[0] - p1[0]), int(p2[1] - p1[1]), int(p2[2] - p1[2])}
}

func translate(p, trans [3]int) [3]int {
	return [3]int{p[0] + trans[0], p[1] + trans[1], p[2] + trans[2]}
}

//////////// Helper Functions //////////////

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
