package perfect

import (
	"errors"
)

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var (
	ErrOnlyPositive = errors.New("input must be positive")
)

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}
	a := aliquot(n)
	switch {
	case a == n:
		return ClassificationPerfect, nil
	case a > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}

func aliquot(n int64) int64 {
	var a int64
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			a += i
		}
	}
	return a
}
