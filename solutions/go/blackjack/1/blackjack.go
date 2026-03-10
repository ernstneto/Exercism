package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	//panic("Please implement the ParseCard function")
    switch card{
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
	//panic("Please implement the FirstTurn function")
    sum := ParseCard(card1) + ParseCard(card2)
	sumDealer := ParseCard(dealerCard)

	// Regra 1: Par de Ases
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// Avaliando as outras regras usando o switch idiomático
	switch {
	case sum == 21:
		// Se o dealer tem 10, figura ou ás (valor >= 10), tem que esperar (S)
		if sumDealer >= 10 {
			return "S"
		}
		// Caso contrário, vitória automática (W)
		return "W"
		
	case sum >= 17 && sum <= 20:
		return "S"
		
	case sum >= 12 && sum <= 16:
		// Se o dealer tem 7 ou mais, pedimos carta (H)
		if sumDealer >= 7 {
			return "H"
		}
		// Caso contrário, ficamos parados (S)
		return "S"
		
	default:
		// Isso cobre o "sum <= 11"
		return "H"
	}
}
