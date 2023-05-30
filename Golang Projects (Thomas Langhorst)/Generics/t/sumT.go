package t

import (
	"golang.org/x/exp/constraints"
)

func SumT[T constraints.Ordered](v1, v2 T) T {
	return v1 + v2
}
