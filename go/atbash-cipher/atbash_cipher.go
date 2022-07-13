package atbash

import (
	"bytes"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var b bytes.Buffer
	var letters int
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			continue
		}
		letters++
		r = unicode.ToLower(r)
		l := r
		if unicode.IsLetter(r) {
			l = 'a' + 25 - (r - 'a')
		}
		b.WriteRune(l)
		if letters%5 == 0 {
			b.WriteRune(' ')
		}
	}
	return strings.TrimSpace(b.String())
}
