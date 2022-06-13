package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.Replace(id, " ", "", -1)
	if len(id) <= 1 {
		return false
	}
	var sum int64
	double := false
	numbers := strings.Split(id, "")
	for i := len(numbers) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(numbers[i], 10, 8)
		if err != nil {
			return false
		}
		if double == true {
			n = 2 * n
			if n > 9 {
				n = n - 9
			}
		}
		sum += n
		double = !double
	}
	return sum%10 == 0
}
