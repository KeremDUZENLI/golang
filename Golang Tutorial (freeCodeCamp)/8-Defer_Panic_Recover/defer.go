package Defer_Panic_Recover

import (
	"fmt"
)

func Defer() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}
