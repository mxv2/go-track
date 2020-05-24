// Package collatzconjecture provides Collatz Conjecture solution
package collatzconjecture

import "errors"

// CollatzConjecture returns count of steps for `n` to reach 1.
// If n is not positive when appropriate error returns.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("not positive number")
	}

	var step int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		step++
	}

	return step, nil
}
