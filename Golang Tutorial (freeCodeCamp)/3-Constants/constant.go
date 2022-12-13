package Constants

import (
	"fmt"
)

func Constant() {
	const a int = 42
	var b int = 27

	fmt.Printf("%v, %T\n", a+b, a+b)
}

const (
	a = iota
	b
	c
)

func Iota() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}

const (
	_  = iota // ignore the first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func IotaSize() {
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
