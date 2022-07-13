package prime

import "fmt"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("invalid index")
	}
	if n == 1 {
		return 2, nil
	}
	var curr int
	for n > 0 {
		curr++
		if isPrime(curr) {
			n--
		}
	}
	return curr, nil
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
