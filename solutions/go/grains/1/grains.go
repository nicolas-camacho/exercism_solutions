package grains

import (
	"errors"
	"math"
)

func Square(n int) (uint64, error) {
	if n > 0 && n <= 64 {
		return uint64(math.Pow(2, float64(n-1))), nil
	}

	return 0, errors.New("the number has to be between 1 and 64")
}

func Total() uint64 {
	return math.MaxUint64
}
