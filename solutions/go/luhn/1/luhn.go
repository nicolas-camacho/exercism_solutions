package luhn

import (
	"strings"
)

//Valid evaluate a number to check if it's valid according to the Luhn algorithm
func Valid(number string) bool {
	number = strings.Replace(number, " ", "", -1)
	if len(number) > 1 {
		sum := 0
		for i := len(number) - 1; i >= 0; i-- {
			digit := int(rune(number[i]) - '0')
			if digit < 0 || digit > 9 {
				return false
			}
			if i%2 == len(number)%2 {
				digit *= 2
			}
			if digit > 9 {
				digit -= 9
			}
			sum += digit
		}

		return sum%10 == 0
	}

	return false
}
