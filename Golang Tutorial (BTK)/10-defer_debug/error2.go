package defer_debug

import "fmt"

func verify(i interface{}) {
	number, ok := i.(int)
	fmt.Println(number, ok)
}

func Error2() {
	var x interface{} = "5"
	verify(x)

	var y interface{} = 5
	verify(y)
}
