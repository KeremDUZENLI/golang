package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	aNums := []int{1, 2, 3}

	for _, num := range aNums {
		pl(num)
	}
}
