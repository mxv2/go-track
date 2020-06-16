// Package isogram checks isogram words
package isogram

import "unicode"

// IsIsogram returns true if word is an isogram,
// otherwise false.
func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}

	letters := make(map[rune]struct{})
	for _, l := range word {
		if unicode.IsSpace(l) || l == '-' {
			continue
		}

		l = unicode.ToLower(l)
		if _, ok := letters[l]; ok {
			return false
		}

		letters[l] = struct{}{}
	}

	return true
}
