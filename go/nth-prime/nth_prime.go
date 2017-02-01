package prime

// Not the quickest way to generate primes but we don't know
// how far we'll need to go so we just check each one
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			return false
		}
	}
	return true
}

// Nth gets the Nth prime
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	found := 0
	for p := 2; ; p++ {
		if !isPrime(p) {
			continue
		}

		found++

		if found == n {
			return p, true
		}
	}
}
