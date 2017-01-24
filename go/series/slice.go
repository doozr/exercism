// Package slice produces slices of an array
package slice

// All returns all substrings of length n
func All(n int, s string) (slices []string) {
	length := len(s)
	if length < n {
		return
	}

	numSlices := length - n + 1
	slices = make([]string, numSlices, numSlices)

	for x := 0; x < numSlices; x++ {
		slices[x] = s[x : x+n]
	}

	return
}

// UnsafeFirst returns the first substring of length n with no bounds checking
func UnsafeFirst(n int, s string) (slice string) {
	slice = s[0:n]
	return
}

// First returns the first substring of length n or sets ok to false
func First(n int, s string) (slice string, ok bool) {
	if len(s) < n {
		ok = false
		return
	}
	slice = s[0:n]
	ok = true
	return
}
