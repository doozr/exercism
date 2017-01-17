// Package gigasecond calculates when somebody is a gigasecond old.
package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds a gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
