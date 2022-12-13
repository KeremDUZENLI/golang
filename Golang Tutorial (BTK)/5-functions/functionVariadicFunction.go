package functions

import "fmt"

func VariadicFunction(numbers ...int) int {
	total := 0

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	fmt.Println(total)
	return total
}
