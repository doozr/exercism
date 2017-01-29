package say

import "strings"

var units = map[uint64]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var weirdTeens = map[uint64]string{
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	15: "fifteen",
	18: "eighteen",
}

var tens = map[uint64]string{
	1: "ten",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var thousands = map[uint64]string{
	1: "thousand",
	2: "million",
	3: "billion",
	4: "trillion",
	5: "quadrillion",
	6: "quintillion",
}

func parseUnits(number uint64) (word string) {
	// Only give a value if there is a value
	if v, ok := units[number]; ok {
		word = v
	}
	return
}

func parseTeens(number uint64) (words string) {
	// Some teen values do not follow the pattern,
	// so have a lookup table for them
	if v, ok := weirdTeens[number]; ok {
		words = v
		return
	}

	// Usual teens are just the unit value with the
	// word "teen" appended to it.
	word := parseUnits(number % 10)
	words = word + "teen"
	return
}

func parseTens(number uint64) (words string) {
	// If this is a teen number, handle it differently
	if number > 10 && number < 20 {
		words = parseTeens(number)
		return
	}

	t := (number / 10) % 10
	u := number % 10
	wordList := make([]string, 0)

	// If any tens, add them
	if t > 0 {
		wordList = append(wordList, tens[t])
	}

	// If any units, add them
	if u > 0 {
		wordList = append(wordList, parseUnits(u))
	}

	words = strings.Join(wordList, " ")
	return
}

func parseHundreds(number uint64) (words string) {
	wordList := make([]string, 0)
	h := (number / 100) % 10
	t := number % 100

	// If there are any hundreds, add them
	if h > 0 {
		wordList = append(wordList, parseUnits(h), "hundred")
	}

	// If there are any tens/units, parse them
	if t > 0 {

		// If there are tens/units AND hundreds, inject
		// a conjoining "and"
		if h > 0 {
			wordList = append(wordList, "and")
		}
		wordList = append(wordList, parseTens(t))
	}

	words = strings.Join(wordList, " ")
	return
}

// Recursively parse groups of thousands
//
// For each thousands-group (thousand=0, million=1, billion=2, etc)
// calculate how many of that group there are in hundreds, tens
// and units. Prepends higher groups.
func parseNumber(thousandGroup, number uint64) (words string) {
	wordList := make([]string, 0)
	t := number % 1000

	// If there is a higher thousands group, append that first
	if number >= 1000 {
		wordList = append(wordList, parseNumber(thousandGroup+1, number/1000))
	}

	// If there is some number of this group, parse it
	if t > 0 {
		wordList = append(wordList, parseHundreds(t))

		// Check that the group has a name (numbers < 1000
		// do not have a group name)
		if w, ok := thousands[thousandGroup]; ok {
			wordList = append(wordList, w)
		}
	}

	words = strings.Join(wordList, " ")
	return
}

// Say numbers as words
//
// Implements the conjoining "and" between hundreds and
// tens/units. Also ditches the hyphen between tens and
// units. Tests updated appropriately.
func Say(number uint64) (words string) {
	if number == 0 {
		words = parseUnits(0)
		return
	}

	return parseNumber(0, number)
}
