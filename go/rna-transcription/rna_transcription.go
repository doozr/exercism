// Package strand translates DNA and RNA strands
package strand

var testVersion = 3

var translation = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA translates a DNA sequence to RNA
func ToRNA(input string) (rna string) {
	nucleotides := make([]rune, len(input))
	for ix, r := range input {
		nucleotides[ix] = translation[r]
	}
	rna = string(nucleotides)
	return
}
