package grains

import (
	"errors"
)

func Square(n int) (uint64, error) {
	if n > 0 && n <= 64 {
		return 1 << (n - 1), nil
	}

	return 0, errors.New("the number has to be between 1 and 64")
}

func Total() uint64 {
	return 1<<64 - 1
}
