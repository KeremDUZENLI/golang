package t

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type myInterface interface {
	int | float64 | ~string | Weekday
}

func MaxWeek[T myInterface](input []T) T {
	var max T

	for _, w := range input {
		if w > max {
			max = w
		}
	}

	return max
}
