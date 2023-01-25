package variable

import (
	"math/rand"
)

type Generic[C any] struct {
	Cards []C
}

func (c *Generic[C]) Addcard(card C) {
	c.Cards = append(c.Cards, card)
}

func (c *Generic[C]) Randomcard() C {
	r := rand.Intn(len(c.Cards))
	return c.Cards[r]
}

func (c *Generic[C]) Name(input string) string {
	return input
}
