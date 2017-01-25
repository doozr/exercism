// Package isogram determines if words are isograms
package isogram

import "unicode"

var testVersion = 1

// IsIsogram returns true if the string is an isogram
func IsIsogram(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}

		r = unicode.ToLower(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
