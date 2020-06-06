// Package scrabble provides computations for the Scrabble
package scrabble

import "strings"

// Score returns score for given word by rules:
// A, E, I, O, U, L, N, R, S, T - 1
// D, G                         - 2
// B, C, M, P                   - 3
// F, H, V, W, Y                - 4
// K                            - 5
// J, X                         - 8
// Q, Z                         - 10
func Score(word string) int {
	letters := []rune(strings.ToLower(word))

	score := 0
	for _, l := range letters {
		score += scoreLetter(l)
	}

	return score
}

func scoreLetter(l rune) int {
	switch l {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}
