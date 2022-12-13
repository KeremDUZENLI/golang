package Pointers

import (
	"fmt"
)

func Pointer() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)

	a = 27
	fmt.Println(a, *b)

	*b = 10
	fmt.Println(a, *b)
}
