// Package acronym provides functions for generating acronyms.
package acronym

import "unicode"

const testVersion = 2

// Abbreviate a string to its acronym.
func Abbreviate(s string) string {
	var acronym []rune
	runes := []rune(s)
	var last rune
	for _, r := range runes {
		if (unicode.IsUpper(r) && !unicode.IsUpper(last)) || (unicode.IsLetter(r) && !unicode.IsLetter(last)) {
			acronym = append(acronym, unicode.ToUpper(r))
		}
		last = r
	}
	return string(acronym)
}
