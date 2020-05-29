// Package dna provides
package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

func newHistogram() map[byte]int {
	h := make(map[byte]int)
	h['A'], h['C'], h['G'], h['T'] = 0, 0, 0, 0
	return h
}

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []byte

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := newHistogram()
	for _, n := range d {
		if _, ok := h[n]; ok {
			h[n]++
		} else {
			return nil, fmt.Errorf("illegal nucleotide symbol: %s", string(n))
		}
	}

	return h, nil
}
