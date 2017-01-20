// Package twelve sings the 12 days of Christmas
package twelve

import (
	"fmt"
	"strings"
)

var testVersion = 1

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

func getGifts(day int) []string {
	if day == 1 {
		return []string{gifts[0]}
	}
	giftsGiven := make([]string, day, day)
	giftsGiven[day-1] = "and " + gifts[0]
	for ix := 0; ix < day-1; ix++ {
		giftsGiven[ix] = gifts[day-1-ix]
	}
	return giftsGiven
}

// Verse gets the verse associated with the given day number
func Verse(day int) (verse string) {
	dayOrdinal := days[day-1]
	giftsGiven := getGifts(day)
	verse = fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s.", dayOrdinal, strings.Join(giftsGiven, ", "))
	return
}

// Song gets all the verses in order
func Song() (song string) {
	for verse := 1; verse <= 12; verse++ {
		song += Verse(verse) + "\n"
	}
	return
}
