// Package gigasecond calculates the moment when
// someone has lived for 10^9 seconds.
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond takes a time t and adds a gigasecond.
func AddGigasecond(t time.Time) time.Time {
	gigaSecond := time.Second * time.Duration(math.Pow10(9))
	return t.Add(gigaSecond)
}
