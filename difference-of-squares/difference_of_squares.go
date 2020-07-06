// Package diffsquares solves math square problems.
// Under the hood, implementations use fast Gauss formulas.
package diffsquares

// SquareOfSum calculates square of sum for numbers until n.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates sum of squares for numbers until n.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates subtraction of SquareOfSum and SumOfSquares results for numbers until n.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
