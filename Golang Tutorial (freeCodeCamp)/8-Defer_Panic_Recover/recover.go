package Defer_Panic_Recover

import (
	"fmt"
	"log"
)

func Recover() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err) // repanic
		}
	}()
	panic("something went wrong")
}
