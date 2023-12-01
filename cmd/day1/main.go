package main

import (
	"log"

	"github.com/mathew/advent-of-code-2023/cmd/day1/trebuchet"
)

func main() {

	results := []int{}

	for _, in := range INPUT {
		dm := trebuchet.DigitMatcher{}

		_, f := trebuchet.First(in, dm)
		_, l := trebuchet.Last(in, dm)

		results = append(results, trebuchet.ConvertStrToInt(trebuchet.CombineRunes(f, l)))
	}

	log.Printf("Day 1, part 1 result: %d", trebuchet.Sum(results...))
}
