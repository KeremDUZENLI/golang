package GoRoutines

import (
	"fmt"
	"time"
)

func GoRoutine() {
	var msg = "Hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)

	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
