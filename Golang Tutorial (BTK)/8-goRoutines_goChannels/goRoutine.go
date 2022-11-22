package goRoutines_goChannels

import (
	"fmt"
	"time"
)

func EvenNumbers() {
	for i := 0; i < 10; i += 2 {
		fmt.Print(" ", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func OddNumbers() {
	for i := 1; i < 10; i += 2 {
		fmt.Print(" ", i)
		time.Sleep(1 * time.Millisecond)
	}
}
