// Package acronym implements function for the abbreviation.
package acronym

import "unicode"

// Abbreviate returns acronym for long name in string.
func Abbreviate(s string) string {
	if len(s) == 0 {
		return ""
	}
	sr := []rune(s)

	acronym, catch := make([]rune, 0), true
	for _, v := range sr {
		if isCatcher(v) {
			catch = true
			continue
		}

		if !unicode.IsLetter(v) {
			catch = false
		}

		if catch {
			acronym = append(acronym, unicode.ToUpper(v))
			catch = false
		}
	}
	return string(acronym)
}

func isCatcher(a rune) bool {
	switch a {
	case ' ', '-', '_':
		return true
	default:
		return false
	}
}
