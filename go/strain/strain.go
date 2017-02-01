// Package strain contains basic implementations of keep and discard for various types
package strain

// Ints is a list of ints
type Ints []int

// Lists is a list of lists of ints
type Lists []Ints

// Strings is a list of strings
type Strings []string

// Keep ints
func (c Ints) Keep(fn func(int) bool) Ints {
	if c == nil {
		return nil
	}

	result := make(Ints, 0)
	for _, v := range c {
		if !fn(v) {
			continue
		}
		result = append(result, v)
	}
	return Ints(result)
}

// Discard ints
func (c Ints) Discard(fn func(int) bool) Ints {
	if c == nil {
		return nil
	}

	result := make(Ints, 0)
	for _, v := range c {
		if fn(v) {
			continue
		}
		result = append(result, v)
	}
	return Ints(result)
}

// Keep lists
func (c Lists) Keep(fn func([]int) bool) Lists {
	result := make(Lists, 0)
	for _, v := range c {
		if !fn(v) {
			continue
		}
		result = append(result, v)
	}
	return Lists(result)
}

// Keep strings
func (c Strings) Keep(fn func(string) bool) Strings {
	result := make(Strings, 0)
	for _, v := range c {
		if !fn(v) {
			continue
		}
		result = append(result, v)
	}
	return Strings(result)
}
