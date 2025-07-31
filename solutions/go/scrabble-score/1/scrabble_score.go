package scrabble

import "strings"

func getLettersScore() map[int][]string {
	return map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}
}

//Score function return the score of string based on its letters
func Score(scrable string) (score int) {

	scrable = strings.ToUpper(scrable)

	for value, list := range getLettersScore() {
		for _, letter := range list {
			if strings.Contains(scrable, letter) {
				score += value * strings.Count(scrable, letter)
			}
		}
	}
	return score
}
