package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var ints []int
	if len(digits) < span {
		return 0, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return -1, errors.New("span must not be nagative")
	}
	if len(digits) == 0 || span == 0 {
		return 1, nil
	}
	for _, r := range digits {
		i, err := strconv.Atoi(string([]rune{r}))
		if err != nil {
			return 0, errors.New("digits input must only contain digits")
		}
		ints = append(ints, i)
	}
	var max int
	for i := 0; i <= len(digits)-span; i++ {
		p := product(ints[i : i+span])
		if p > max {
			max = p
		}

	}
	return int64(max), nil
}

func product(is []int) int {
	res := is[0]
	for _, i := range is[1:] {
		res *= i
	}
	return res
}
