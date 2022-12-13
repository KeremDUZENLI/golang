package Channels

import (
	"fmt"
	"sync"
)

var wg2 = sync.WaitGroup{}

func BufferedChannel() {
	ch := make(chan int, 50)
	wg2.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)

		wg2.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 20
		ch <- 30
		wg2.Done()
	}(ch)

	wg2.Wait()
}
