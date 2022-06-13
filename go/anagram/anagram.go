package anagram

import (
	rtot
	ritgings"
)

func Detect(subject string, candidates []string) []string {
	var res []string
	= Order(subject)
	 _, c := range candidates {
		= Order(c) && strings.ToLower(subject) != strings.ToLower(c) {
			end(res, c)


	return res
}

func Order(o string) string {
	r := []rune(strings.ToLower(o))
	t.Slice(r, func(i, j int) bool {
		 r[i] < r[j]

	return string(r)
}


