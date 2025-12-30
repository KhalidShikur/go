package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
        	return 11
        case "two":
        	return 2
        case "three":
        	return 3
        case "four":
        	return 4
        case "five":
        	return 5
        case "six":
        	return 6
        case "seven":
        	return 7
        case "eight":
        	return 8
        case "nine":
        	return 9
        case "ten", "jack", "queen", "king":
        	return 10
        default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myCard1 := ParseCard(card1)
    myCard2 := ParseCard(card2)
    delerCard := ParseCard(dealerCard)
    var returns string

    switch {
        case myCard1 == 11 && myCard2 == 11:
        	returns = "P"
        	break
        case myCard1+myCard2 == 21 && (delerCard == 11 || delerCard == 10):
        	returns = "S"
        	break
        case myCard1+myCard2 >= 17 && myCard1+myCard2 <= 20:
        	returns = "S"
        	break
        case myCard1+myCard2 >= 12 && myCard1+myCard2 <= 16 && delerCard < 7:
        	returns = "S"
        	break
        case myCard1+myCard2 >= 12 && myCard1+myCard2 <= 16 && delerCard >= 7:
        	returns = "H"
        	break
        case myCard1+myCard2 <= 11:
        	returns = "H"
        	break
        default:
        	returns = "W"
    }

    return returns
}
