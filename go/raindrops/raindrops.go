// Package raindrops counts the raindrops.
package raindrops

import "strconv"

const testVersion = 2

// Convert the number into a raindrop string.
func Convert(x int) string {
	if x%3 == 0 {
		if x%5 == 0 {
			if x%7 == 0 {
				return "PlingPlangPlong"
			}
			return "PlingPlang"
		}
		if x%7 == 0 {
			return "PlingPlong"
		}
		return "Pling"
	}
	if x%5 == 0 {
		if x%7 == 0 {
			return "PlangPlong"
		}
		return "Plang"
	}
	if x%7 == 0 {
		return "Plong"
	}
	return strconv.Itoa(x)
}
