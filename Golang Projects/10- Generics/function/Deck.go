package function

import (
	"generics/variable"
)

func DeckFuncNew() variable.Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	DeckNew := variable.Deck{}

	for _, suit := range suits {
		for _, rank := range ranks {
			// DeckNew.Cards = append(DeckNew.Cards, suit, rank)
			DeckNew.Addcard(PlayingcardFuncNew(suit, rank))
		}
	}

	return DeckNew
}

func DeckFuncPoint() *variable.Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	DeckNew := &variable.Deck{}

	for _, suit := range suits {
		for _, rank := range ranks {
			// DeckNew.cards = append(DeckNew.cards, suit, rank)
			DeckNew.Addcard(PlayingcardFuncPoint(suit, rank))
		}
	}

	return DeckNew
}
