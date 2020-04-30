// Package hamming implements Hamming distance calculation
package hamming

import "errors"

// Distance returns calculated Hamming distance between two words.
// An error is returned if words have not equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("words not equal")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
