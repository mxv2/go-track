// Package diffsquares solves math square problems.
package diffsquares

// SquareOfSum calculates square of sum for numbers until n.
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares calculates sum of squares for numbers until n.
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference calculates subtraction of SquareOfSum and SumOfSquares results for numbers until n.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
