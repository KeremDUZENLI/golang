package defer_debug

import "fmt"

func Defer11() {
	fmt.Println("1")
}

func Defer12() {
	defer Defer11()
	fmt.Println("2")
}

func Defer13() {
	defer Defer11()
	defer Defer12()
	fmt.Println("3")
}
