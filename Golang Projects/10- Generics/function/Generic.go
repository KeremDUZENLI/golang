package function

import (
	"generics/variable"
)

func GenericFuncNew() variable.Generic[variable.Playingcard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	DeckNew := variable.Generic[variable.Playingcard]{}

	for _, suit := range suits {
		for _, rank := range ranks {
			DeckNew.Addcard(PlayingcardFuncNew(suit, rank))
		}
	}

	return DeckNew
}

func GenericFuncPoint() *variable.Generic[*variable.Playingcard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	DeckNew := &variable.Generic[*variable.Playingcard]{}

	for _, suit := range suits {
		for _, rank := range ranks {
			DeckNew.Addcard(PlayingcardFuncPoint(suit, rank))
		}
	}

	return DeckNew
}

func GenericFuncPointTradingcard() *variable.Generic[*variable.Tradingcard] {
	collectables := []string{"Sammy", "Droplets", "Spaces", "App Platform"}
	DeckNew := &variable.Generic[*variable.Tradingcard]{}

	for _, collectable := range collectables {
		DeckNew.Addcard(TradingcardFuncPoint(collectable))
	}

	return DeckNew
}
