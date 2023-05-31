package function

import (
	"fmt"
	"generics/variable"
)

// func GenericFuncNewT() variable.GenericT[variable.Playingcard] {
// 	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
// 	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
// 	DeckNew := variable.GenericT[variable.Playingcard]{}

// 	for _, suit := range suits {
// 		for _, rank := range ranks {
// 			DeckNew.AddcardT(PlayingcardFuncNew(suit, rank))
// 		}
// 	}

// 	return DeckNew
// }

func GenericFuncPointT() *variable.GenericT[*variable.Playingcard] {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	DeckNew := &variable.GenericT[*variable.Playingcard]{}

	for _, suit := range suits {
		for _, rank := range ranks {
			DeckNew.AddcardT(PlayingcardFuncPoint(suit, rank))
		}
	}

	return DeckNew
}

func PrintString[C variable.Interfac](card C) {
	fmt.Println("\nPrintString: ", card.String())
}

func PrintName[C variable.Interfac](card C) {
	fmt.Println("\nPrintName: ", card.Name())
}
