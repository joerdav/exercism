package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, tiles := range in {
		for _, tile := range tiles {
			result[strings.ToLower(tile)] = score
		}
	}
	return result
}
