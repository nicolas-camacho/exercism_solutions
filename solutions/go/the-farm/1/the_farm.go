package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	message string
}

func (err InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", err.cows, err.message)
}

func DivideFood(fooder FodderCalculator, cows int) (float64, error) {
	fooderAmount, err := fooder.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fooder.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (fooderAmount / float64(cows)) * factor, nil
}

func ValidateInputAndDivideFood(fooder FodderCalculator, cows int) (float64, error) {
	if cows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fooder, cows)
}

func ValidateNumberOfCows(cows int) error {
	err := InvalidCowsError{
		cows: cows,
	}
	if cows == 0 {
		err.message = "no cows don't need food"
		return &err
	} else if cows < 0 {
		err.message = "there are no negative cows"
		return &err
	}
	return nil
}
