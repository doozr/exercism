// Package protein allows translation of codons to proteins
package protein

var testVersion = 1

var codonTranslations = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon translates a single codon to a protein
func FromCodon(codon string) (protein string) {
	return codonTranslations[codon]
}

// FromRNA converts an RNA string into a list of proteins
func FromRNA(rna string) (proteins []string) {
	for x := 0; x < len(rna); x += 3 {
		codon := rna[x : x+3]
		protein := FromCodon(codon)
		if protein == "STOP" {
			return
		}
		proteins = append(proteins, protein)
	}
	return
}
