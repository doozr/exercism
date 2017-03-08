// Package dna contains DNA strand calculations
package dna

import "errors"

var testVersion = 2

// DNA strand
type DNA string

// Histogram of byte keys to int values
type Histogram map[byte]int

func isValidNucleotide(nucleotide byte) (valid bool) {
	switch nucleotide {
	case 'G', 'T', 'A', 'C':
		valid = true
	}
	return
}

// Count the number of times a nucleotide appears in a strand
func (strand DNA) Count(nucleotide byte) (count int, err error) {
	if !isValidNucleotide(nucleotide) {
		err = errors.New("Invalid nucleotide")
		return
	}

	for ix := 0; ix < len(strand); ix++ {
		if strand[ix] == nucleotide {
			count++
		}
	}

	return
}

// Counts returns a histogram of counts of each nucleotide
func (strand DNA) Counts() (histogram Histogram, err error) {
	histogram = make(Histogram)
	var count int
	var total int

	for _, nucleotide := range []byte{'G', 'T', 'A', 'C'} {
		count, _ = strand.Count(nucleotide)
		histogram[nucleotide] = count
		total += count
	}

	if total != len(strand) {
		err = errors.New("Invalid nucleotides in strand")
	}

	return
}
