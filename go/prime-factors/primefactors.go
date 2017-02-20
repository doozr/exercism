// Package prime works with prime numbers
package prime

var testVersion = 2

// Factors returns the prime factors of n
func Factors(n int64) []int64 {
	factors := make([]int64, 0)

	// Handle 2 (the only even prime) separately so subsequent
	// checks can increment by 2
	// Borrowed from markwest1's submission:
	// http://exercism.io/submissions/0128f00f48e3400780d828cb7d0c5b25
	for n%2 == 0 {
		n = n / 2
		factors = append(factors, 2)
	}

	for i := int64(3); i*i <= n; i += 2 {
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
