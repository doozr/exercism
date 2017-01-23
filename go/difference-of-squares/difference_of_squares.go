// Package diffsquares does calculations with sums and squares
package diffsquares

// SquareOfSums calculates the square of the sum of 1 to limit
func SquareOfSums(limit int) int {
	squareOfSum := 1
	for x := 2; x <= limit; x++ {
		squareOfSum += x
	}
	squareOfSum *= squareOfSum
	return squareOfSum
}

// SumOfSquares calculates the sum of the squares of 1 to limit
func SumOfSquares(limit int) int {
	sumOfSquares := 1
	for x := 2; x <= limit; x++ {
		sumOfSquares += x * x
	}
	return sumOfSquares
}

// Difference calculates the difference between SquareOfSums(limit) and SumOfSquares(limit)
func Difference(limit int) int {
	return SquareOfSums(limit) - SumOfSquares(limit)
}
