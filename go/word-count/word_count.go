package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Frequency map of words
type Frequency map[string]int

// WordCount creates a frequency map
//
// Strictly speaking the tests only require apostrophe to be
// handled differently, but we can support hyphens too at little
// extra cost
//
// Added two extra test cases:
//
// {
// 	"with hyphenation",
// 	"I joined a co-operative",
// 	Frequency{"i": 1, "joined": 1, "a": 1, "co-operative": 1},
// },
// {
// 	"with dangling hyphens",
// 	"I have -redacted- this statement",
// 	Frequency{"i": 1, "have": 1, "redacted": 1, "this": 1, "statement": 1},
// },
func WordCount(phrase string) Frequency {
	frequency := Frequency{}
	inWord := false
	wordStart := 0
	specials := "'-"

	// Function to extract a word and add it to the frequency map
	addWord := func(ix int) {
		// Trim punctuation from the ends of words
		word := strings.Trim(phrase[wordStart:ix], specials)
		frequency[word]++
	}

	// Normalise the phrase once
	phrase = strings.ToLower(phrase)

	for ix, r := range phrase {
		// Allow letters, numbers and certain special punctuation
		isWordChar := unicode.IsLetter(r) ||
			unicode.IsDigit(r) ||
			strings.ContainsRune(specials, r)

		// If it is a word char start a word
		if isWordChar && !inWord {
			wordStart = ix
			inWord = true
		}

		// If not a word char terminate the current word
		if !isWordChar && inWord {
			inWord = false
			addWord(ix)
		}
	}

	// Catch the last word
	if inWord {
		addWord(len(phrase))
	}

	return frequency
}
