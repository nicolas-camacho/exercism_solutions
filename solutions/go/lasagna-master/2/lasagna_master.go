package lasagna

// PreparationTime returns the total preparation time in minutes for a lasagna
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	}

	return len(layers) * timePerLayer
}

// Quantities returns the quantities of noodles and sauce needed for a lasagna
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}

		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient adds the secret ingredient from the friends list to my list
func AddSecretIngredient(friendsList, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

// ScaleRecipe scales the quantities of ingredients for the given portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaled []float64
	for _, quantity := range quantities {
		scaled = append(scaled, (quantity/2)*float64(portions))
	}
	return scaled
}
