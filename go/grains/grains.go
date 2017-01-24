// Package grains calculates grains
package grains

import "fmt"

var testVersion = 1

var squares = make([]uint64, 64, 64)

// Square calculates how many grains on a chess board square
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, fmt.Errorf("Square must be 1-64")
	}

	ix := square - 1
	if squares[ix] == 0 {
		if ix == 0 {
			squares[ix] = 1
		} else {
			v, _ := Square(ix)
			squares[ix] = v << 1
		}
	}

	return squares[ix], nil
}

// Total calculates the total number of grains on a board
func Total() uint64 {
	var t uint64
	for s := 1; s <= 64; s++ {
		st, _ := Square(s)
		t = t + st
	}
	return t
}
