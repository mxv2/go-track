// Package isogram checks isogram words
package isogram

import "unicode"

// IsIsogram returns true if word is an isogram,
// otherwise false.
func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}

	var letterBits int32
	for _, l := range word {
		if unicode.IsSpace(l) || l == '-' {
			continue
		}

		l = unicode.ToLower(l) - 'a'
		bit := int32(1 << l)
		if letterBits&bit > 0 {
			return false
		}

		letterBits |= bit
	}

	return true
}
