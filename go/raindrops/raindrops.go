// Package raindrops is a number to string package like fizzbuzz.
package raindrops

import (
	"fmt"
)

// Convert is a string that converts its int input into
// a string of some combination of Pling, Plang and Plong.
func Convert(v int) (s string) {
	if v%3 == 0 {
		s += "Pling"
	}
	if v%5 == 0 {
		s += "Plang"
	}
	if v%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = fmt.Sprint(v)
	}
	return
}
