package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("index out of range")
	}
	if number == 1 {
		return 1, nil
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var n uint64 = 0
	for i := 1; i < 65; i++ {
		s, _ := Square(i)
		n += s
	}
	return n
}
