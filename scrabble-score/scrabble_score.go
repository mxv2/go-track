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
	letters := strings.ToLower(word)

	score := 0
	for _, l := range letters {
		switch l {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}

	return score
}
