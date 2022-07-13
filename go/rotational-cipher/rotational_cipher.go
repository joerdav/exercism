package rotationalcipher

import (
	"strings"
	"unicode"
)

const characters = "abcdefghijklmnopqrstuvwxyz"

func RotationalCipher(plain string, shiftKey int) string {
	cypher := createMap(shiftKey)
	return strings.Map(func(r rune) rune {
		l, ok := cypher[unicode.ToLower(r)]
		if !ok {
			return r
		}
		if unicode.IsUpper(r) {
			return unicode.ToUpper(l)
		}
		return l

	}, plain)
}

func createMap(sk int) map[rune]rune {
	result := map[rune]rune{}
	for i, l := range characters {
		result[l] = []rune(characters)[(i+sk)%26]
	}
	return result
}
