// Package blackjack simulates the first turn of a Blackjack game
package blackjack

var cards = map[string]int{
	"other": 0,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"ace":   11,
}

const blackjack int = 21

const (
	stand = "S"
	hit   = "H"
	split = "P"
	win   = "W"
)

func isBlackjack(playerScore int) bool {
	return playerScore == blackjack
}

func largeHandTurn(playerScore int, dealerScore int) string {
	var turn string
	switch {
	case !isBlackjack(playerScore):
		turn = split
	case dealerScore != 10 && dealerScore != 11:
		turn = win
	default:
		turn = stand
	}
	return turn
}

func smallHandTurn(playerScore int, dealerScore int) string {
	var turn string
	switch {
	case (playerScore >= 12 && playerScore <= 16) && dealerScore >= 7:
		turn = hit
	case playerScore <= 11:
		turn = hit
	default:
		turn = stand
	}
	return turn
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if playerScore < blackjack {
		return smallHandTurn(playerScore, dealerScore)
	}
	return largeHandTurn(playerScore, dealerScore)
}
