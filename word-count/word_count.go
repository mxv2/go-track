// Package wordcount provides word counting function in text.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is a storage, which counts strings in a case-insensitive way.
type Frequency map[string]int

// WordCount counts frequency of occurrences words in input.
// Word is a number (1234), simple string (ABCD) or with apostrophe.
// Quoted words are ignored.
func WordCount(input string) Frequency {
	freq := make(map[string]int)

	wordStartIndex := -1
	for peakIndex, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'' {
			if wordStartIndex < 0 {
				wordStartIndex = peakIndex
			}
			continue
		}

		if wordStartIndex < 0 {
			continue
		}

		countWord(freq, input[wordStartIndex:peakIndex])
		wordStartIndex = -1
	}
	if wordStartIndex >= 0 && wordStartIndex < len(input) {
		countWord(freq, input[wordStartIndex:])
	}

	return freq
}

func countWord(freq Frequency, word string) {
	word = strings.ToLower(strings.Trim(word, "'"))

	if _, ok := freq[word]; !ok {
		freq[word] = 0
	}
	freq[word]++
}
