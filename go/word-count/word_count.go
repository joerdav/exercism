package wordcount

import (
	"regexp"
	"strings"
)

var ignored = regexp.MustCompile(`([^a-z0-9'])`)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = ignored.ReplaceAllString(strings.ToLower(phrase), " ")
	r := Frequency{}
	for _, p := range strings.Fields(phrase) {
		r[strings.Trim(p, "'")]++
	}
	return r
}
