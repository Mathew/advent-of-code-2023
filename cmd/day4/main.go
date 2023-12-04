package main

import (
	"log"
	"slices"
)

func main() {
	partOne()
}

type Game struct {
	id      int
	results []int
	card    []int
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

func partOne() {
	TotalScore := 0

	for _, game := range INPUT {
		scorer := Scorer()
		finalScore := 0

		for _, r := range game.results {
			if slices.Contains(game.card, r) {
				finalScore = scorer()
			}
		}

		TotalScore += finalScore
	}

	log.Printf("Day 4, part 1: %v", TotalScore)
}
