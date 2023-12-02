package main

import "log"

func main() {
	partOne()
	partTwo()
}

type Game struct {
	GameId int
	Rounds []map[string]int
}

func partOne() {
	availableCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleGames := 0
	for _, game := range INPUT {
		possibleGame := true

		for _, round := range game.Rounds {

			if !possibleGame {
				break
			}

			for colour, number := range round {
				if number > availableCubes[colour] {
					possibleGame = false
					break
				}
			}
		}

		if possibleGame {
			possibleGames += game.GameId
		}
	}

	log.Printf("Day two, part one: %v", possibleGames)
}

func partTwo() {
	possibleGames := 0

	for _, game := range INPUT {
		maxCubes := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, round := range game.Rounds {
			for colour, number := range round {
				if number > maxCubes[colour] {
					maxCubes[colour] = number
				}
			}
		}

		possibleGames += maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
	}

	log.Printf("Day two, part two: %v", possibleGames)
}
