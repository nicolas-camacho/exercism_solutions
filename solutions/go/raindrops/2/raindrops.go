package raindrops

import "strconv"

// Convert turn a number into a string depending if is factor of 3, 5 or 7
func Convert(number int) (result string) {

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if len(result) == 0 {
		return strconv.Itoa(number)
	}

	return result
}
