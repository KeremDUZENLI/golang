package t

type ordered interface {
	int | float64 | ~string
}

func Max[T ordered](input []T) T {
	var max T

	for _, v := range input {
		if v > max {
			max = v
		}
	}

	return max
}
