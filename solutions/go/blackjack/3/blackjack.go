package blackjack

const (
	hit   = "H"
	split = "P"
	stand = "S"
	win   = "W"
)

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

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardsValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerSum := ParseCard(card1) + ParseCard(card2)

	if playerSum > 21 {
		return split
	}

	dealerCardValue := ParseCard(dealerCard)

	if playerSum < 21 {
		switch {
		case playerSum > 16:
			return stand
		case playerSum < 12:
			return hit
		default:
			if dealerCardValue > 6 {
				return hit
			}
			return stand
		}
	}

	switch dealerCardValue {
	case 10, 11:
		return stand
	default:
		return win
	}
}
