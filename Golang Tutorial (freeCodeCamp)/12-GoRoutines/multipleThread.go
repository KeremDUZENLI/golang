package GoRoutines

import (
	"fmt"
	"runtime"
)

func MultipleThread() {
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
