package isogram

import "unicode"

//IsIsogram determine if a phare is an isogram by checkig if any letter is been repeated
func IsIsogram(phrase string) bool {
	charsMap := make(map[rune]bool)

	for _, character := range phrase {
		lowerChar := unicode.ToLower(character)
		if !unicode.IsLetter(lowerChar) {
			continue
		}
		if charsMap[lowerChar] {
			return false
		}
		charsMap[lowerChar] = true
	}

	return true

}
