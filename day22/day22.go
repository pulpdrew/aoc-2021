package day22

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reactor := Octree{}

	for _, line := range strings.Split(string(input), "\n") {
		xStart, xEnd := strings.Index(line, "x=")+2, strings.Index(line, ",y=")
		yStart, yEnd := strings.Index(line, "y=")+2, strings.Index(line, ",z=")
		zStart, zEnd := strings.Index(line, "z=")+2, len(line)

		x1, _ := strconv.Atoi(strings.Split(line[xStart:xEnd], "..")[0])
		x2, _ := strconv.Atoi(strings.Split(line[xStart:xEnd], "..")[1])
		y1, _ := strconv.Atoi(strings.Split(line[yStart:yEnd], "..")[0])
		y2, _ := strconv.Atoi(strings.Split(line[yStart:yEnd], "..")[1])
		z1, _ := strconv.Atoi(strings.Split(line[zStart:zEnd], "..")[0])
		z2, _ := strconv.Atoi(strings.Split(line[zStart:zEnd], "..")[1])

		b := bounds{x1, x2 + 1, y1, y2 + 1, z1, z2 + 1}
		on := strings.Contains(line, "on")
		reactor.set(b, on)
	}

	fmt.Printf("Day 22 - Part 1: %v initialization cubes are on\n", reactor.count(bounds{-50, 51, -50, 51, -50, 51}))
	fmt.Printf("Day 22 - Part 2: %v total cubes are on\n", reactor.count(bounds{math.MinInt, math.MaxInt, math.MinInt, math.MaxInt, math.MinInt, math.MaxInt}))

}

type bounds struct {
	x1, x2, y1, y2, z1, z2 int
}

func (b bounds) partition(x, y, z int) (partitions []bounds) {
	partitions = append(partitions, b)

	temp := []bounds{}
	if b.x1 < x && b.x2 > x {
		for _, inst := range partitions {
			lower := inst
			lower.x2 = x
			upper := inst
			upper.x1 = x
			temp = append(temp, lower)
			temp = append(temp, upper)
		}
		partitions = temp
	}

	temp = []bounds{}
	if b.y1 < y && b.y2 > y {
		for _, inst := range partitions {
			lower := inst
			lower.y2 = y
			upper := inst
			upper.y1 = y
			temp = append(temp, lower)
			temp = append(temp, upper)
		}
		partitions = temp
	}

	temp = []bounds{}
	if b.z1 < z && b.z2 > z {
		for _, inst := range partitions {
			lower := inst
			lower.z2 = z
			upper := inst
			upper.z1 = z
			temp = append(temp, lower)
			temp = append(temp, upper)
		}
		partitions = temp
	}

	return partitions
}

type Octree struct {

	// true iff this is an interior node (eg. has children)
	interior bool

	// if this is a leaf node and on is true, then the region is entirely ON
	on bool

	// the axes on which to partition
	x, y, z int

	// the partitions node representing each of the partitions
	partitions [8]*Octree
}

func (n *Octree) partitionContaining(b bounds) *Octree {
	x, y, z := 0, 0, 0

	if b.x1 >= n.x {
		x = 1
	}

	if b.y1 >= n.y {
		y = 1
	}

	if b.z1 >= n.z {
		z = 1
	}

	return n.partitions[(x<<2)+(y<<1)+z]
}

func (n *Octree) set(b bounds, on bool) {

	if !n.interior && (n.on == on) {
		return
	}

	if !n.interior {
		// Partition this node based on the min bounds
		n.interior = true
		n.x = b.x1
		n.y = b.y1
		n.z = b.z1
		for i := 0; i < len(n.partitions); i++ {
			n.partitions[i] = new(Octree)
			n.partitions[i].on = n.on
		}

		// Partition the child based on the max bounds
		child := n.partitionContaining(b)
		child.interior = true
		child.x = b.x2
		child.y = b.y2
		child.z = b.z2
		grandchildren := &child.partitions
		for i := 0; i < len(grandchildren); i++ {
			grandchildren[i] = new(Octree)
			grandchildren[i].on = n.on
		}

		child.partitionContaining(b).on = on
		return
	}

	for _, split := range b.partition(n.x, n.y, n.z) {
		n.partitionContaining(split).set(split, on)
	}
}

func (n *Octree) count(bounds bounds) (count int) {
	if !n.interior && n.on {
		count = (bounds.x2 - bounds.x1) * (bounds.y2 - bounds.y1) * (bounds.z2 - bounds.z1)
	} else if n.interior {
		for _, split := range bounds.partition(n.x, n.y, n.z) {
			count += n.partitionContaining(split).count(split)
		}
	}
	return count
}
