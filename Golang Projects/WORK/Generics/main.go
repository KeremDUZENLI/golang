package main

import (
	"generics/function"
	"generics/variable"
)

func main() {

	// Playingcard -----------------------------------------------------------

	// playingcardNew := function.PlayingcardFuncNew("SUITN", "RANKN")
	// fmt.Println(playingcardNew)

	// playingcardPoint := function.PlayingcardFuncPoint("SUITP", "RANKP")
	// fmt.Println(playingcardPoint)

	// playingcardCustom := variable.Playingcard{
	// 	Suit: "SUITC",
	// 	Rank: "RANKC",
	// }
	// fmt.Println(playingcardCustom.String())

	// Tradingcard -------------------------------------------------------------

	// TradingcardNew := function.TradingcardFuncNew("xxx")
	// fmt.Printf("\n--- drawing playing card ---\n %v\n", TradingcardNew)

	// TradingcardPoint := function.TradingcardFuncPoint("zzz")
	// fmt.Printf("\n--- drawing playing card ---\n %v\n", TradingcardPoint)

	// TradingcardCustom := variable.Tradingcard{
	// 	Collectablename: "CUSTOM",
	// }
	// fmt.Printf("\n--- drawing playing card ---\n %v\n", TradingcardCustom.String())

	// Deck -----------------------------------------------------------------

	// deckNew := function.DeckFuncNew()
	// fmt.Printf("\n--- drawing playing card ---\n %v\n", deckNew)

	// deckPoint := function.DeckFuncPoint()
	// fmt.Printf("\n--- drawing playing card ---\n %v\n", deckPoint)

	// deckPointCard := deckPoint.Randomcard()
	// fmt.Printf("\ndrew card: %s\n", deckPointCard)

	// deckPointCardSpec, ok := deckPointCard.(*variable.Playingcard)
	// if !ok {
	// 	fmt.Printf("card received wasn't a playing card!\n")
	// 	os.Exit(1)
	// }
	// fmt.Printf("card suit: %s\n", deckPointCardSpec.Suit)
	// fmt.Printf("card rank: %s\n", deckPointCardSpec.Rank)

	// var listeOut variable.Deck
	// function.DeckFuncCustom(&listeOut.Cards)
	// fmt.Printf("\n%v\n", listeOut.Cards)

	// Generic ------------------------------------------------------------

	// GenericNewDeck := function.GenericFuncNew()
	// fmt.Printf("\n--- drawing playing card ---\n %v\n", GenericNewDeck)

	// GenericNewDeckCard := GenericNewDeck.Randomcard()
	// fmt.Printf("\ndrewcard: %s\n", GenericNewDeckCard)

	// GenericPointDeck := function.GenericFuncPoint()
	// fmt.Printf("\n--- drawing playing card ---\n %v\n", GenericPointDeck)

	// GenericPointDeckCard := GenericPointDeck.Randomcard()
	// fmt.Printf("\ndrewcard: %s\n", GenericPointDeckCard)
	// fmt.Printf("card suit: %s\n", GenericPointDeckCard.Suit)
	// fmt.Printf("card rank: %s\n", GenericPointDeckCard.Rank)

	// GenericT ----------------------------------------------------------

	// GenericPointTradingcard := function.GenericFuncPointT()
	// fmt.Printf("\n--- drawing trading card ---\n %v\n", GenericPointTradingcard)

	// GenericPointTradingcardCard := GenericPointTradingcard.Randomcard()
	// fmt.Printf("\ndrew card: %s", GenericPointTradingcardCard)
	// fmt.Printf("\ncard collectable name:: %s\n", GenericPointTradingcardCard.Collectablename)

	GenericPointDeckT := function.GenericFuncPoint()
	GenericPointDeckCardT := GenericPointDeckT.Randomcard()
	function.PrintCard[*variable.Playingcard](GenericPointDeckCardT)
}
