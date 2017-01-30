package diamond

import (
	"fmt"
	"strings"
)

var testVersion = 1

// Gen creates a diamond of letters
func Gen(b byte) (diamond string, err error) {
	endRune := rune(b)
	if endRune < 'A' || endRune > 'Z' {
		err = fmt.Errorf("Invalid rune")
		return
	}

	if endRune == 'A' {
		diamond = "A\n"
		return
	}

	lines := []string{}

	outerSpaces := int(endRune) - 65
	innerSpaces := -1
	for r := 'A'; r <= endRune; r++ {
		var line string

		if innerSpaces > 0 {
			line = fmt.Sprintf("%[1]*[2]c%[3]*c%-[1]*[2]c\n", outerSpaces+1, r, innerSpaces, ' ')
		} else {
			line = fmt.Sprintf("%[1]*[2]c%[3]*[4]c\n", outerSpaces+1, r, outerSpaces, ' ')
		}

		lines = append(lines, string(line))
		outerSpaces--
		innerSpaces += 2
	}

	for x := len(lines) - 2; x >= 0; x-- {
		lines = append(lines, lines[x])
	}

	diamond = strings.Join(lines, "")
	return
}
