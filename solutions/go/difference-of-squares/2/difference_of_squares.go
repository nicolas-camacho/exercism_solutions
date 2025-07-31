package diffsquares

// SquareOfSum return the square of the sum for the amount of natural number defined
func SquareOfSum(number int) int {
	result := number * (number + 1) / 2
	return result * result
}

// SumOfSquares return the sum of squares for the amout of natural numbers defined
func SumOfSquares(number int) int {
	return number * (number + 1) * (2*number + 1) / 6
}

// Difference return the diference of the square of sum and sum of square for the
// amount of natural numbers defined
func Difference(limit int) (result int) {
	return SquareOfSum(limit) - SumOfSquares(limit)
}
