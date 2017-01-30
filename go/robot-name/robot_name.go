// Package robotname names robots
package robotname

import "math/rand"

var names = []string{}

// Robot with a name
type Robot struct {
	name string
}

// generateAllNames generates every possible name and shuffles the list
func generateAllNames() {
	for a := 'A'; a <= 'Z'; a++ {
		for b := 'A'; b <= 'Z'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					for e := '0'; e <= '9'; e++ {
						runes := []rune{a, b, c, d, e}
						names = append(names, string(runes))
					}
				}
			}
		}
	}
	for i := range names {
		j := rand.Intn(len(names))
		names[i], names[j] = names[j], names[i]
	}
}

// generateName randomly generates a name using rand.
// Not guaranteed random, not guaranteed unique.
//
// To be able to use something like UUIDv4 for generating
// unique names without a lookup table we need a bit more
// entropy than 5 chars in a fixed format.
//
// Instead we generate every possible name, shuffle the list,
// then take them in order.
//
// 329ms to generate the list, but 485ns for every allocation
// after that. Worth the initial time, I think.
func generateName() (name string) {
	if len(names) == 0 {
		generateAllNames()
	}
	name = names[0]
	names = names[1:]
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
