package allyourbase

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrInvalidBase  = errors.New("invalid base")
	ErrInvalidDigit = errors.New("invalid digit")
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 || outputBase < 2 {
		return nil, ErrInvalidBase
	}
	var res int
	for i, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, ErrInvalidDigit
		}
		res += d * int(math.Pow(float64(outputBase), float64(len(inputDigits)-(i+1))))
	}
	return intToSlice(res), nil
}
func intToSlice(i int) []int {
	var s []int
	for _, d := range fmt.Sprint(i) {
		s = append(s, int(d-'0'))
	}
	return s
}
