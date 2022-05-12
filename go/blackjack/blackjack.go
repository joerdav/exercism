package blackjack

var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	val, _ := cardValues[card]
	return val
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1, c2, d := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	sum := c1 + c2
	if c1 == c2 && c1 == 11 {
		return "P"
	}
	if sum == 21 && d < 10 {
		return "W"
	}
	if sum >= 17 || sum >= 12 && d < 7 {
		return "S"
	}
	return "H"
}
