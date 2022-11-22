package goRoutines_goChannels

import (
	"fmt"
	"time"
)

func EvenNumbers_(evenCh chan int) {
	var total int
	for i := 0; i < 10; i += 2 {
		total += i
		fmt.Println("Working on EvenNumbers")
		time.Sleep(1 * time.Second)
	}

	evenCh <- total
}

func OddNumbers_(oddCh chan int) {
	var total int
	for i := 1; i < 10; i += 2 {
		total += i
		fmt.Println("Working on OddNumbers")
		time.Sleep(1 * time.Second)
	}

	oddCh <- total
}
