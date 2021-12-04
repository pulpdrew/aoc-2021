package day04

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoBoard struct {
	spaces [25]int
	marked [25]bool
}

func (board *BingoBoard) Get(row int, col int) int {
	return board.spaces[row*5+col]
}

func (board *BingoBoard) Set(row int, col int, value int) {
	board.spaces[row*5+col] = value
}

func (board *BingoBoard) IsMarked(row int, col int) bool {
	return board.marked[row*5+col]
}

func (board *BingoBoard) Mark(row int, col int) {
	board.marked[row*5+col] = true
}

func (board *BingoBoard) MarkValue(value int) {
	for i, space := range board.spaces {
		if value == space {
			board.marked[i] = true
		}
	}
}

func (board *BingoBoard) Print() {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board.IsMarked(row, col) {
				fmt.Printf("[%2v]", board.Get(row, col))
			} else {
				fmt.Printf(" %2v ", board.Get(row, col))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (board *BingoBoard) FromInput(inputs []string) {
	for row, line := range inputs {
		numbers := strings.Split(strings.ReplaceAll(strings.Trim(line, " "), "  ", " "), " ")
		for col, value := range numbers {
			iValue, _ := strconv.ParseInt(value, 10, 64)
			board.Set(row, col, int(iValue))
		}
	}
}

func (board *BingoBoard) HasWon() bool {
	// Check rows and cols
	for i := 0; i < 5; i++ {
		allRowsMarked := true
		allColsMarked := true
		for j := 0; j < 5; j++ {
			if !board.IsMarked(i, j) {
				allColsMarked = false
			}
			if !board.IsMarked(j, i) {
				allRowsMarked = false
			}
		}

		if allColsMarked || allRowsMarked {
			return true
		}
	}

	return false
}

func (board *BingoBoard) Score() int {
	sum := 0
	for i, marked := range board.marked {
		if !marked {
			sum += board.spaces[i]
		}
	}

	return sum
}

func Run() {
	puzzleInput, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(puzzleInput), "\n")

	calledStrings := strings.Split(lines[0], ",")
	calledNumbers := []int{}
	for _, calledString := range calledStrings {
		calledNumber, _ := strconv.Atoi(calledString)
		calledNumbers = append(calledNumbers, calledNumber)
	}

	boards := []BingoBoard{}
	for start := 2; start < len(lines); start += 6 {
		board := BingoBoard{}
		board.FromInput(lines[start : start+5])
		boards = append(boards, board)
	}

	part1(calledNumbers, boards)
	part2(calledNumbers, boards)
}

func part1(calledNumbers []int, boards []BingoBoard) {
	for _, n := range calledNumbers {
		for i := range boards {
			boards[i].MarkValue(n)
		}

		for _, board := range boards {
			if board.HasWon() {
				fmt.Printf("Day 4 - Part 1: Winning Board Score = %v, last number called = %v, product = %v\n", board.Score(), n, n*board.Score())
				return
			}
		}
	}

}

func part2(calledNumbers []int, boards []BingoBoard) {
	boardsLeft := boards
	for _, n := range calledNumbers {
		for i := range boardsLeft {
			boardsLeft[i].MarkValue(n)
		}

		newBoardsLeft := []BingoBoard{}
		for _, board := range boardsLeft {
			if !board.HasWon() {
				newBoardsLeft = append(newBoardsLeft, board)
			}
		}

		if len(newBoardsLeft) == 0 {
			fmt.Printf("Day 4 - Part 2: Last Winning Board Score = %v, last number called = %v, product = %v\n", boardsLeft[0].Score(), n, n*boardsLeft[0].Score())
			return
		}

		boardsLeft = newBoardsLeft
	}

}
