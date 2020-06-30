// Package etl provides transform functions
package etl

import (
	"strings"
)

// Transform returns pivoted table representation.
// String elements cast to lower-case.
func Transform(table map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, letters := range table {
		for _, l := range letters {
			l = strings.ToLower(l)
			result[l] = score
		}
	}
	return result
}
