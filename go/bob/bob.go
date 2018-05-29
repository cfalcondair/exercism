// Package bob responds to statements in rule based way
// based on the content of the statement said to bob.
package bob

import (
	"regexp"
)

// Hey takes a remark and returns what Bob
// says to the remark.
func Hey(remark string) string {
	question := regexp.MustCompile("\\?\\s*$")
	yell := regexp.MustCompile("^[!\\^*\\($#@%0-9A-Z ,]+!$")
	yellNoExclamation := regexp.MustCompile("^[A-Z ]+$")
	gibberish := regexp.MustCompile("^[A-Z]+$")
	silence := regexp.MustCompile("^\\s*$")
	yellQuestion := regexp.MustCompile("^[A-Z ,]+\\?$")

	switch {
	case silence.MatchString(remark):
		return "Fine. Be that way!"
	case yellQuestion.MatchString(remark):
		return "Calm down, I know what I'm doing!"
	case yell.MatchString(remark) || gibberish.MatchString(remark) || yellNoExclamation.MatchString(remark):
		return "Whoa, chill out!"
	case question.MatchString(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}
