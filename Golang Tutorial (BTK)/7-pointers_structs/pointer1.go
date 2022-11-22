package pointers_structs

import "fmt"

func Pointer1(number *int) {
	*number = *number + 1
	fmt.Printf("Pointer(): %v\n", *number)

}
