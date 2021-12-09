package day08

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Observation struct {
	patterns []string
	outputs  []string
}

func Run() {
	puzzleInput, err := os.ReadFile("./day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	observations := []Observation{}
	for _, line := range strings.Split(string(puzzleInput), "\n") {
		halves := strings.Split(line, " | ")
		observations = append(observations, Observation{
			patterns: strings.Split(halves[0], " "),
			outputs:  strings.Split(halves[1], " "),
		})
	}

	part1(observations)
	part2(observations)
}

func part1(observations []Observation) {
	count := 0
	for _, observation := range observations {
		for _, output := range observation.outputs {
			if len(output) != 5 && len(output) != 6 {
				count++
			}
		}
	}

	fmt.Printf("Day 8 - Part 1: 1, 4, 7, and 8 appear %v times\n", count)
}

type SegmentSet struct {
	segments uint8
}

func (set SegmentSet) Contains(segment int) bool {
	return set.segments&(1<<segment) != 0
}

func (set *SegmentSet) Insert(segment int) {
	set.segments |= (1 << segment)
}

func (set *SegmentSet) Remove(segment int) {
	set.segments &= ^(1 << segment)
}

func (set SegmentSet) ToSlice() []int {
	segments := []int{}
	for segment := 0; segment < 7; segment++ {
		if set.Contains(segment) {
			segments = append(segments, segment)
		}
	}
	return segments
}

func (set SegmentSet) Size() int {
	count := 0
	for segment := 0; segment < 7; segment++ {
		if set.Contains(segment) {
			count++
		}
	}
	return count
}

func (set *SegmentSet) InitFull() {
	set.segments = 0b1111111
}

func intersection(a SegmentSet, b SegmentSet) SegmentSet {
	return SegmentSet{
		segments: a.segments & b.segments,
	}
}

func union(a SegmentSet, b SegmentSet) SegmentSet {
	return SegmentSet{
		segments: a.segments | b.segments,
	}
}

const ZERO_SEGMENTS = uint8(0b0111111)
const ONE_SEGMENTS = uint8(0b0000110)
const TWO_SEGMENTS = uint8(0b1011011)
const THREE_SEGMENTS = uint8(0b1001111)
const FOUR_SEGMENTS = uint8(0b1100110)
const FIVE_SEGMENTS = uint8(0b1101101)
const SIX_SEGMENTS = uint8(0b1111101)
const SEVEN_SEGMENTS = uint8(0b0000111)
const EIGHT_SEGMENTS = uint8(0b1111111)
const NINE_SEGMENTS = uint8(0b1101111)

func mapToDigit(mapping map[rune]int, signals string) int {
	segments := SegmentSet{}
	for _, signal := range signals {
		segments.Insert(mapping[signal])
	}

	switch segments.segments {
	case ZERO_SEGMENTS:
		return 0
	case ONE_SEGMENTS:
		return 1
	case TWO_SEGMENTS:
		return 2
	case THREE_SEGMENTS:
		return 3
	case FOUR_SEGMENTS:
		return 4
	case FIVE_SEGMENTS:
		return 5
	case SIX_SEGMENTS:
		return 6
	case SEVEN_SEGMENTS:
		return 7
	case EIGHT_SEGMENTS:
		return 8
	case NINE_SEGMENTS:
		return 9
	default:
		return -1
	}
}

func isValidMapping(mapping map[rune]int, patterns []string) bool {
	for _, signals := range patterns {
		if mapToDigit(mapping, signals) == -1 {
			return false
		}
	}
	return true
}

func removeKnown(known map[rune]int, unknown *[7]SegmentSet) {
	for signal, segment := range known {
		unknown[signal-'a'] = SegmentSet{}

		for i := range unknown {
			unknown[i].Remove(segment)
		}
	}
}

func getMapped(potentialMappings [7]SegmentSet) map[rune]int {
	mapping := make(map[rune]int)
	for _, signal := range "abcdefg" {
		if potentialMappings[signal-'a'].Size() == 1 {
			mapping[signal] = potentialMappings[signal-'a'].ToSlice()[0]
		}
	}
	return mapping
}

func totalMappings(potentialMappings [7]SegmentSet) int {
	count := 0
	for _, set := range potentialMappings {
		count += set.Size()
	}
	return count
}

func getAllMappings(known map[rune]int, potentialMappings [7]SegmentSet) []map[rune]int {
	removeKnown(known, &potentialMappings)

	// base cases
	if len(known) == 7 {
		return []map[rune]int{known}
	}

	// Recursive Case
	for _, signal := range "abcdefg" {
		if potentialMappings[signal-'a'].Size() > 0 {
			mappings := []map[rune]int{}

			for _, segment := range potentialMappings[signal-'a'].ToSlice() {
				newKnown, newPotentialMappings := make(map[rune]int), potentialMappings
				for k, v := range known {
					newKnown[k] = v
				}
				newKnown[signal] = segment

				mappings = append(mappings, getAllMappings(newKnown, newPotentialMappings)...)
			}

			return mappings
		}

	}

	return []map[rune]int{}
}

func deduceMappingFrom(patterns []string) map[rune]int {

	potentialMappings := [7]SegmentSet{}
	for i := range potentialMappings {
		potentialMappings[i].InitFull()
	}

	// Use the special cases to eliminate invalid mappings
	for _, signals := range patterns {
		validSegments, invalidSegments := SegmentSet{}, SegmentSet{}
		validSegments.InitFull()
		invalidSegments.InitFull()

		switch len(signals) {
		case 2: // this is a 1
			validSegments, invalidSegments = SegmentSet{ONE_SEGMENTS}, SegmentSet{segments: ^ONE_SEGMENTS}
		case 4: // this is a 4
			validSegments, invalidSegments = SegmentSet{FOUR_SEGMENTS}, SegmentSet{^FOUR_SEGMENTS}
		case 3: // this is a 7
			validSegments, invalidSegments = SegmentSet{SEVEN_SEGMENTS}, SegmentSet{^SEVEN_SEGMENTS}
		}

		for _, signal := range "abcdefg" {
			if strings.ContainsRune(signals, signal) {
				potentialMappings[signal-'a'] = intersection(potentialMappings[signal-'a'], validSegments)
			} else {
				potentialMappings[signal-'a'] = intersection(potentialMappings[signal-'a'], invalidSegments)
			}
		}
	}

	mappings := getAllMappings(getMapped(potentialMappings), potentialMappings)

	for _, mapping := range mappings {
		if isValidMapping(mapping, patterns) {
			return mapping
		}
	}

	return make(map[rune]int)
}

func part2(observations []Observation) {

	sum := 0
	for _, observation := range observations {
		mapping := deduceMappingFrom(observation.patterns)

		output := 0
		for _, signals := range observation.outputs {
			digit := mapToDigit(mapping, signals)
			output *= 10
			output += digit
		}

		sum += output
	}

	fmt.Printf("Day 8 - Part 2: The sum of the outputs is %v\n", sum)
}
