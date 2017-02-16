package foodchain

import (
	"fmt"
	"strings"
)

var testVersion = 3

// FoodChain describes the verse structure
type FoodChain struct {
	name        string
	commentary  string
	description string
}

var foodChain = []FoodChain{
	{name: "fly"},
	{name: "spider", commentary: "It wriggled and jiggled and tickled inside her.", description: "that wriggled and jiggled and tickled inside her"},
	{name: "bird", commentary: "How absurd to swallow a bird!"},
	{name: "cat", commentary: "Imagine that, to swallow a cat!"},
	{name: "dog", commentary: "What a hog, to swallow a dog!"},
	{name: "goat", commentary: "Just opened her throat and swallowed a goat!"},
	{name: "cow", commentary: "I don't know how she swallowed a cow!"},
	{name: "horse", commentary: "She's dead, of course!"},
}

// Verse returns one verse
func Verse(num int) (verse string) {
	// Switch to 0 based indexing for the actual verse array
	num--

	var lines []string
	food := foodChain[num]

	lines = append(lines, fmt.Sprintf("I know an old lady who swallowed a %s.", food.name))

	if food.commentary != "" {
		lines = append(lines, food.commentary)
	}

	isLastVerse := num >= len(foodChain)-1
	if isLastVerse {
		verse = strings.Join(lines, "\n")
		return
	}

	for l := num; l > 0; l-- {
		food = foodChain[l]
		prevFood := foodChain[l-1]

		var description = ""
		if prevFood.description != "" {
			description = " " + prevFood.description
		}

		line := fmt.Sprintf("She swallowed the %s to catch the %s%s.", food.name, prevFood.name, description)
		lines = append(lines, line)
	}

	lines = append(lines, "I don't know why she swallowed the fly. Perhaps she'll die.")

	verse = strings.Join(lines, "\n")
	return
}

// Verses returns a range of verses
func Verses(start, end int) string {
	var verses []string
	for x := start; x <= end; x++ {
		verses = append(verses, Verse(x))
	}
	return strings.Join(verses, "\n\n")
}

// Song returns the whole song
func Song() string {
	return Verses(1, len(foodChain))
}
