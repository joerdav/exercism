package raindrops

import "fmt"

var drops = []struct {
	num   int
	sound string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	r := ""
	for _, v := range drops {
		if number%v.num == 0 {
			r += v.sound
		}
	}
	if r == "" {
		return fmt.Sprint(number)
	}
	return r
}
