package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands do not match.")
	}
	var count int
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
