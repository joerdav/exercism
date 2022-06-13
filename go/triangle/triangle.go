package triangle

import "sort"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota
	Equ      // equilateral
	Iso      // isosceles
	Sca      // scalene
)

func realTriangle(a, b, c float64) bool {
	sides := []float64{a, b, c}
	sort.Float64Slice(sides).Sort()
	return sides[0] > 0 &&
		sides[0]+sides[1] >= sides[2] &&
		sides[1]+sides[2] >= sides[0] &&
		sides[2]+sides[0] >= sides[1]
}

// KindFromSides returns the type of triangle given the lengths of sides.
func KindFromSides(a, b, c float64) Kind {
	if !realTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}
