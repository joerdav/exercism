package sieve

func Sieve(limit int) []int {
	marked := map[int]bool{}
	for x := 2; x <= limit; x++ {
		if marked[x] {
			continue
		}
		for y := 2; y <= limit; y++ {
			r := x * y
			if r > limit {
				break
			}
			marked[r] = true
		}
	}
	var primes []int
	for x := 2; x <= limit; x++ {
		if !marked[x] {
			primes = append(primes, x)
		}
	}
	return primes

}
