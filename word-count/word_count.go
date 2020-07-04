package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

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
