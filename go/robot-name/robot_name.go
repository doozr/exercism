// Package robotname names robots
package robotname

import "crypto/rand"

var names = map[string]bool{}

// Robot with a name
type Robot struct {
	name string
}

// These generator functions are horrible. They skew the probability by
// the simple fact that max(byte) is not divisible by either 26 or 10.
// But frankly it doesn't matter because even distribution of characters
// is not a requirement.

func toLetter(b byte) rune {
	return rune(b%26 + 'A')
}

func toDigit(b byte) rune {
	return rune(b%10 + '0')
}

// generateName randomly generates a name using cryptorand
// guaranteed random, not guaranteed unique
//
// To be able to use something like UUIDv4 for generating
// unique names without a lookup table we need a bit more
// entropy than 5 chars in a fixed format.
func generateName() (name string) {
	for name == "" || names[name] {
		bytes := make([]byte, 5)
		runes := make([]rune, 5)
		rand.Read(bytes)

		runes[0] = toLetter(bytes[0])
		runes[1] = toLetter(bytes[1])
		runes[2] = toDigit(bytes[2])
		runes[3] = toDigit(bytes[3])
		runes[4] = toDigit(bytes[4])
		name = string(runes)
	}
	names[name] = true
	return
}

// Name of the robot
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = generateName()
	}
	return r.name
}

// Reset the robot to factory settings
func (r *Robot) Reset() {
	r.name = ""
}
