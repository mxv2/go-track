// Package twofer implements simple counting.
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith returns two-fer counting message.
func ShareWith(name string) string {
	if isBlank(name) {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}

func isBlank(s string) bool {
	if s = strings.Trim(s, " "); len(s) == 0 {
		return true
	}
	return false
}
