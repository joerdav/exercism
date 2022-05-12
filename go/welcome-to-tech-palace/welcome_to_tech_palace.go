package techpalace

import (
	"bytes"
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var b bytes.Buffer
	b.WriteString(strings.Repeat("*", numStarsPerLine) + "\n")
	b.WriteString(welcomeMsg + "\n")
	b.WriteString(strings.Repeat("*", numStarsPerLine))
	return b.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.Trim(strings.Split(oldMsg, "\n")[1], "*"))
}
