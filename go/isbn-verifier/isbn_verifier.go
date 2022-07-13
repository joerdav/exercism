package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	total := 0
	for i, d := range isbn {
		m := 10 - i
		if d == 'X' && i == 9 {
			total += 10 * m
			continue
		}
		if !unicode.IsDigit(d) {
			return false
		}
		n, _ := strconv.Atoi(string([]rune{d}))
		total += n * m
	}
	return total%11 == 0
}
