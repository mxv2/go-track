// Package strand provides implementation for DNA to RNA transforming
package strand

var transcripts = map[byte]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns RNA transformed from passed DNA `dna`
// If dna contains illegal nucleotide when result will be empty
func ToRNA(dna string) string {
	dnab := []byte(dna)
	rnab := make([]byte, len(dnab))
	for i, dn := range dnab {
		if rn, ok := transcripts[dn]; ok {
			rnab[i] = rn
		} else {
			return ""
		}
	}

	return string(rnab)
}
