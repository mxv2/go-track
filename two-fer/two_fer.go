// Package twofer implements simple counting.
package twofer

import (
	"strings"
)

// ShareWith returns two-fer counting message.
func ShareWith(name string) string {
	if isBlank(name) {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

func isBlank(s string) bool {
	if s = strings.Trim(s, " "); len(s) == 0 {
		return true
	}
	return false
}
