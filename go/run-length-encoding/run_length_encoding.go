package encode

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	if len(input) < 1 {
		return ""
	}
	var (
		r     bytes.Buffer
		curr  rune = []rune(input)[0]
		count int  = 1
	)
	for _, l := range input[1:] {
		if l == curr {
			count++
			continue
		}
		write(&r, curr, count)
		curr = l
		count = 1
	}
	write(&r, curr, count)
	return r.String()
}

func write(w io.Writer, r rune, c int) {
	if c > 1 {
		fmt.Fprint(w, c)
	}
	w.Write([]byte(string([]rune{r})))
}

func RunLengthDecode(input string) string {
	if len(input) < 1 {
		return ""
	}
	var (
		r   bytes.Buffer
		num string
	)
	for _, l := range input {
		if unicode.IsDigit(l) {
			num = num + string([]rune{l})
			continue
		}
		n, _ := strconv.Atoi(num)
		if n == 0 {
			n = 1
		}
		r.WriteString(strings.Repeat(string([]rune{l}), n))
		num = ""
	}
	return r.String()
}
