// Package encode provides implementation for Run-Length encoding.
// This implementation assumes that input to encode contains only letters A-Z and space.
package encode

import (
	"bytes"
	"strconv"
	"unicode"
)

// RunLengthEncode encodes input by RLE algorithm.
// Example, "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"  ->  "12WB12W3B24WB".
func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var output bytes.Buffer

	var counter int
	var prev rune
	for i, cur := range input {
		switch {
		case i == 0:
			counter = 1
			prev = cur
		case cur != prev:
			encode(&output, counter, prev)

			counter = 1
			prev = cur
		default:
			counter++
		}
	}
	encode(&output, counter, prev)

	return output.String()
}

func encode(output *bytes.Buffer, counter int, symbol rune) {
	if counter > 1 {
		output.WriteString(strconv.Itoa(counter))
	}
	output.WriteRune(symbol)
}

// RunLengthDecode decodes string by RLE algorithm.
// Example, "12WB12W3B24WB"  ->  "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB".
func RunLengthDecode(encoded string) string {
	if len(encoded) == 0 {
		return ""
	}

	var output bytes.Buffer

	var point int
	for i, sym := range encoded {
		if unicode.IsDigit(sym) {
			continue
		}

		counter := 1
		if point != i {
			num, err := strconv.Atoi(encoded[point:i])
			if err != nil {
				return ""
			}
			counter = num
		}

		decode(&output, counter, sym)

		point = i + 1
	}

	return output.String()
}

func decode(output *bytes.Buffer, counter int, symbol rune) {
	for i := 0; i < counter; i++ {
		output.WriteRune(symbol)
	}
}
