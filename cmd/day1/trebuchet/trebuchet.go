package trebuchet

import (
	"strconv"
	"unicode"
)

type Matcher interface {
	Matches(rune) bool
}

type DigitMatcher struct{}

func (dm DigitMatcher) Matches(r rune) bool {
	return unicode.IsDigit(r)
}

func First(str string, matcher Matcher) (bool, rune) {
	rs := []rune(str)

	l := len(rs)
	for x := 0; x < l; x++ {
		if matcher.Matches(rs[x]) {
			return true, rs[x]
		}
	}

	return false, rune(0)
}

func Last(str string, matcher Matcher) (bool, rune) {
	rs := []rune(str)

	l := len(rs)
	for x := l - 1; x >= 0; x-- {
		if matcher.Matches(rs[x]) {
			return true, rs[x]
		}
	}

	return false, rune(0)
}

func CombineRunes(a, b rune) string {
	return string([]rune{a, b})
}

func ConvertStrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic("trebuchet: String to Int Conversion")
	}

	return i
}

func Sum(is ...int) int {
	sum := 0
	for _, i := range is {
		sum += i
	}

	return sum
}
