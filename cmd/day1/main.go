package main

import (
	"log"

	"github.com/mathew/advent-of-code-2023/cmd/day1/trebuchet"
	"github.com/mathew/advent-of-code-2023/internal/structures"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	results := []int{}
	dm := &trebuchet.DigitMatcher{}
	for _, in := range INPUT {
		_, f := trebuchet.First(in, dm)
		_, l := trebuchet.Last(in, dm)

		results = append(results, trebuchet.ConvertStrToInt(string(f)+string(l)))
	}

	log.Printf("Day 1, part 1 result: %d", trebuchet.Sum(results...))
}

func partTwo() {
	wordsMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	words := structures.Keys(wordsMap)

	reverseWordsMap := map[string]string{
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
	reverseWords := structures.Keys(reverseWordsMap)

	dm := &trebuchet.DigitMatcher{}
	wm := trebuchet.NewWordMatcher(words)
	wc := trebuchet.NewWordConverter(wordsMap, wm)
	fwodm := trebuchet.NewMultiMatcher(wc, dm)

	rwm := trebuchet.NewWordMatcher(reverseWords)
	rwc := trebuchet.NewWordConverter(reverseWordsMap, rwm)
	rfwodm := trebuchet.NewMultiMatcher(rwc, dm)

	results := []int{}
	for _, in := range INPUT {
		_, f := trebuchet.First(in, fwodm)
		_, l := trebuchet.Last(in, rfwodm)

		combined := string(f) + string(l)
		results = append(results, trebuchet.ConvertStrToInt(combined))
	}

	log.Printf("Day 1, part 2 result: %d", trebuchet.Sum(results...))
}
