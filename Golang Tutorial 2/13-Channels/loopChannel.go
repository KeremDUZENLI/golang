package Channels

import (
	"fmt"
	"sync"
)

var wg3 = sync.WaitGroup{}

func LoopChannel() {
	ch := make(chan int, 50)
	wg3.Add(2)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg3.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 40
		ch <- 50
		close(ch)
		wg3.Done()
	}(ch)

	wg3.Wait()
}
