package isogram

import "unicode"

//IsIsogram determine if a phare is an isogram by checkig if any letter is been repeated
func IsIsogram(phrase string) bool {
	charsMap := make(map[rune]bool)
	uniqueChars := ""

	for _, character := range phrase {
		if unicode.IsLetter(character) {
			lowerChar := unicode.ToLower(character)
			if _, exists := charsMap[lowerChar]; !exists {
				charsMap[lowerChar] = true
				uniqueChars += string(lowerChar)
			}
		} else {
			uniqueChars += string(character)
		}
	}

	return len(phrase) == len(uniqueChars)

}
