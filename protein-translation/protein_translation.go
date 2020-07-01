// Package protein provides an interface for translating
// Codon and RNA sequence to Protein sequence.
package protein

import "errors"

var stop = "STOP"

var proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

// ErrInvalidBase is the error for undefined codon.
var ErrInvalidBase = errors.New("not a codon sequence")

// ErrStop declares end of the sequence.
var ErrStop = errors.New("stop a sequence")

// FromCodon translates codon (3-nucleotides) to proteins.
// ErrStop returns for stop-sequence proteins.
//
// If codon is not defined, when ErrInvalidBase returns.
func FromCodon(codon string) (string, error) {
	protein, ok := proteins[codon]

	if !ok {
		return "", ErrInvalidBase
	}
	if protein == stop {
		return "", ErrStop
	}

	return protein, nil
}

// FromRNA collects proteins from rna sequence.
//
// If error occurs via translating codon to protein,
// when returns proteins with appropriate error.
func FromRNA(rna string) ([]string, error) {
	proteins := make([]string, 0)

	acc := make([]rune, 0, 3)
	for _, nucleotide := range rna {
		acc = append(acc, nucleotide)
		if len(acc) != 3 {
			continue
		}

		protein, err := FromCodon(string(acc))
		if err != nil {
			if errors.Is(err, ErrStop) {
				break
			}
			return proteins, err
		}
		proteins = append(proteins, protein)

		acc = acc[:0]
	}

	return proteins, nil
}
