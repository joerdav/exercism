package cipher

import (
	"bytes"
	"strings"
	"unicode"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift int

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return -1
		}
		r = unicode.ToLower(r)
		return letters[(((int(r-'a') + int(c)) + 26) % 26)]
	}, input)
}

func (c shift) Decode(input string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return -1
		}
		r = unicode.ToLower(r)
		return letters[((int(r-'a')-int(c))+26)%26]
	}, input)
}

type vigenere []int

func NewVigenere(key string) Cipher {
	if len(key) < 3 || strings.ToLower(key) != key {
		return nil
	}
	key = strings.ToLower(key)
	var v vigenere
	for _, r := range key {
		if !unicode.IsLetter(r) {
			return nil
		}
		v = append(v, int(r-'a'))
	}
	return v
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	var b bytes.Buffer
	var i int
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		c := int(r - 'a')
		ru := (v[i%len(v)] + c) % 26
		b.WriteRune(rune(ru + 'a'))
		i++

	}
	return b.String()
}

func (v vigenere) Decode(input string) string {
	input = strings.ToLower(input)
	var b bytes.Buffer
	var i int
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		c := int(r - 'a')
		ru := (c + 26 - v[i%len(v)]) % 26
		b.WriteRune(rune(ru + 'a'))
		i++

	}
	return b.String()
}
