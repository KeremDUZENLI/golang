package Challenge19_CalculatingTheDifferenceBetweenSquares

import (
	"math"
)

func DiffSquares(n, m int) int {
	x1 := float64(n)
	x2 := float64(m)

	return int(math.Pow(x1, 2)) - int(math.Pow(x2, 2))
}
