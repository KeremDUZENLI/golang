package main

import (
	"fmt"
	"math/rand"
	"os"
)

// 0------------------------------------------------------------------------

type PlayingCard struct {
	Suit string
	Rank string
}

type Deck struct {
	cards []any
}

// 1------------------------------------------------------------------------

func PlayingCardFuncN(x string, y string) PlayingCard {
	return PlayingCard{
		Suit: x,
		Rank: y,
	}
}

func PlayingCardFuncP(x string, y string) *PlayingCard {
	return &PlayingCard{
		Suit: x,
		Rank: y,
	}
}

func (PlayingCard *PlayingCard) String() string {
	return fmt.Sprintf("%s - %s", PlayingCard.Suit, PlayingCard.Rank)
}

// 2------------------------------------------------------------------------

func PlayingCardDeckFuncN(listeOut *[]any) {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, rank := range ranks {
			*listeOut = append(*listeOut, PlayingCardFuncN(suit, rank))
		}
	}
}

func PlayingCardDeckFuncP() Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	DeckNew := Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			DeckNew.cards = append(DeckNew.cards, suit, rank)
		}
	}

	return DeckNew
}

// 3------------------------------------------------------------------------

func PlayingCardDeckFuncPS() *Deck {
	suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	DeckNew := &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			// DeckNew.cards = append(DeckNew.cards, suit, rank)
			DeckNew.AddCard(PlayingCardFuncN(suit, rank))
		}
	}

	return DeckNew
}

func (d *Deck) AddCard(card any) {
	d.cards = append(d.cards, card)
}

func (d *Deck) RandomCard() any {
	r := rand.Intn(len(d.cards))
	return d.cards[r]
}

func main() {

	// 1--------------------------------------------------------------------

	// s1 := PlayingCardFuncN("SUITN", "RANKN")
	// fmt.Println(s1)

	// s2 := PlayingCardFuncP("SUITP", "RANKP")
	// fmt.Println(s2)

	// s3 := PlayingCard{
	// 	Suit: "SUIT",
	// 	Rank: "RANK",
	// }
	// fmt.Println(s3.String())

	// 2--------------------------------------------------------------------

	// var x Deck
	// PlayingCardDeckFuncN(&x.cards)
	// fmt.Println(x.cards)
	// fmt.Printf("\n\n")

	// fmt.Println(PlayingCardDeckFuncP())
	// fmt.Printf("\n\n")
	// fmt.Println(PlayingCardDeckFuncPS())
	// fmt.Printf("\n\n")

	// fmt.Println(x.RandomCard())

	// 3--------------------------------------------------------------------

	deck := PlayingCardDeckFuncPS()
	fmt.Printf("\n\n--- drawing playing card ---\n\n %v", deck)

	card := deck.RandomCard()
	fmt.Printf("\n\ndrew card: %s\n\n", card)

	playingCard, ok := card.(PlayingCard)
	if !ok {
		fmt.Printf("card received wasn't a playing card!\n")
		os.Exit(1)
	}
	fmt.Printf("card suit: %s\n", playingCard.Suit)
	fmt.Printf("card rank: %s\n", playingCard.Rank)
}
