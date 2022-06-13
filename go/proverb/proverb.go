// Package proverb provides a function for generating a proverb.
package proverb

import "fmt"

// Proverb returns a list of verses given some nouns.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	res := make([]string, len(rhyme), len(rhyme))
	for i := 0; i < len(rhyme)-1; i++ {
		res[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	res[len(res)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return res
}
