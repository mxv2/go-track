// Package summultiples provides calculation multiplication sum.
package summultiples

// SumMultiples calculates the sum of numbers in the range from 1 to limit (excluding),
// that are divisible by at least one of the divisors
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
mainLoop:
	for i := 0; i < limit; i++ {
		for _, d := range divisors {
			if d != 0 && i%d == 0 {
				sum += i
				continue mainLoop
			}
		}
	}

	return sum
}
