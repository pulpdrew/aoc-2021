package day23

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	// Part 1
	contents := [19]amphipod{}
	for i := 0; i <= 10; i++ {
		contents[i] = amphipod{exists: lines[1][i+1] != '.', _type: int(lines[1][i+1]) - int('A')}
	}
	contents[11] = amphipod{exists: lines[2][3] != '.', _type: int(lines[2][3]) - int('A')}
	contents[12] = amphipod{exists: lines[2][5] != '.', _type: int(lines[2][5]) - int('A')}
	contents[13] = amphipod{exists: lines[2][7] != '.', _type: int(lines[2][7]) - int('A')}
	contents[14] = amphipod{exists: lines[2][9] != '.', _type: int(lines[2][9]) - int('A')}
	contents[15] = amphipod{exists: lines[3][3] != '.', _type: int(lines[3][3]) - int('A')}
	contents[16] = amphipod{exists: lines[3][5] != '.', _type: int(lines[3][5]) - int('A')}
	contents[17] = amphipod{exists: lines[3][7] != '.', _type: int(lines[3][7]) - int('A')}
	contents[18] = amphipod{exists: lines[3][9] != '.', _type: int(lines[3][9]) - int('A')}

	s := P1State{
		contents: contents,
	}
	DfsP1(s, []P1State{s})
	fmt.Printf("Day 23 - Part 1: minimum energy = %v\n", minSolveCost)

	// Part 2
	contents2 := [27]amphipod{}
	for i := 0; i <= 10; i++ {
		contents2[i] = amphipod{exists: lines[1][i+1] != '.', _type: int(lines[1][i+1]) - int('A')}
	}
	contents2[11] = amphipod{exists: lines[2][3] != '.', _type: int(lines[2][3]) - int('A')}
	contents2[12] = amphipod{exists: lines[2][5] != '.', _type: int(lines[2][5]) - int('A')}
	contents2[13] = amphipod{exists: lines[2][7] != '.', _type: int(lines[2][7]) - int('A')}
	contents2[14] = amphipod{exists: lines[2][9] != '.', _type: int(lines[2][9]) - int('A')}
	contents2[23] = amphipod{exists: lines[3][3] != '.', _type: int(lines[3][3]) - int('A')}
	contents2[24] = amphipod{exists: lines[3][5] != '.', _type: int(lines[3][5]) - int('A')}
	contents2[25] = amphipod{exists: lines[3][7] != '.', _type: int(lines[3][7]) - int('A')}
	contents2[26] = amphipod{exists: lines[3][9] != '.', _type: int(lines[3][9]) - int('A')}

	contents2[15] = amphipod{exists: true, _type: 3}
	contents2[16] = amphipod{exists: true, _type: 2}
	contents2[17] = amphipod{exists: true, _type: 1}
	contents2[18] = amphipod{exists: true, _type: 0}
	contents2[19] = amphipod{exists: true, _type: 3}
	contents2[20] = amphipod{exists: true, _type: 1}
	contents2[21] = amphipod{exists: true, _type: 0}
	contents2[22] = amphipod{exists: true, _type: 2}

	s2 := P2State{
		contents: contents2,
	}

	minSolveCost = math.MaxInt
	DfsP2(s2)
	fmt.Printf("Day 23 - Part 2: minimum energy = %v\n", minSolveCost)
}

var minSolveCost = math.MaxInt

type amphipod struct {
	exists bool
	_type  int
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

func isRoom(i int) bool {
	return i > 10
}

func isHallway(i int) bool {
	return i >= 0 && i <= 10
}

func isJustOutsideRoom(i int) bool {
	return i == 2 || i == 4 || i == 6 || i == 8
}

func roomNumber(i int) int {
	return (i - 11) % 4
}

func roomDepth(i int) int {
	return (i-11)/4 + 1
}

func cost(a amphipod) int {
	switch a._type {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	}
	panic("invalid amphipod passed to cost()")
}

func getNearestHallwaySpace(i int) int {
	return roomNumber(i)*2 + 2
}

func isReachable(s P1State, source, dest int) (reachable bool, moves int) {
	start, stop := source, dest

	if s.isOccupied(dest) {
		return false, 0
	}

	// Swap to eliminate a case
	if isHallway(start) && isRoom(stop) {
		start, stop = stop, start
	}

	if isRoom(start) && isHallway(stop) {
		hallwayNearestToStart := getNearestHallwaySpace(start)

		// Check horizontal movement between hallwayNearestToStart and stop
		for cur := min(hallwayNearestToStart, stop); cur <= max(hallwayNearestToStart, stop); cur++ {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
		}

		// Check vertical movement between start and hallwayNearestToStart
		cur := start
		for isRoom(cur) {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
			cur -= 4
		}

		return true, roomDepth(start) + max(hallwayNearestToStart, stop) - min(hallwayNearestToStart, stop)
	}

	if isRoom(start) && isRoom(stop) {
		hallwayNearestToStart := getNearestHallwaySpace(start)
		hallwayNearestToStop := getNearestHallwaySpace(stop)

		// Check Horizontal movement
		for i := min(hallwayNearestToStart, hallwayNearestToStop); i <= max(hallwayNearestToStart, hallwayNearestToStop); i++ {
			if s.isOccupied(i) {
				return false, 0
			}
		}

		// Check vertical movement between start and hallwayNearestToStart
		cur := start
		for isRoom(cur) {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
			cur -= 4
		}

		// Check vertical movement between stop and hallwayNearestToStop
		cur = stop
		for isRoom(cur) {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
			cur -= 4
		}

		return true, roomDepth(start) + roomDepth(stop) + max(hallwayNearestToStart, hallwayNearestToStop) - min(hallwayNearestToStart, hallwayNearestToStop)
	}

	return false, 0
}

func getDestinationRoom(a amphipod, depth int) int {
	return 11 + 4*(depth-1) + a._type
}

///////////////// Day 2 Specific //////////////////////

var P2Costs = make(map[[27]amphipod]int)

func DfsP2(current P2State) {
	leastCost, hasExplored := P2Costs[current.contents]
	if !hasExplored || current.cost < leastCost {
		P2Costs[current.contents] = current.cost

		if current.isSolved() && current.cost < minSolveCost {
			minSolveCost = current.cost
		}

		if !current.isSolved() {
			for _, n := range current.next() {
				DfsP2(n)
			}
		}
	}
}

type P2State struct {
	contents [27]amphipod
	cost     int
}

func (s P2State) isSolved() bool {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			index := 4*row + 11 + col
			occupant := s.contents[index]
			if !occupant.exists || occupant._type != col {
				return false
			}
		}
	}
	return true
}

func (s P2State) isOccupied(i int) bool {
	return s.contents[i].exists
}

func (s P2State) next() (nexts []P2State) {

	for source, occupant := range s.contents {
		if occupant.exists {

			// Amphipod in the hallway must go to its destination
			if isHallway(source) && !isDestinationOccupiedByOtherTypeP2(s, occupant) {
				for d := 4; d > 0; d-- {
					reachable, moves := isReachableP2(s, source, getDestinationRoom(occupant, d))
					if reachable {
						nexts = append(nexts, P2State{
							contents: moveP2(s.contents, source, getDestinationRoom(occupant, d)),
							cost:     s.cost + cost(occupant)*moves,
						})
						break
					}
				}
			}

			// Amphipod in a room can go to hallway or destination
			if isRoom(source) && !isInBestSpaceP2(s, occupant, source) {
				for hallwaySpace := 0; hallwaySpace <= 10; hallwaySpace++ {
					reachable, moves := isReachableP2(s, source, hallwaySpace)
					if reachable && !isJustOutsideRoom(hallwaySpace) {
						nexts = append(nexts, P2State{
							contents: moveP2(s.contents, source, hallwaySpace),
							cost:     s.cost + cost(occupant)*moves,
						})
					}
				}

				if !isDestinationOccupiedByOtherTypeP2(s, occupant) {
					for d := 4; d > 0; d-- {
						reachable, moves := isReachableP2(s, source, getDestinationRoom(occupant, d))
						if reachable {
							nexts = append(nexts, P2State{
								contents: moveP2(s.contents, source, getDestinationRoom(occupant, d)),
								cost:     s.cost + cost(occupant)*moves,
							})
							break
						}
					}
				}
			}
		}
	}

	return nexts
}

func isReachableP2(s P2State, source, dest int) (reachable bool, moves int) {
	start, stop := source, dest

	if s.isOccupied(dest) {
		return false, 0
	}

	// Swap to eliminate a case
	if isHallway(start) && isRoom(stop) {
		start, stop = stop, start
	}

	if isRoom(start) && isHallway(stop) {
		hallwayNearestToStart := getNearestHallwaySpace(start)

		// Check horizontal movement between hallwayNearestToStart and stop
		for cur := min(hallwayNearestToStart, stop); cur <= max(hallwayNearestToStart, stop); cur++ {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
		}

		// Check vertical movement between start and hallwayNearestToStart
		cur := start
		for isRoom(cur) {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
			cur -= 4
		}

		return true, roomDepth(start) + max(hallwayNearestToStart, stop) - min(hallwayNearestToStart, stop)
	}

	if isRoom(start) && isRoom(stop) {
		hallwayNearestToStart := getNearestHallwaySpace(start)
		hallwayNearestToStop := getNearestHallwaySpace(stop)

		// Check Horizontal movement
		for i := min(hallwayNearestToStart, hallwayNearestToStop); i <= max(hallwayNearestToStart, hallwayNearestToStop); i++ {
			if s.isOccupied(i) {
				return false, 0
			}
		}

		// Check vertical movement between start and hallwayNearestToStart
		cur := start
		for isRoom(cur) {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
			cur -= 4
		}

		// Check vertical movement between stop and hallwayNearestToStop
		cur = stop
		for isRoom(cur) {
			if s.isOccupied(cur) && cur != source {
				return false, 0
			}
			cur -= 4
		}

		return true, roomDepth(start) + roomDepth(stop) + max(hallwayNearestToStart, hallwayNearestToStop) - min(hallwayNearestToStart, hallwayNearestToStop)
	}

	return false, 0
}

func isInBestSpaceP2(s P2State, a amphipod, space int) bool {
	n, d := roomNumber(space), roomDepth(space)

	// Am I in the wrong room? Or not in a room?
	if !isRoom(space) || n != a._type {
		return false
	}

	// Are there empty spaces or incorrect amphopods trapped below?
	for below := d + 1; below <= 4; below++ {
		occupant := s.contents[getDestinationRoom(a, below)]
		if !occupant.exists || occupant._type != n {
			return false
		}
	}

	return true
}

func isDestinationOccupiedByOtherTypeP2(s P2State, a amphipod) bool {
	for d := 1; d <= 4; d++ {
		occupant := s.contents[getDestinationRoom(a, d)]
		if occupant.exists && occupant._type != a._type {
			return true
		}
	}
	return false
}

func moveP2(contents [27]amphipod, src, dest int) [27]amphipod {
	copy := contents
	copy[dest] = contents[src]
	copy[src] = amphipod{}
	return copy
}

///////////////// Day 1 Specific //////////////////////

var P1Costs = make(map[[19]amphipod]int)

func DfsP1(current P1State, path []P1State) {
	leastCost, hasExplored := P1Costs[current.contents]
	if !hasExplored || current.cost < leastCost {
		P1Costs[current.contents] = current.cost

		if current.isSolved() && current.cost < minSolveCost {
			minSolveCost = current.cost
		}

		if !current.isSolved() {
			for _, n := range current.next() {
				DfsP1(n, append(path, n))
			}
		}
	}
}

type P1State struct {
	contents [19]amphipod
	cost     int
}

func (s P1State) isSolved() bool {
	for row := 0; row < 2; row++ {
		for col := 0; col < 4; col++ {
			index := 4*row + 11 + col
			occupant := s.contents[index]
			if !occupant.exists || occupant._type != col {
				return false
			}
		}
	}
	return true
}

func (s P1State) isOccupied(i int) bool {
	return s.contents[i].exists
}

func isInBestSpace(s P1State, a amphipod, space int) bool {
	isInDestinationTop := space == getDestinationRoom(a, 1)
	isInDestinationBottom := space == getDestinationRoom(a, 2)
	return isInDestinationBottom || (isInDestinationTop && !isDestinationOccupiedByOtherType(s, a))
}

func isDestinationOccupiedByOtherType(s P1State, a amphipod) bool {
	top := s.contents[getDestinationRoom(a, 1)]
	bottom := s.contents[getDestinationRoom(a, 2)]
	return (top.exists && top._type != a._type) || (bottom.exists && bottom._type != a._type)
}

func move(contents [19]amphipod, src, dest int) [19]amphipod {
	copy := contents
	copy[dest] = contents[src]
	copy[src] = amphipod{}
	return copy
}

func (s P1State) next() (nexts []P1State) {

	for source, occupant := range s.contents {
		if occupant.exists {

			// Amphipod in the hallway must go to its destination
			if isHallway(source) && !isDestinationOccupiedByOtherType(s, occupant) {
				bottomReachable, bottomMoves := isReachable(s, source, getDestinationRoom(occupant, 2))
				topReachable, topMoves := isReachable(s, source, getDestinationRoom(occupant, 1))
				if bottomReachable {
					nexts = append(nexts, P1State{
						contents: move(s.contents, source, getDestinationRoom(occupant, 2)),
						cost:     s.cost + cost(occupant)*bottomMoves,
					})
				} else if topReachable {
					nexts = append(nexts, P1State{
						contents: move(s.contents, source, getDestinationRoom(occupant, 1)),
						cost:     s.cost + cost(occupant)*topMoves,
					})
				}
			}

			// Amphipod in a room can go to hallway or destination
			if isRoom(source) && !isInBestSpace(s, occupant, source) {
				for hallwaySpace := 0; hallwaySpace <= 10; hallwaySpace++ {
					reachable, moves := isReachable(s, source, hallwaySpace)
					if reachable && !isJustOutsideRoom(hallwaySpace) {
						nexts = append(nexts, P1State{
							contents: move(s.contents, source, hallwaySpace),
							cost:     s.cost + cost(occupant)*moves,
						})
					}
				}

				if !isDestinationOccupiedByOtherType(s, occupant) {
					bottomReachable, bottomMoves := isReachable(s, source, getDestinationRoom(occupant, 2))
					topReachable, topMoves := isReachable(s, source, getDestinationRoom(occupant, 1))
					if bottomReachable {
						nexts = append(nexts, P1State{
							contents: move(s.contents, source, getDestinationRoom(occupant, 2)),
							cost:     s.cost + cost(occupant)*bottomMoves,
						})
					} else if topReachable {
						nexts = append(nexts, P1State{
							contents: move(s.contents, source, getDestinationRoom(occupant, 1)),
							cost:     s.cost + cost(occupant)*topMoves,
						})
					}
				}
			}
		}
	}

	return nexts
}
