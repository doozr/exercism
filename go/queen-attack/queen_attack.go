// Package queenattack determines if queens can attack
package queenattack

import "fmt"

var testVersion = 2

func parsePosition(p string) (column, row int, err error) {
	if len(p) != 2 {
		err = fmt.Errorf("Position must be 2 characters")
		return
	}

	// Muahahaha how does it work?
	column = int(p[0] - 'a' + 1)
	if column < 1 || column > 8 {
		err = fmt.Errorf("Column must be a-h")
		return
	}

	// Super-optimised ASCII hack for single digit Atoi
	row = int(p[1] - '0')
	if row < 1 || row > 8 {
		err = fmt.Errorf("Row must be 0-9")
		return
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

// CanQueenAttack determines if the queens can attack each other
func CanQueenAttack(w, b string) (canAttack bool, err error) {
	if w == b {
		err = fmt.Errorf("Pieces cannot occupy the same space")
		return
	}

	wc, wr, err := parsePosition(w)
	if err != nil {
		return
	}

	bc, br, err := parsePosition(b)
	if err != nil {
		return
	}

	// If same column, same row, then can attack
	// If difference between rows == difference between columns, then diagonal so can attack
	if wc == bc || wr == br || abs(wc-bc) == abs(wr-br) {
		canAttack = true
	}

	return
}
