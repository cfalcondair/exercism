// Package hamming implements measuring the hamming distance.
package hamming

import (
	"fmt"
)

type NotSameLength struct {
	a, b string
}

func (e NotSameLength) Error() string {
	return fmt.Sprintf("String a = %v not the same length as string b %v", e.a, e.b)
}

// Distance measuresthe hamming distance between two strings
func Distance(a, b string) (dist int, e error) {
	if len(a) != len(b) {
		dist = -1
		e = NotSameLength{a, b}
	} else {
		for i, _ := range a {
			if a[i] != b[i] {
				dist++
			}
		}
	}
	return
}
