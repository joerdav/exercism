package isogram

import "strings"

func IsIsogram(word string) bool {
	found := map[rune]bool{}
	for _, l := range strings.ToLower(word) {
		if l == '-' || l == ' ' {
			continue
		}
		if found[l] {
			return false
		}
		found[l] = true
	}
	return true
}
