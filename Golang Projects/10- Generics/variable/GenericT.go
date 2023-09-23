package variable

import (
	"math/rand"
)

type GenericT[T Interfac] struct {
	CardsT []T
}

type Interfac interface {
	String() string
	Name() string
}

func (t *GenericT[T]) AddcardT(card T) []T {
	t.CardsT = append(t.CardsT, card)
	return t.CardsT
}

func (t *GenericT[T]) RandomcardT() T {
	r := rand.Intn(len(t.CardsT))
	return t.CardsT[r]
}
