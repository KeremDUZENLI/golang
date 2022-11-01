package main

import (
	"fmt"
)

func Calculate(x int) (y int) {
	y = x + 2
	return y
}

func main() {
	fmt.Println("Go Testing Tutorial")
	fmt.Println(Calculate(2))
}
