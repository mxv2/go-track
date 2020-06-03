// Package romannumerals provides interchanging
// arabic to roman numbers functionality
package romannumerals

import "errors"

var table = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

var ranks = []int{1000, 500, 100, 50, 10, 5, 1}
var subRanks = []int{100, 100, 10, 10, 1, 1, 1}

// ToRomanNumeral converts `n` to roman number as string.
// Returns error when number does not satisfy range (1 to 3000).
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", errors.New("expect number from 1 to 3000")
	}

	res := ""
	for i := 0; n > 0; {
		rnk := ranks[i]
		if n >= rnk {
			n -= rnk
			res += table[rnk]
			continue
		}

		srnk := subRanks[i]
		specialCase := rnk - srnk
		if n-specialCase >= 0 && n-specialCase < srnk {
			n -= specialCase
			res += table[srnk] + table[rnk]
			continue
		}

		i++
	}

	return res, nil
}
