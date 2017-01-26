// Package palindrome calculates palindrome products
package palindrome

import "fmt"

var testVersion = 1

// Product represents a product and its factorisations
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func isPalindrome(i int) bool {
	if i%10 == 0 {
		// If the number ends with 0 it can't be palindromic
		return false
	}

	// Incrementally pop digits from n and push onto r
	n := i
	r := 0
	for n > 0 {
		d := n % 10
		r = r*10 + d
		n = n / 10
	}

	// If the reversed number equals the original, this is palandromic
	return i == r
}

// Products returns various products that are palindromic
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax...")
		return
	}

	var min, max int
	factorisations := make(map[int][][2]int)

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
				fs, ok := factorisations[product]
				if !ok {
					fs = make([][2]int, 0)
				}
				factorisations[product] = append(fs, [2]int{x, y})
			}
		}
	}

	if min == 0 {
		err = fmt.Errorf("No palindromes...")
	}

	pmin = Product{Product: min, Factorizations: factorisations[min]}
	pmax = Product{Product: max, Factorizations: factorisations[max]}
	return
}
