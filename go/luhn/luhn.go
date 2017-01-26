// Package luhn implements luhn checks for card validity
package luhn

import (
	"strconv"
	"unicode"
)

var testVersion = 1

// parse a string to extract digits
// spaces are ignored, other characters cause an error
func parse(s string) (digits []int, ok bool) {
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return
		}
		digits = append(digits, digit)
	}
	ok = true
	return
}

// Valid returns true if the provided string passes luhn check
func Valid(s string) (valid bool) {
	digits, ok := parse(s)
	if !ok {
		return
	}

	l := len(digits) - 1
	if l < 1 {
		return
	}

	var sum int
	for x := 0; x <= l; x++ {
		digit := digits[l-x]

		if x%2 == 0 {
			sum += digit
			continue
		}

		digit *= 2
		if digit > 9 {
			// Note: same as adding the digits for numbers 10 <= n <= 19
			digit -= 9
		}
		sum += digit
	}

	return sum%10 == 0
}
