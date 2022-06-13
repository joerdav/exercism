package romannumerals

import (
	"fmt"
)

var numbers = []struct {
	number  int
	letters string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3000 {
		return "", fmt.Errorf("number should be 0 < x < 3001")
	}
	var r []rune
	for input != 0 {
		for _, n := range numbers {
			if n.number > input {
				continue
			}
			input -= n.number
			r = append(r, []rune(n.letters)...)
			break
		}
	}
	return string(r), nil
}
