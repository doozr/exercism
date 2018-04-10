// Package pangram contains functions to check for pangrams.
package pangram

import "strings"

const testVersion = 1

var alphabet = "abcdefghijklmnopqrstuvwxyz"

// IsPangram returns true if the string is a pangram.
func IsPangram(s string) bool {
	lowered := strings.ToLower(s)
	for _, l := range alphabet {
		if !strings.Contains(lowered, string(l)) {
			return false
		}
	}
	return true
}