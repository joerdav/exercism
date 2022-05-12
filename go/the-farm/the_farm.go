package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError int

func (s SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s)
}

var (
	ErrNegativeFodder = errors.New("negative fodder")
	ErrDivideByZero   = errors.New("division by zero")
)

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}
	if fodder < 0 {
		return 0, ErrNegativeFodder
	}
	if err == ErrScaleMalfunction {
		fodder *= 2
	}
	if cows == 0 {
		return 0, ErrDivideByZero
	}
	if cows < 0 {
		return 0, SillyNephewError(cows)
	}
	return fodder / float64(cows), nil
}
