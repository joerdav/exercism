package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var nonDigits = regexp.MustCompile(`[^\d]`)

type phoneNumber struct {
	areaCode, firstChunk, secondChunk string
}

func (p phoneNumber) Number() string {
	return p.areaCode + p.firstChunk + p.secondChunk
}

func (p phoneNumber) Format() string {
	return fmt.Sprintf("(%s) %s-%s", p.areaCode, p.firstChunk, p.secondChunk)
}

func parse(number string) (phoneNumber, error) {
	number = nonDigits.ReplaceAllString(number, "")
	if len(number) == 11 && number[0] == '1' {
		number = string(number[1:])
	}
	if len(number) != 10 {
		return phoneNumber{}, fmt.Errorf("wrong digit len, len(%s) == %d", number, len(number))
	}
	if number[0] < '2' || number[3] < '2' {
		return phoneNumber{}, errors.New("first and fourth digits should be 1 < n < 10")
	}
	return phoneNumber{number[:3], number[3:6], number[6:]}, nil
}

func Number(phoneNumber string) (string, error) {
	n, err := parse(phoneNumber)
	if err != nil {
		return "", err
	}
	return n.Number(), nil
}

func AreaCode(phoneNumber string) (string, error) {
	n, err := parse(phoneNumber)
	if err != nil {
		return "", err
	}
	return n.areaCode, nil
}

func Format(phoneNumber string) (string, error) {
	n, err := parse(phoneNumber)
	if err != nil {
		return "", err
	}
	return n.Format(), nil
}
