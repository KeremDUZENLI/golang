package GoRoutines

import (
	"fmt"
	"time"
)

func GoRoutines() {
	var msg = "Hello"

	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}
