package beer

import (
	"fmt"
	"strconv"
	"strings"
)

var testVersion = 1

func firstLine(n int) (line string) {
	format := "%[1]s bottle%[2]s of beer on the wall, %[1]s bottle%[2]s of beer.\n"
	switch n {
	case 1:
		line = fmt.Sprintf(format, "1", "")
	case 0:
		line = fmt.Sprintf(format, "no more", "s")
		line = "N" + line[1:]
	default:
		line = fmt.Sprintf(format, strconv.Itoa(n), "s")
	}
	return
}

func secondLine(n int) (line string) {
	format := "Take %[1]s down and pass it around, %[2]s bottle%[3]s of beer on the wall.\n"
	switch n {
	case 1:
		line = fmt.Sprintf(format, "one", "1", "")
	case 0:
		line = fmt.Sprintf(format, "it", "no more", "s")
	case -1:
		line = fmt.Sprint("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
	default:
		line = fmt.Sprintf(format, "one", strconv.Itoa(n), "s")
	}
	return
}

// Verse of the song
func Verse(n int) (verse string, err error) {
	if n > 99 || n < 0 {
		err = fmt.Errorf("Verse must be between 0 and 99")
		return
	}
	verse = firstLine(n) + secondLine(n-1)
	return
}

// Verses in range s to e, in descending order
func Verses(s, e int) (verses string, err error) {
	if s < e {
		err = fmt.Errorf("Start must be greater than end")
		return
	}

	v := make([]string, 0)
	var verse string
	for x := s; x >= e; x-- {
		verse, err = Verse(x)
		if err != nil {
			return
		}

		v = append(v, verse)
	}
	verses = strings.Join(v, "\n") + "\n"
	return
}

// Song in its entirety
func Song() (song string) {
	song, _ = Verses(99, 0)
	return
}
