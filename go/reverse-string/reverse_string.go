package reverse

func Reverse(input string) string {
	var b []rune
	for _, r := range input {
		b = append([]rune{r}, b...)
	}
	return string(b)
}
