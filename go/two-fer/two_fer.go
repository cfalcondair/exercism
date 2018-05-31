// Package twofer just lets you know who your sharing with.
package twofer

// ShareWith returns a who you are sharing with
func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}
	return "One for " + s + ", one for me."
}
