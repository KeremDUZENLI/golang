package variable

import (
	"fmt"
	"math/rand"
)

type GenericT[T Interfac] struct {
	Cards []T
}

type Interfac interface {
	fmt.Stringer

	NameT() string
}

func (t *Generic[T]) AddcardT(card T) {
	t.Cards = append(t.Cards, card)
}

func (t *GenericT[T]) RandomcardT() T {
	r := rand.Intn(len(t.Cards))
	return t.Cards[r]
}

func (t *GenericT[T]) NameT(input string) string {
	return input
}
