package raindrops

import "strconv"

//Convert turn a number into a string depending if is factor of 3, 5 or 7
func Convert(number int) (result string) {
	sounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for factor, sound := range sounds {
		if number%factor == 0 {
			result += sound
		}
	}

	if len(result) == 0 {
		return strconv.Itoa(number)
	}

	return result
}
