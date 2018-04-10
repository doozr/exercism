package triangle

import "math"

const testVersion = 3

func sortSides(a, b, c float64) (float64, float64, float64) {
	if a > b && a > c {
		return b, c, a
	}
	if b > a && b > c {
		return a, c, b
	}
	return a, b, c
}

func isBadSide(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}

// KindFromSides calculates the kind of triangle described.
func KindFromSides(a, b, c float64) Kind {
	if isBadSide(a) || isBadSide(b) || isBadSide(c) {
		return NaT
	}

	o, a, h := sortSides(a, b, c)

	if o+a < h {
		return NaT
	}

	if o == a && a == h {
		if o == 0 {
			return NaT
		}
		return Equ
	}

	if o != a && a != h && h != o {
		return Sca
	}

	return Iso
}

// Kind of triangle
type Kind int

const (
	// NaT is not a triangle
	NaT Kind = iota

	// Equ is an equilateral
	Equ

	// Iso is isosceles
	Iso

	// Sca is scalene
	Sca
)