package pointers_structs

import "fmt"

func Pointer2(numbers []int) {
	numbers[0] = 100
	fmt.Printf("Pointer2(): %v\n", numbers)
}
