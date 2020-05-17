// Package raindrops provides number conversion tool.
package raindrops

import "strconv"

// Convert returns string by rules for number:
// number has 3 as a factor, add 'Pling' to the result.
// number has 5 as a factor, add 'Plang' to the result.
// number has 7 as a factor, add 'Plong' to the result.
// number does not have any of 3, 5, or 7 as a factor,
// the result should be the digits of the number.
func Convert(a int) string {
	var b string
	if a%3 == 0 {
		b += "Pling"
	}
	if a%5 == 0 {
		b += "Plang"
	}
	if a%7 == 0 {
		b += "Plong"
	}
	if len(b) == 0 {
		return strconv.Itoa(a)
	}
	return b
}
