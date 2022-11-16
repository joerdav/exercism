package prime

func Factors(n int64) []int64 {
	pf := []int64{}
	div := int64(2)
	for n != 1 {
		if n%div != 0 {
			div++
			continue
		}
		n = n / div
		pf = append(pf, div)
	}
	return pf
}
