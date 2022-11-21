package arrays

import "fmt"

func Array3() {
	var numbers [2][3]int

	numbers[0][0] = 1
	numbers[0][1] = 2
	numbers[0][2] = 3
	numbers[1][0] = 4
	numbers[1][1] = 5
	numbers[1][2] = 6

	for i := 0; i < 2; i++ {
		for l := 0; l < 3; l++ {
			fmt.Print(numbers[i][l])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
