// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Calculates the moment when someone has lived for 10^9 seconds.
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	gigaSecond := time.Second * time.Duration(math.Pow10(9))
	return t.Add(gigaSecond)
}
