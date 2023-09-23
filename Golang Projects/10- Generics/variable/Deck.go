package variable

import (
	"math/rand"
)

type Deck struct {
	Cards []any
}

func (d *Deck) Addcard(Card any) {
	d.Cards = append(d.Cards, Card)
}

func (d *Deck) Randomcard() any {
	r := rand.Intn(len(d.Cards))
	return d.Cards[r]
}
