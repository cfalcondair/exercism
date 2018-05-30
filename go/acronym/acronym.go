// Package acronym creates an acronym from a string.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate creates an acronym from a string.
func Abbreviate(s string) (acronym string) {
	reg := regexp.MustCompile("([A-Z])[^\\s]*(?:\\s|$)")
	s = strings.Replace(s, "-", " ", 100)
	s = strings.ToUpper(s)
	wordCount := len(strings.Split(s, " "))
	result := reg.FindAllStringSubmatch(s, wordCount)
	for _, regexMatch := range result {
		acronym += regexMatch[1]
	}
	return
}
