// Package stringset implements a set of strings
package stringset

import (
	"fmt"
	"strings"
)

var testVersion = 4

// Set represents a set of strings
type Set struct {
	values *map[string]bool
}

// New instance of Set
func New() Set {
	values := make(map[string]bool)
	return Set{
		values: &values,
	}
}

// NewFromSlice creates a new Set and populates it
func NewFromSlice(ss []string) Set {
	values := make(map[string]bool)
	for _, s := range ss {
		values[s] = true
	}
	return Set{
		values: &values,
	}
}

// String returns a stringified version
func (s Set) String() string {
	ss := make([]string, len(*s.values))
	ix := 0
	for s := range *s.values {
		ss[ix] = fmt.Sprintf("%q", s)
		ix++
	}
	return fmt.Sprintf("{%s}", strings.Join(ss, ", "))
}

// IsEmpty returns true if the set is empty
func (s Set) IsEmpty() bool {
	return len(*s.values) == 0
}

// Has returns true if the set has the requested value
func (s Set) Has(v string) bool {
	return (*s.values)[v]
}

// Add a value to the set
func (s Set) Add(v string) {
	(*s.values)[v] = true
}

// Subset returns true if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	for v := range *s1.values {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 do not have any common values
func Disjoint(s1, s2 Set) bool {
	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}
	for v := range *s1.values {
		if s2.Has(v) {
			return false
		}
	}
	return true
}

// Equal returns true if s1 and s2 have the same values
func Equal(s1, s2 Set) bool {
	if len(*s1.values) != len(*s2.values) {
		return false
	}
	for v := range *s1.values {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

// Intersection returns a set of values that exist in s1 and s2
func Intersection(s1, s2 Set) Set {
	s3 := New()
	for v := range *s1.values {
		if s2.Has(v) {
			s3.Add(v)
		}
	}
	return s3
}

// Difference returns values from s1 that do not exist in s2
func Difference(s1, s2 Set) Set {
	s3 := New()
	for v := range *s1.values {
		if !s2.Has(v) {
			s3.Add(v)
		}
	}
	return s3
}

// Union returns a set that contains all values in s1 and s2
func Union(s1, s2 Set) Set {
	s3 := New()
	for v := range *s1.values {
		s3.Add(v)
	}
	for v := range *s2.values {
		s3.Add(v)
	}
	return s3
}
