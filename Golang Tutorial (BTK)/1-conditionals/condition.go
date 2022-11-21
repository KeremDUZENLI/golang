package conditionals

import "fmt"

func Compare(a int, b int, c int) int {
	if a > b && a > c {
		fmt.Println(a)
		return a
	} else if b > a && b > c {
		fmt.Println(b)
		return b
	} else {
		fmt.Println(c)
		return c
	}
}
