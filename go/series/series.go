package series

import "fmt"

func All(n int, s string) []string {
	if len(s) < n {
		return nil
	}
	results := make([]string, len(s)-n+1)
	for i := 0; i <= (len(s) - n); i++ {
		fmt.Println(i, i+n, s[i:i+n])
		results[i] = s[i : i+n]
	}
	return results
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
