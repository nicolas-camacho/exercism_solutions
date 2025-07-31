//Package grains package is an algorithm for solve the wheat and chessboard problem
package grains

import "errors"

//Square function return the number of grains based on a given square (integer)
func Square(n int) (uint64, error) {
	if n > 0 && n <= 64 {
		return 1 << (n - 1), nil
	}

	return 0, errors.New("the number has to be between 1 and 64")
}

//Total function return the number of grains for every square in a chessboard
func Total() uint64 {
	return 1<<64 - 1
}
