package ocr

import (
	"strconv"
	"strings"
)

var digits = []string{`
 _
| |
|_|
`, `

  |
  |
`, `
 _
 _|
|_
`, `
 _
 _|
 _|
`, `

|_|
  |
`, `
 _
|_
 _|
`, `
 _
|_
|_|
`, `
 _
  |
  |
`, `
 _
|_|
|_|
`, `
 _
|_|
 _|
`,
}

// recognizeDigit compares
func recognizeDigit(input string) (digit string) {
	lines := strings.Split(input, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " ")
	}
	input = "\n" + strings.Join(lines, "\n")
	for i, d := range digits {
		if input == d {
			digit = strconv.Itoa(i)
			return
		}
	}
	return "?"
}

// minLen returns the length of the shortest of an array of strings
func minLen(input []string) int {
	m := 0
	for _, s := range input {
		if m == 0 || len(s) < m {
			m = len(s)
		}
	}
	return m
}

// splitDigits splits an array of 4 lines into an array of 3x4 digits
func splitDigits(input []string) (digits []string) {
	length := minLen(input)

	for i := 0; i <= length-3; i += 3 {
		var lines []string
		for _, line := range input {
			lines = append(lines, line[i:i+3])
		}
		digits = append(digits, strings.Join(lines, "\n"))
	}
	return
}

// splitDigitLines splits a string into groups of 4 lines
func splitDigitLines(input string) (digitLines [][]string) {
	// Split into individual lines
	lines := strings.Split(input, "\n")

	// Strip leading blanks
	for len(lines[0]) == 0 {
		lines = lines[1:]
	}

	for i := 0; i <= len(lines)-4; i += 4 {
		digitLines = append(digitLines, lines[i:i+4])
	}
	return
}

// Recognize digits in a string
func Recognize(input string) (output []string) {
	for _, digitLine := range splitDigitLines(input) {
		var digits string
		for _, digit := range splitDigits(digitLine) {
			digits = digits + recognizeDigit(digit)
		}
		output = append(output, digits)
	}
	return
}
