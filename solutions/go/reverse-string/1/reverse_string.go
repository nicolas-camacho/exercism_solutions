package reverse

//Reverse return the reversed form of an original string
func Reverse(original string) (reversed string) {
	for _, letter := range original {
		reversed = string(letter) + reversed
	}

	return reversed
}
