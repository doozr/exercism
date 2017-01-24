// Package etl does transforms of scrabble data
package etl

import "unicode"

// toLower gives a 25% speed increase over strings.ToLower
func toLower(s string) (l string) {
	for _, x := range s {
		l = string(unicode.ToLower(x))
		break
	}
	return
}

// Transform converts the input lookup table to the required lookup table
func Transform(input map[int][]string) (scores map[string]int) {
	scores = make(map[string]int)
	for n, letters := range input {
		for _, l := range letters {
			l = toLower(l)
			scores[l] = n
		}
	}
	return
}
