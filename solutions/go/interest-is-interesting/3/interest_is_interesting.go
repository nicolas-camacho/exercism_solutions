package interest

const (
	lessThanZero            float32 = 3.213
	zeroToThousand          float32 = 0.5
	thousandToFiveThousand  float32 = 1.621
	greaterThanFiveThousand float32 = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance >= 0 && balance < 1000 {
		return zeroToThousand
	} else if balance >= 1000 && balance < 5000 {
		return thousandToFiveThousand
	} else if balance >= 5000 {
		return greaterThanFiveThousand
	}

	return lessThanZero
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}

	return years
}
