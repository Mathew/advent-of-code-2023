package main

import (
	"log"
	"slices"
)

func main() {
	partOne()
	partTwo()
}

type Game struct {
	id      int
	results []int
	card    []int
}

func (g Game) Matches() []int {
	matches := []int{}

	for _, r := range g.results {
		if slices.Contains(g.card, r) {
			matches = append(matches, r)
		}
	}

	return matches

}

func Scorer() func() int {
	x := 0

	return func() int {
		if x == 0 {
			x = 1
			return x
		}

		x += x
		return x
	}
}

func xrange(start, end, limit int) []int {
	arr := []int{}

	for i := start; i < end+1; i++ {
		if i > limit {
			break
		}

		arr = append(arr, i)
	}

	return arr

}

func partOne() {
	TotalScore := 0

	for _, game := range INPUT {
		scorer := Scorer()
		finalScore := 0

		for range game.Matches() {
			finalScore = scorer()
		}

		TotalScore += finalScore
	}

	log.Printf("Day 4, part 1: %v", TotalScore)
}

func partTwo() {
	scratchCardMap := map[int]int{}
	scratchCardCount := map[int]int{}
	total := 0

	for _, game := range INPUT {
		scratchCardMap[game.id] = len(game.Matches())
		scratchCardCount[game.id] += 1

		total += scratchCardCount[game.id]

		for _, id := range xrange(game.id+1, game.id+scratchCardMap[game.id], 207) {
			scratchCardCount[id] += 1 * scratchCardCount[game.id]
		}
	}

	log.Printf("Day 4, part 2: %v", total)
}
