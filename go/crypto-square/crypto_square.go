// Package cryptosquare forms cryptosquares
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

var testVersion = 2

func normalise(s string) (n string) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			n += string(unicode.ToLower(r))
		}
		if unicode.IsNumber(r) {
			n += string(r)
		}
	}
	return
}

func numRows(l uint) (r uint) {
	// Use math.Floor instead of math.Ceil to
	// match behaviour as described in the README.
	// Leave as math.Ceil to match behaviour as
	// expected by the tests.
	r = uint(math.Ceil(math.Sqrt(float64(l))))
	return
}

// Encode a cypher-square
//
// Note that the tests contradict the README
// The README says that c >= r and c - r <= 1
// The tests expect that r >= c and r - c <= 1
//
// The first example, "splunk", expects "su pn lk".
// This has 2 columns and 3 rows. A clear
// contradiction of the README. It should expect
// "sln puk", which has 3 columns and 2 rows.
//
// This implementation follows the tests, not the
// requirements, as is proper. See the comment
// in `func numRows` to change the behaviour.
func Encode(plain string) (encoded string) {
	normalised := normalise(plain)
	length := uint(len(normalised))
	rows := numRows(length)

	var lines []string
	for row := uint(0); row < rows; row++ {
		var line string
		for column := row; column < length; column += rows {
			if column < length {
				line += normalised[column : column+1]
			}
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, " ")
}
