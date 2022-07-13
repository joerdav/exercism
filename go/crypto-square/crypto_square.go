package cryptosquare

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	var b bytes.Buffer
	rs := []rune(normal(pt))
	c, r := rect(len(rs))
	fmt.Println(string(rs))
	fmt.Println(c, r)
	for x := 0; x < c; x++ {
		for y := 0; y < r; y++ {
			idx := y*c + x
			fmt.Println(y*c, x, idx)
			if len(rs) <= idx {
				b.WriteRune(' ')
				continue
			}
			b.WriteRune(rs[idx])
		}
		if x != c-1 {
			b.WriteRune(' ')
		}
	}
	return b.String()
}

func normal(s string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
}

func rect(l int) (int, int) {
	for r := 1; r <= l; r++ {
		for c := 1; c <= l; c++ {
			if r*c >= l && c <= r && c-r <= 1 {
				return r, c
			}
		}
	}
	return 0, 0
}
