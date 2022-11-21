package arrays

import "fmt"

func Array2() {
	var cities [5]int
	for i := 0; i < 5; i++ {
		cities[i] = i
	}
	fmt.Println(cities)

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
}
