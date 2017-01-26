// Package palindrome calculates palindrome products
package palindrome

import (
	"fmt"
	"strconv"
)

var testVersion = 1

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func isPalindrome(i int) bool {
	s := strconv.Itoa(i)
	length := len(s)
	halfway := length / 2
	end := length - 1
	for ix := 0; ix <= halfway; ix++ {
		if s[ix] != s[end-ix] {
			return false
		}
	}
	return true
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax...")
		return
	}

	var min, max int
	factors := make(map[int][][2]int)

	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			product := x * y

			if !isPalindrome(product) {
				continue
			}

			if product < min || min == 0 {
				min = product
			}

			if product > max {
				max = product
			}

			if product == min || product == max {
				fs, ok := factors[product]
				if !ok {
					fs = make([][2]int, 0)
				}
				factors[product] = append(fs, [2]int{x, y})
			}
		}
	}

	if min == 0 {
		err = fmt.Errorf("No palindromes...")
	}

	pmin = Product{Product: min, Factorizations: factors[min]}
	pmax = Product{Product: max, Factorizations: factors[max]}
	return
}
