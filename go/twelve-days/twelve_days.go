package twelve

import (
	"bytes"
	"fmt"
	"strings"
)

var numbers = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}
var gifts = [...]string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"a Partridge in a Pear Tree",
}

func Verse(i int) string {
	g := gifts[len(gifts)-i:]
	if len(g) > 1 {
		return fmt.Sprintf("On the %v day of Christmas my true love gave to me: %v, and %v.", numbers[i], strings.Join(g[:len(g)-1], ", "), g[len(g)-1])
	}
	return fmt.Sprintf("On the %v day of Christmas my true love gave to me: %v.", numbers[i], g[0])
}

func Song() string {
	b := new(bytes.Buffer)
	for i := 1; i <= 12; i++ {
		fmt.Fprintln(b, Verse(i))
	}
	return strings.TrimSpace(b.String())
}
