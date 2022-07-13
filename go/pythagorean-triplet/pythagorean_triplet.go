package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var t []Triplet
	for c := min; c <= max; c++ {
		for b := min; b < c; b++ {
			for a := min; a < b; a++ {
				if a*a+b*b == c*c {
					t = append(t, Triplet{a, b, c})
				}
			}
		}
	}
	return t
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var t []Triplet
	for c := 1; c < p; c++ {
		for b := 1; b < c; b++ {
			for a := 1; a < b; a++ {
				if a+b+c == p && a*a+b*b == c*c {
					t = append([]Triplet{{a, b, c}}, t...)
				}
			}
		}
	}
	return t
}
