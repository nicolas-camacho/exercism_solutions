package diffsquares

// SquareOfSum return the square of the sum for the amount of natural number defined
func SquareOfSum(limit int) (result int) {
	for i := 1; i <= limit; i++ {
		result += i
	}

	return result * result
}

// SumOfSquares return the sum of squares for the amout of natural numbers defined
func SumOfSquares(limit int) (result int) {
	for i := 1; i <= limit; i++ {
		result += i * i
	}

	return result
}

// Difference return the diference of the square of sum and sum of square for the
// amount of natural numbers defined
func Difference(limit int) (result int) {
	return SquareOfSum(limit) - SumOfSquares(limit)
}
