package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	for x := 1; x <= 5; x++ {
		pl(x)
	}

	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}
}
