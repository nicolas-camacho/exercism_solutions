package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardsValues = map[string]int{
		"ace": 11,
		"ten": 10, "jack": 10, "queen": 10, "king": 10,
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
		"other": 0,
	}

	return cardsValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	var playerSum = ParseCard(card1) + ParseCard(card2)
	var dealerCardValue = ParseCard(dealerCard)

	if playerSum == 21 && dealerCardValue < 10 {
		return "W"
	}

	if (playerSum >= 17 && playerSum <= 20) || (playerSum == 21 && dealerCardValue >= 10) || (playerSum >= 12 && playerSum <= 16 && dealerCardValue < 7) {
		return "S"
	}

	return "H"
}
