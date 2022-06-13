package clock

import "fmt"

const Day = 1440

// Define the Clock type here.
type Clock int

func New(h, m int) Clock {
	c := Clock(m+h*60) % Day
	if c < 0 {
		c = c + Day
	}
	return c
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
