// Package clock implements a basic dateless clock with minute precision
package clock

import "fmt"

const testVersion = 4

type Clock struct {
	Hour   int
	Minute int
}

// normalise the total minutes to be a positive int <= 1440.
//
// 1440 is the number of minutes in a day. If there are <0 or
// >1440 minutes, calculate how many that comes to.
func normalise(hour, minute int) (total int) {
	total = hour*60 + minute // total minutes requested, positive or negative
	total = total % 1440     // fit within +/- one day
	total = total + 1440     // add a whole day to make ensure a positive value
	total = total % 1440     // fit the new value within one day
	return
}

// New creates a new clock instance normalised to a 24 hour clock.
//
// Negative hour and minute values are deducted accordingly.
func New(hour, minute int) Clock {
	total := normalise(hour, minute)
	hour = total / 60
	minute = total % 60
	return Clock{
		Hour:   hour,
		Minute: minute,
	}
}

// String returns the time as a string in the format HH:MM.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add more minutes and normalise to a 24 hour clock.
func (c Clock) Add(minute int) Clock {
	return New(c.Hour, c.Minute+minute)
}
