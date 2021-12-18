package day18

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	lines := strings.Split(string(input), "\n")
	sum := parse(lines[0])
	for i := 1; i < len(lines); i++ {
		left := sum
		right := parse(lines[i])
		sum = add(left, right)
	}

	fmt.Printf("Day 18 - Part 1: The sum of all the numbers has magnitude %v\n", sum.magnitude())

	// Part 2
	largestMagnitude := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}

			s1, s2 := parse(lines[i]), parse(lines[j])
			sum := add(s1, s2)
			mag := sum.magnitude()

			if mag > largestMagnitude {
				largestMagnitude = mag
			}
		}
	}

	fmt.Printf("Day 18 - Part 2: The largest magnitude of any two-number sum is %v\n", largestMagnitude)
}

type pair struct {
	left, right, parent *pair
	value               int
}

func parse(s string) (p *pair) {
	p, _ = parseHelper([]rune(s))
	return p
}

func parseHelper(source []rune) (p *pair, remaining []rune) {
	p = new(pair)

	if source[0] == '[' {
		source = source[1:] // [

		p.left, source = parseHelper(source)
		p.left.parent = p
		source = source[1:] // ,

		p.right, source = parseHelper(source)
		p.right.parent = p
		source = source[1:] // ]

		p.value = -1
	} else {
		p = newPair(int(source[0] - '0'))
		source = source[1:]
	}

	return p, source
}

func (p *pair) toString() string {
	if p.left != nil {
		return fmt.Sprintf("[%v,%v]", p.left.toString(), p.right.toString())
	} else {
		return fmt.Sprint(p.value)
	}
}

func newPair(value int) *pair {
	s := new(pair)
	s.value = value
	s.left = nil
	s.right = nil
	s.parent = nil
	return s
}

func add(l, r *pair) *pair {
	s := new(pair)
	s.left = l
	s.right = r
	s.parent = nil
	s.value = -1

	l.parent = s
	r.parent = s

	reduce(s)

	return s
}

func reduce(s *pair) {
	for {

		p := findPairToExplode(s, 0)
		if p != nil {
			p.explode()
			continue
		}

		p = findPairToSplit(s)
		if p != nil {
			p.split()
			continue
		}

		break
	}
}

func findPairToExplode(s *pair, currentDepth int) *pair {
	if s == nil || (currentDepth >= 4 && s.value == -1) {
		return s
	}

	l := findPairToExplode(s.left, currentDepth+1)
	if l != nil {
		return l
	}

	return findPairToExplode(s.right, currentDepth+1)
}

func findPairToSplit(p *pair) *pair {
	if p == nil || p.value >= 10 {
		return p
	}

	l := findPairToSplit(p.left)
	if l != nil {
		return l
	}

	return findPairToSplit(p.right)
}

func (s *pair) split() {
	s.left = newPair(s.value / 2)
	s.left.parent = s

	s.right = newPair(s.value - (s.value / 2))
	s.right.parent = s

	s.value = -1
}

func (s *pair) explode() {
	right, left := s.right.value, s.left.value

	addToSuccessor(s, right)
	addToPredecessor(s, left)

	s.value = 0
	s.left = nil
	s.right = nil
}

func addToSuccessor(original *pair, value int) {
	previous, current := original, original.parent

	for current != nil && current.right == previous {
		previous = current
		current = current.parent
	}

	if current == nil {
		return
	}

	current = current.right

	for current.value == -1 {
		current = current.left
	}

	current.value += value
}

func addToPredecessor(original *pair, value int) {
	previous, current := original, original.parent

	for current != nil && current.left == previous {
		previous = current
		current = current.parent
	}

	if current == nil {
		return
	}

	current = current.left

	for current.value == -1 {
		current = current.right
	}

	current.value += value
}

func (p *pair) magnitude() int {
	if p.value != -1 {
		return p.value
	} else {
		return 3*p.left.magnitude() + 2*p.right.magnitude()
	}
}
