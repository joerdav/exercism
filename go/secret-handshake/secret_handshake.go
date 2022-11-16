package secret

const (
	Wink uint = 1 << iota
	DoubleBlink
	CloseEyes
	Jump
	Reverse
)

var actions = []struct {
	mask uint
	text string
}{
	{Wink, "wink"},
	{DoubleBlink, "double blink"},
	{CloseEyes, "close your eyes"},
	{Jump, "jump"},
}

func Handshake(code uint) []string {
	var result []string
	rev := code&Reverse != 0
	for i := 0; i < len(actions); i++ {
		a := actions[i]
		if rev {
			a = actions[(len(actions)-1)-i]
		}
		if a.mask&code != 0 {
			result = append(result, a.text)
		}
	}
	return result
}
