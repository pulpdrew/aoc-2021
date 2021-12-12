package day10

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Run() {
	puzzleInput, err := os.ReadFile("./day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(puzzleInput), "\n")

	syntaxErrorScore := 0
	autocompleteScores := []int{}
	for _, input := range inputs {
		i, r, unmatched := parseUntilFailure(input)
		if i != -1 {
			syntaxErrorScore += scoreRuneP1(r) // Part 1
		} else {
			autocompleteScore := 0
			for cur := unmatched.Front(); cur != nil; cur = cur.Next() {
				autocompleteScore = autocompleteScore*5 + scoreRuneP2(cur.Value.(rune))
			}
			autocompleteScores = append(autocompleteScores, autocompleteScore)
		}
	}

	fmt.Printf("Day 10 - Part 1: syntax error score = %v\n", syntaxErrorScore)

	sort.Ints(autocompleteScores)
	fmt.Printf("Day 10 - Part 2: auto-complete score = %v\n", autocompleteScores[len(autocompleteScores)/2])

}

func parseUntilFailure(input string) (indexOfFailure int, firstFailingRune rune, unmatchedRunes *list.List) {

	stack := list.New()
	for i, r := range input {
		switch r {
		case '(', '[', '<', '{':
			stack.PushFront(r)
		case ')':
			if stack.Front().Value != '(' {
				return i, r, stack
			} else {
				stack.Remove(stack.Front())
			}
		case ']':
			if stack.Front().Value != '[' {
				return i, r, stack
			} else {
				stack.Remove(stack.Front())
			}
		case '>':
			if stack.Front().Value != '<' {
				return i, r, stack
			} else {
				stack.Remove(stack.Front())
			}
		case '}':
			if stack.Front().Value != '{' {
				return i, r, stack
			} else {
				stack.Remove(stack.Front())
			}
		default:
			panic("Invalid rune found")
		}
	}

	return -1, -1, stack
}

func scoreRuneP1(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
}

func scoreRuneP2(r rune) int {
	switch r {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	default:
		return 0
	}
}
