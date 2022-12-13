package loops

import (
	"fmt"
)

func Loop2() {
	var number int
	var result string

	fmt.Println("Number: ")
	fmt.Scanln(&number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			result = "Not Prime"
			break
		} else {
			result = "Prime"
		}
	}

	fmt.Printf("%v is %v", number, result)
}
