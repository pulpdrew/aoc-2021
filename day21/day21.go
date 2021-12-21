package day21

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	input, err := os.ReadFile("./day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	player1Space, _ := strconv.Atoi(strings.Split(strings.Split(string(input), "\n")[0], ": ")[1])
	player2Space, _ := strconv.Atoi(strings.Split(strings.Split(string(input), "\n")[1], ": ")[1])
	player1Score, player2Score := 0, 0
	nextRoll := 1
	player1Turn := true

	for player1Score < 1000 && player2Score < 1000 {
		if player1Turn {
			player1Space, nextRoll = playPart1(player1Space, nextRoll)
			player1Score += player1Space
		} else {
			player2Space, nextRoll = playPart1(player2Space, nextRoll)
			player2Score += player2Space
		}
		player1Turn = !player1Turn
	}

	if player1Turn {
		fmt.Printf("Day 21 - Part 1: Score of the loser (%v) x Total Rolls (%v) = %v\n", player1Score, nextRoll-1, player1Score*(nextRoll-1))
	} else {
		fmt.Printf("Day 21 - Part 1: Score of the loser (%v) x Total Rolls (%v) = %v\n", player2Score, nextRoll-1, player2Score*(nextRoll-1))
	}

	// Part 2
	player1Space, _ = strconv.Atoi(strings.Split(strings.Split(string(input), "\n")[0], ": ")[1])
	player2Space, _ = strconv.Atoi(strings.Split(strings.Split(string(input), "\n")[1], ": ")[1])
	start := state{
		p1Space:     player1Space,
		p2Space:     player2Space,
		p1Score:     0,
		p2Score:     0,
		player1Turn: true,
	}

	outcomes := countWins(start)
	maxWins := outcomes.p1Wins
	if outcomes.p2Wins > maxWins {
		maxWins = outcomes.p2Wins
	}
	fmt.Printf("Day 21 - Part 2: The player who won the most won in %v universes.\n", maxWins)
}

//////////// Part 1 /////////////

func playPart1(startSpace int, startRoll int) (endSpace, endRoll int) {
	modSafeRoll := startRoll - 1
	modSafeSpace := startSpace - 1

	moves := modSafeRoll%100 + (modSafeRoll+1)%100 + (modSafeRoll+2)%100 + 3
	endSpace = (modSafeSpace+moves)%10 + 1

	return endSpace, startRoll + 3
}

/////////////// Part 2 /////////////////////////

type state struct {
	p1Space, p2Space int
	p1Score, p2Score int
	player1Turn      bool
}

type outcomes struct {
	p1Wins, p2Wins uint64
}

var cache map[state]outcomes = make(map[state]outcomes)

func countWins(start state) (o outcomes) {

	// Base cases
	if start.p1Score >= 21 {
		return outcomes{p1Wins: 1, p2Wins: 0}
	} else if start.p2Score >= 21 {
		return outcomes{p1Wins: 0, p2Wins: 1}
	}

	// Cache case
	cached, isCached := cache[start]
	if isCached {
		return cached
	}

	// Recursive case
	for _, nextState := range generateNextStates(start) {
		nextOutcome := countWins(nextState)
		o.p1Wins += nextOutcome.p1Wins
		o.p2Wins += nextOutcome.p2Wins
	}

	cache[start] = o
	return o
}

func generateNextStates(start state) (nextStates []state) {

	for roll1 := 1; roll1 <= 3; roll1++ {
		for roll2 := 1; roll2 <= 3; roll2++ {
			for roll3 := 1; roll3 <= 3; roll3++ {
				next := start
				next.player1Turn = !start.player1Turn

				moves := roll1 + roll2 + roll3
				if start.player1Turn {
					next.p1Space = (next.p1Space+moves-1)%10 + 1
					next.p1Score += next.p1Space
				} else {
					next.p2Space = (next.p2Space+moves-1)%10 + 1
					next.p2Score += next.p2Space
				}

				nextStates = append(nextStates, next)
			}
		}
	}

	return nextStates
}
