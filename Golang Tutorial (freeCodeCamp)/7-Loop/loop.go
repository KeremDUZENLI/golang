package Loop

import (
	"fmt"
)

func Loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func Collection() {
	s := [...]string{"A", "B", "C"}

	for k, v := range s {
		fmt.Println(k, v)
	}
}
