// Package phonenumber deals with American phone numbers
package phonenumber

import "fmt"

var testVersion = 1

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// Number validator
func Number(input string) (number string, err error) {
	digits := make([]rune, 0)
	for _, r := range input {
		if !isDigit(r) {
			continue
		}
		digits = append(digits, r)
	}

	if len(digits) == 11 {
		if digits[0] != '1' {
			err = fmt.Errorf("Invalid national code")
			return
		}
		digits = digits[1:]
	}

	if len(digits) != 10 {
		err = fmt.Errorf("Incorrect number length")
		return
	}

	number = string(digits)
	return
}

// AreaCode of the number
func AreaCode(input string) (areaCode string, err error) {
	input, err = Number(input)
	if err != nil {
		return
	}

	areaCode = input[0:3]
	return
}

// Format the phone number correctly
func Format(input string) (number string, err error) {
	input, err = Number(input)
	if err != nil {
		return
	}

	number = fmt.Sprintf("(%s) %s-%s", input[0:3], input[3:6], input[6:])
	return
}
