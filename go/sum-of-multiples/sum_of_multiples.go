// Package summultiples sums multiples
package summultiples

// SumMultiples sums multiples of a list of numbers up to a given limit
func SumMultiples(limit int, divisors ...int) (total int) {
	filter := make([]bool, limit)
	for _, divisor := range divisors {
		for value := divisor; value < limit; value += divisor {
			filter[value] = true
		}
	}
	for value, add := range filter {
		if add {
			total += value
		}
	}
	return
}
