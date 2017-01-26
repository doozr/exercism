// Package sieve contains a Sieve of Eratosthenes
package sieve

// Sieve produces primes up to limit using the Sieve of Eratosthenes
func Sieve(limit int) []int {
	marked := make([]bool, limit+1)
	primes := make([]int, 0)
	for x := 2; x <= limit; x++ {
		if marked[x] {
			continue
		}
		primes = append(primes, x)
		for y := x * 2; y <= limit; y += x {
			marked[y] = true
		}
	}
	return primes
}
