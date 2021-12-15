package day15

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Run() {

	part1()
	part2()

}

func part1() {

	puzzleInput, err := os.ReadFile("./day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nodes := [][]node{}
	for row, risks := range strings.Split(string(puzzleInput), "\n") {
		nodeRow := []node{}
		for col, risk := range risks {
			nodeRow = append(nodeRow, node{
				localRisk:           int(risk - '0'),
				leastCumulativeRisk: math.MaxInt,
				row:                 row,
				col:                 col,
			})
		}
		nodes = append(nodes, nodeRow)
	}

	n := leastRiskPath(nodes)
	fmt.Printf("Day 15 - Part 1: The least risky path has cumulative risk %v\n", n.leastCumulativeRisk)
}

func part2() {
	puzzleInput, err := os.ReadFile("./day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(puzzleInput), "\n")
	inputRowCount := len(lines)
	inputColCount := len(lines[0])

	nodes := [][]node{}
	for row := 0; row < 5*inputRowCount; row++ {
		nodes = append(nodes, []node{})
		for col := 0; col < 5*inputColCount; col++ {
			input := int(lines[row%inputRowCount][col%inputColCount] - '0')
			risk := (input+row/inputRowCount+col/inputColCount-1)%9 + 1
			nodes[row] = append(nodes[row], node{
				localRisk:           risk,
				leastCumulativeRisk: math.MaxInt,
				row:                 row,
				col:                 col,
			})
		}
	}

	n := leastRiskPath(nodes)
	fmt.Printf("\nDay 15 - Part 2: The least risky path has cumulative risk %v\n", n.leastCumulativeRisk)
}

type node struct {
	localRisk           int
	leastCumulativeRisk int
	row, col            int
	addedToFringe       bool
}

/////////////////// Priority Queue Implementation ////////////////////////

type NodeHeap []*node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return (h[i].leastCumulativeRisk) < (h[j].leastCumulativeRisk)
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/////////////////// Helper Functions ////////////////////////////

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

var offsets = [4][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func getNeighbors(n *node, nodes [][]node) []*node {
	neighbors := []*node{}
	rows, cols := len(nodes), len(nodes[0])

	for _, offset := range offsets {
		row, col := n.row+offset[0], n.col+offset[1]
		if row >= 0 && row < rows && col >= 0 && col < cols {
			neighbors = append(neighbors, &nodes[row][col])
		}
	}

	return neighbors
}

///////////////////// A* ///////////////////////////////

func leastRiskPath(nodes [][]node) *node {

	rows, cols := len(nodes), len(nodes[0])

	// Create a Priority Queue of "discovered but not explored nodes"
	fringe := NodeHeap{}
	heap.Init(&fringe)

	// Start with the starting node in the fringe
	nodes[0][0].leastCumulativeRisk = 0
	nodes[0][0].addedToFringe = true
	heap.Push(&fringe, &nodes[0][0])

	// Explore until the end is reached
	for len(fringe) > 0 {
		current := heap.Pop(&fringe).(*node)

		// Update each neighbor's cumulative risk + add to the fringe if necessary
		for _, neighbor := range getNeighbors(current, nodes) {
			neighbor.leastCumulativeRisk = min(neighbor.leastCumulativeRisk, neighbor.localRisk+current.leastCumulativeRisk)

			if !neighbor.addedToFringe {
				heap.Push(&fringe, neighbor)
				neighbor.addedToFringe = true
			}
		}

		heap.Init(&fringe)
	}

	return &nodes[rows-1][cols-1]
}
