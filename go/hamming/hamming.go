// Package hamming calculates the hamming difference between two DNA strands.
package hamming

import "fmt"

const testVersion = 5

// Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Mismatched lengths")
	}

	differences := 0
	for ix := range a {
		if a[ix] != b[ix] {
			differences += 1
		}
	}

	return differences, nil
}
