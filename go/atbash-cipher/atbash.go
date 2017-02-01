// Package atbash contains an atbash implementation
package atbash

import "strings"

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isBoundary(i int) bool {
	return i%5 == 4
}

// Atbash encodes a string using the atbash cipher
func Atbash(plain string) (cipher string) {
	plain = strings.ToLower(plain)
	runes := make([]rune, 0)

	i := 0
	for _, r := range plain {
		switch {
		case isLetter(r):
			runes = append(runes, 'z'-r+'a')
		case isDigit(r):
			runes = append(runes, r)
		default:
			continue
		}

		if isBoundary(i) {
			runes = append(runes, ' ')
		}
		i++
	}

	cipher = strings.TrimSpace(string(runes))
	return
}
