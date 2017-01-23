// Package house generates a song about a house that Jack built
package house

import (
	"fmt"
	"strings"
)

var testVersion = 1

type part struct {
	name   string
	action string
}

var parts = []part{
	{`malt`, `lay in`},
	{`rat`, `ate`},
	{`cat`, `killed`},
	{`dog`, `worried`},
	{`cow with the crumpled horn`, `tossed`},
	{`maiden all forlorn`, `milked`},
	{`man all tattered and torn`, `kissed`},
	{`priest all shaven and shorn`, `married`},
	{`rooster that crowed in the morn`, `woke`},
	{`farmer sowing his corn`, `kept`},
	{`horse and the hound and the horn`, `belonged to`},
}

// Verse generates a single verse of the song
func Verse(num int) string {
	verse := `This is `
	for x := num - 2; x >= 0; x-- {
		verse += fmt.Sprintf("the %s\nthat %s ", parts[x].name, parts[x].action)
	}
	verse += fmt.Sprintf(`the house that Jack built.`)
	return verse
}

// Song generates the whole song
func Song() string {
	song := make([]string, 12, 12)
	for x := 0; x < 12; x++ {
		song[x] = Verse(x + 1)
	}
	return strings.Join(song, "\n\n")
}
