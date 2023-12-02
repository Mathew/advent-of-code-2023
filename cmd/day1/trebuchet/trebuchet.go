package trebuchet

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Matcher interface {
	Matches(rune) (bool, string)
}

type DigitMatcher struct{}

func NewDigitMatcher() Matcher {
	return DigitMatcher{}
}

func (dm DigitMatcher) Matches(r rune) (bool, string) {
	return unicode.IsDigit(r), string(r)
}

type WordConverter struct {
	wordsMapping map[string]string
	matcher      Matcher
}

func NewWordConverter(wordsMapping map[string]string, matcher Matcher) Matcher {
	return WordConverter{wordsMapping, matcher}
}

func (wm WordConverter) Matches(r rune) (bool, string) {
	ok, val := wm.matcher.Matches(r)
	if ok {
		val = wm.wordsMapping[val]
	}

	return ok, val
}

type WordMatcher struct {
	context string
	words   []string
	reverse bool
}

func NewWordMatcher(words []string) Matcher {
	wm := WordMatcher{"", words, false}
	return &wm
}

func NewReverseWordMatcher(words []string) Matcher {
	wm := WordMatcher{"", words, true}
	return &wm

}

func (wm *WordMatcher) Matches(r rune) (bool, string) {
	if unicode.IsDigit(r) {
		wm.context = ""
		return false, ""
	}

	if wm.reverse {
		wm.context = string(r) + wm.context
	} else {
		wm.context = wm.context + string(r)
	}

	for _, w := range wm.words {
		if w == wm.context {
			match := wm.context
			wm.context = ""
			return true, match
		}

		if strings.Contains(w, wm.context) {
			return false, ""
		}
	}
	wm.context = ""
	return false, ""
}

type MultiMatcher struct {
	words  Matcher
	digits Matcher
}

func NewMultiMatcher(words Matcher, digits Matcher) Matcher {
	mm := MultiMatcher{words, digits}
	return &mm
}

func (fwodm *MultiMatcher) Matches(r rune) (bool, string) {
	ok, val := fwodm.words.Matches(r)
	if ok {
		return true, val
	}

	ok, val = fwodm.digits.Matches(r)
	if ok {
		return true, val
	}

	return false, ""
}

func First(str string, matcher Matcher) (bool, string) {
	rs := []rune(str)

	l := len(rs)
	for x := 0; x < l; x++ {
		ok, val := matcher.Matches(rs[x])
		if ok {
			return true, val
		}
	}

	return false, ""
}

func Last(str string, matcher Matcher) (bool, string) {
	rs := []rune(str)

	l := len(rs)
	for x := l - 1; x >= 0; x-- {
		ok, val := matcher.Matches(rs[x])
		if ok {
			return true, val
		}
	}

	return false, ""
}

func ConvertStrToInt(str string) int {
	println("converting:")
	println(str)
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("%v", err)
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
