package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var vName string = "Kerem"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4
	v4 = 5.4

	pl(vName, "\n", v1, v2, v3, v4)
}
