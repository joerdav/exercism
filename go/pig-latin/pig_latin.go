package piglatin

import (
	"strings"
	"unicode"
)

func Sentence(sentence string) string {
	ws := strings.Fields(sentence)
	for i := range ws {
		ws[i] = word(ws[i])
	}
	return strings.Join(ws, " ")
}

func word(word string) string {
	rs := []rune(word)
	if isVowel(rs[0]) || strings.HasPrefix(strings.ToLower(word), "yt") || strings.HasPrefix(strings.ToLower(word), "xr") {
		return word + "ay"
	}
	s, e := splitConsonantGroups(word)
	return e + s + "ay"
}

func splitConsonantGroups(s string) (string, string) {
	sl := strings.ToLower(s)
	if strings.HasPrefix(sl, "qu") {
		return s[:2], s[2:]
	}
	clusterlen := 0
	for i, l := range sl {
		if isVowel(l) || (l == 'y' && i != 0) {
			break
		}
		if strings.HasPrefix(sl[i:], "qu") {
			clusterlen += 2
			break
		}
		clusterlen++
	}
	return s[:clusterlen], s[clusterlen:]
}

func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
