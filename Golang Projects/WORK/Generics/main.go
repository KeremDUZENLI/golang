package main

import (
	"fmt"
	"generics/function"
	"generics/variable"
)

func main() {

	// Playingcard -----------------------------------------------------------

	fmt.Println("\nPlayingcardFuncNew: ", function.PlayingcardFuncNew("SUITN", "RANKN"))
	fmt.Println("\nPlayingcardFuncPoint: ", function.PlayingcardFuncPoint("SUITP", "RANKP"))

	playingcardCustom := variable.Playingcard{Suit: "SUITC", Rank: "RANKC"}
	fmt.Println("\nString: ", playingcardCustom.String())
	fmt.Println("\nName: ", playingcardCustom.Name())

	// Tradingcard -------------------------------------------------------------

	fmt.Println("\nTradingcardFuncNew: ", function.TradingcardFuncNew("xxx"))
	fmt.Println("\nTradingcardFuncPoint: ", function.TradingcardFuncPoint("xxx"))

	tradingcardCustom := variable.Tradingcard{Collectablename: "CUSTOM"}
	fmt.Println("\nString: ", tradingcardCustom.String())
	fmt.Println("\nName: ", tradingcardCustom.Name())

	// Deck -----------------------------------------------------------------

	fmt.Println("\nDeckFuncNew: ", function.DeckFuncNew())
	fmt.Println("\nDeckFuncPoint: ", function.DeckFuncPoint())

	randomcard := function.DeckFuncPoint().Randomcard()
	fmt.Println("\nRandomcard: ", randomcard)
	fmt.Println("\nRandomcard String: ", randomcard.(*variable.Playingcard).String())
	fmt.Println("\nRandomcard Name: ", randomcard.(*variable.Playingcard).Name())
	fmt.Println("\nRandomcard Rank: ", randomcard.(*variable.Playingcard).Rank)
	fmt.Println("\nRandomcard Suit: ", randomcard.(*variable.Playingcard).Suit)

	// Generic ------------------------------------------------------------

	fmt.Println("\nGenericFuncNew: ", function.GenericFuncNew())
	fmt.Println("\nGenericFuncPoint: ", function.GenericFuncPoint())

	RandomcardGeneric := function.GenericFuncPoint().Randomcard()
	fmt.Println("\nRandomcardGeneric: ", RandomcardGeneric)
	fmt.Println("\nRandomcardGeneric String: ", RandomcardGeneric.String())
	fmt.Println("\nRandomcardGeneric Name: ", RandomcardGeneric.Name())
	fmt.Println("\nRandomcardGeneric Rank: ", RandomcardGeneric.Rank)
	fmt.Println("\nRandomcardGeneric Suit: ", RandomcardGeneric.Suit)

	fmt.Println("\nGenericFuncPointTradingCard: ", function.GenericFuncPointTradingcard())

	RandomcardTradingGeneric := function.GenericFuncPointTradingcard().Randomcard()
	fmt.Println("\nRandomcardTradingGeneric: ", RandomcardTradingGeneric)
	fmt.Println("\nRandomcardTradingGeneric String: ", RandomcardTradingGeneric.String())
	fmt.Println("\nRandomcardTradingGeneric Name: ", RandomcardTradingGeneric.Name())
	fmt.Println("\nRandomcardTradingGeneric Collectablename: ", RandomcardTradingGeneric.Collectablename)

	// GenericT ----------------------------------------------------------

	fmt.Println("\nGenericFuncPointT: ", function.GenericFuncPointT())
	fmt.Println("\nGenericFuncPointT Cards: ", function.GenericFuncPointT().CardsT)

	randomcardGenericT := function.GenericFuncPointT().RandomcardT()
	fmt.Println("\nRandomcardGenericT: ", randomcardGenericT)
	fmt.Println("\nRandomcardGenericT String: ", randomcardGenericT.String())
	fmt.Println("\nRandomcardGenericT Name: ", randomcardGenericT.Name())
	fmt.Println("\nRandomcardGenericT Rank: ", function.GenericFuncPointT().RandomcardT().Rank)
	fmt.Println("\nRandomcardGenericT Suit: ", function.GenericFuncPointT().RandomcardT().Suit)

	// -------------------------------------------------------------------

	function.PrintString[*variable.Playingcard](RandomcardGeneric)
	function.PrintString[*variable.Tradingcard](RandomcardTradingGeneric)

	function.PrintName[*variable.Playingcard](RandomcardGeneric)
	function.PrintName[*variable.Tradingcard](RandomcardTradingGeneric)

	// -------------------------------------------------------------------

	x := &variable.Playingcard{
		Suit: "NEW",
		Rank: "1000",
	}

	fmt.Println("\nAdd Card New: ", function.GenericFuncPoint().Addcard(x))

	fmt.Println("\nAdd Card New: ", function.GenericFuncPointT().AddcardT(x))
}
