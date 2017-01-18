// Package bob contains the bot called Bob.
package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func isShouting(s string) (shouting bool) {
	hasLetters := false
	for _, r := range s {
		if unicode.IsLower(r) {
			return
		}
		hasLetters = hasLetters || unicode.IsUpper(r)
	}
	return hasLetters
}

func isQuestion(s string) (question bool) {
	question = strings.HasSuffix(s, "?")
	return
}

// Hey handles Bob's responss to input strings.
func Hey(s string) (response string) {
	s = strings.TrimSpace(s)
	switch {
	case s == "":
		response = "Fine. Be that way!"
	case isShouting(s):
		response = "Whoa, chill out!"
	case isQuestion(s):
		response = "Sure."
	default:
		response = "Whatever."
	}
	return
}
