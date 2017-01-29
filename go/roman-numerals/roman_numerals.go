package romannumerals

import "fmt"

var testVersion = 3

func translateThousands(digit int) (numerals []rune) {
	numerals = make([]rune, 0)
	for x := 0; x < digit; x++ {
		numerals = append(numerals, 'M')
	}
	return
}

func translateDigit(digit int, unit, five, ten rune) (numerals []rune) {
	switch digit {
	case 1:
		numerals = []rune{unit}
	case 2:
		numerals = []rune{unit, unit}
	case 3:
		numerals = []rune{unit, unit, unit}
	case 4:
		numerals = []rune{unit, five}
	case 5:
		numerals = []rune{five}
	case 6:
		numerals = []rune{five, unit}
	case 7:
		numerals = []rune{five, unit, unit}
	case 8:
		numerals = []rune{five, unit, unit, unit}
	case 9:
		numerals = []rune{unit, ten}
	}
	return
}

func translateHundreds(digit int) (numerals []rune) {
	return translateDigit(digit, 'C', 'D', 'M')
}

func translateTens(digit int) (numerals []rune) {
	return translateDigit(digit, 'X', 'L', 'C')
}

func translateUnits(digit int) (numerals []rune) {
	return translateDigit(digit, 'I', 'V', 'X')
}

func concat(first, second []rune) []rune {
	if len(second) == 0 {
		return first
	}
	return append(first, second...)
}

// ToRomanNumeral converts arabic numbers to roman numerals
func ToRomanNumeral(arabic int) (roman string, err error) {
	if arabic <= 0 || arabic >= 4000 {
		err = fmt.Errorf("Arabic number should be between 1 and 3999")
		return
	}

	numerals := make([]rune, 0)
	numerals = concat(numerals, translateThousands((arabic/1000)%10))
	numerals = concat(numerals, translateHundreds((arabic/100)%10))
	numerals = concat(numerals, translateTens((arabic/10)%10))
	numerals = concat(numerals, translateUnits(arabic%10))

	roman = string(numerals)
	return
}
