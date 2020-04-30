// Package hamming implements Hamming distance calculation
package hamming

import "errors"

// Distance returns calculated Hamming distance between two words.
// An error is returned if words have not equal length.
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("words not equal")
	}

	distance := 0
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}

	return distance, nil
}
