// Package diffsquares provides integers square sum utils
package diffsquares

// SquareOfSums returns square of sums
func SquareOfSums(n int) int {
	// @see https://www.wolframalpha.com/input/?i=1+%2B+2+%2B+3+%2B+...+%2B+n
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns sum of squares
func SumOfSquares(n int) int {
	// @see https://www.wolframalpha.com/input/?i=1+%2B+4+%2B+9%2B+...+%2B+n
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns difference between square of sums and sum of squares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
