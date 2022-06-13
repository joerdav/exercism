package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	found := make(map[rune]bool)
	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r) {
			continue
		}
		found[r] = true
	}
	return len(found) == 26
}
