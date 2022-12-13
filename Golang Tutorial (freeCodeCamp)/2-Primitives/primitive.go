package Primitives

import (
	"fmt"
)

func Boolean() {
	var n bool

	fmt.Printf("%v, %T\n", n, n)
}

func Integer() {
	var a int = 10
	var b int8 = 3

	fmt.Println(a + int(b))
}

func Operator() {
	a := 10
	b := 3

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
}

func String() {
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
