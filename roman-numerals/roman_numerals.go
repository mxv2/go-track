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

// ToRomanNumeral converts `n` to roman number as string.
// Returns error when number does not satisfy range (1 to 3000).
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", errors.New("expect number from 1 to 3000")
	}

	res := ""
	for n > 0 {
		switch {
		case n >= 1000:
			n -= 1000
			res += table[1000]
		case n-900 >= 0 && n-900 < 100:
			n -= 900
			res += table[100] + table[1000]
		case n >= 500:
			n -= 500
			res += table[500]
		case n-400 >= 0 && n-400 < 100:
			n -= 400
			res += table[100] + table[500]
		case n >= 100:
			n -= 100
			res += table[100]
		case n-90 >= 0 && n-90 < 10:
			n -= 90
			res += table[10] + table[100]
		case n >= 50:
			n -= 50
			res += table[50]
		case n-40 >= 0 && n-40 < 10:
			n -= 40
			res += table[10] + table[50]
		case n >= 10:
			n -= 10
			res += table[10]
		case n-9 >= 0 && n-9 < 1:
			n -= 9
			res += table[1] + table[10]
		case n >= 5:
			n -= 5
			res += table[5]
		case n-4 >= 0 && n-4 < 1:
			n -= 4
			res += table[1] + table[5]
		case n >= 1:
			n--
			res += table[1]
		}
	}

	return res, nil
}
