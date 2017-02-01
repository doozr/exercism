// Package igpay converts English to Pig Latin
package igpay

import "strings"

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiou", r)
}

func startsWithVowelSound(word string) bool {
	runes := []rune(word)
	return isVowel(runes[0]) ||
		(strings.ContainsRune("xy", runes[0]) && !isVowel(runes[1]))
}

func findFirstVowel(word string) int {
	firstVowel := 0
	for ix, r := range word {
		// Special case u following q
		if r == 'u' && word[ix-1:ix] == "q" {
			continue
		}

		if !isVowel(r) {
			continue
		}

		firstVowel = ix
		break
	}
	return firstVowel
}

// PigLatin converts English to Pig Latin
func PigLatin(input string) (pigLatin string) {
	words := strings.Split(input, " ")
	pigWords := make([]string, 0)
	for _, word := range words {
		if startsWithVowelSound(word) {
			pigWords = append(pigWords, word+"ay")
			continue
		}

		firstVowel := findFirstVowel(word)
		word = word[firstVowel:] + word[0:firstVowel] + "ay"

		pigWords = append(pigWords, word)
	}
	pigLatin = strings.Join(pigWords, " ")
	return
}
