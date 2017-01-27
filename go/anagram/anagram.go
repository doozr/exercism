package anagram

import (
	"sort"
	"strings"
)

// RuneSorter sorts arrays of runs into lexicographic order
type RuneSorter []rune

func (rs RuneSorter) Less(x, y int) bool {
	return rs[x] < rs[y]
}

func (rs RuneSorter) Swap(x, y int) {
	rs[x], rs[y] = rs[y], rs[x]
}

func (rs RuneSorter) Len() int {
	return len(rs)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(RuneSorter(r))
	return string(r)
}

// Detect anagrams in a list of candidates
func Detect(anagram string, candidates []string) []string {
	anagram = strings.ToLower(anagram)
	sortedAnagram := sortString(anagram)

	matches := make([]string, 0)

	for _, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		if anagram == candidate || len(anagram) != len(candidate) {
			continue
		}

		sortedCandidate := sortString(candidate)
		if sortedAnagram != sortedCandidate {
			continue
		}

		matches = append(matches, candidate)
	}
	return matches
}
