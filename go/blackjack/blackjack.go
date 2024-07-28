package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	age:=0

	switch card{
	case "ace":
		age=11
	case "two":
		age=2
	case "three":
		age=3
	case "four":
		age=4
	case "five":
		age=5
	case "six":
		age=6
	case "seven":
		age=7
	case "eight":
		age=8
	case "nine":
		age=9
	case "ten":
		age=10
	case "jack":
		age=10
	case "queen":
		age=10
	case "king":
		age=10
	default:
		age=0
	}

	return age
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	sum:=card1Value+card2Value
	dealerCardValue := ParseCard(dealerCard)

	move := ""

	switch{
	case card1Value==11 && card2Value==11:
		move="P"
	case sum==21 && dealerCardValue!=11 && dealerCardValue!=10:
		move="W"
	case sum==21 && (dealerCardValue==11 || dealerCardValue==10):
		move="S"
	case sum>=17 && sum<=20:
		move="S"
	case sum>=12 && sum<=16 && dealerCardValue<7:
		move="S"
	case sum>=12 && sum<=16 && dealerCardValue>=7:
		move="H"
	case sum<=11:
		move="H"
	}

	return move
}
