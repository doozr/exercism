// Package transpose calculates the transpose of nested arrays
package transpose

func maxLength(input []string) int {
	maxLen := 0
	for ix := range input {
		l := len(input[ix])
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

func makeRuneArrays(input []string) (runes [][]rune) {
	for _, s := range input {
		runes = append(runes, []rune(s))
	}
	return
}

// Transpose a nested array of strings
//
// Works by making slices through the input arrays at
// each "row" position in reverse column order and
// working out if there should be a rune, a space, or
// nothing at each position.
//
// e.g. taking this slice:
//
// ab|c|d
// ef|g|h
// ij
//
// produces the output string "cg". There is no padding
// because g is the last valid char on the line.
func Transpose(input []string) (transpose []string) {
	rows := maxLength(input)
	columns := len(input)
	runes := makeRuneArrays(input)

	for r := 0; r < rows; r++ {
		pad := false
		row := make([]rune, columns)

		// Iterate over columns backwards to aid padding
		for c := columns - 1; c >= 0; c-- {
			switch {
			case r < len(runes[c]):
				// If there is a rune at this position, add it
				row[c] = runes[c][r]
				pad = true
			case pad:
				// otherwise, if a rune has been written, pad the column
				row[c] = ' '
			default:
				// or if no rune written yet, remove the column entirely
				row = row[:c]
			}
		}

		transpose = append(transpose, string(row))
	}

	return
}
