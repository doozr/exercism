// Package prime works with prime numbers
package prime

var testVersion = 2

// Factors returns the prime factors of n
func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(2); i*i <= n; i++ {
		for n%i == 0 {
			n = n / i
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
