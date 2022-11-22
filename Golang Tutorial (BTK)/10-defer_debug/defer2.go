package defer_debug

import "fmt"

func test(number int) {
	defer deferFunc()

	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func deferFunc() {
	fmt.Println("Deferred...")
}

func Defer2() {
	test(10)
}
