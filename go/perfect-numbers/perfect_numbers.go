// Package perfect classifies numbers
package perfect

import "errors"

var testVersion = 1

// Classification of numbers
type Classification int

// ErrOnlyPositive is returned if a number less than 1 is passed in
var ErrOnlyPositive = errors.New("Only positive numbers can be calculated")

const (
	// ClassificationDeficient represents deficient numbers
	ClassificationDeficient = iota

	// ClassificationPerfect represents perfect numbers
	ClassificationPerfect

	// ClassificationAbundant represents abundant numbers
	ClassificationAbundant
)

func sumFactors(n uint64) uint64 {
	sum := uint64(1)

	f := uint64(2)
	for f*f < n {
		if n%f == 0 {
			sum += f + n/f
		}
		f++
	}
	if f*f == n {
		sum += f
	}
	return sum
}

// Classify a number as deficient, perfect or abundant
func Classify(n uint64) (c Classification, err error) {
	if n <= 0 {
		err = ErrOnlyPositive
		return
	}

	if n == 1 {
		c = ClassificationDeficient
		return
	}

	sum := sumFactors(n)
	switch {
	case sum < n:
		c = ClassificationDeficient
	case sum > n:
		c = ClassificationAbundant
	default:
		c = ClassificationPerfect
	}
	return
}
