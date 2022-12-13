package main

import (
	"fmt"
	"time"
)

func SendValue(c chan int) {
	fmt.Println("Start Goroutine")
	time.Sleep(1 * time.Second)
	c <- 8
	fmt.Println("End Goroutine")
}

func main() {
	fmt.Println("Channels: ")

	values := make(chan int, 2)
	defer close(values)

	go SendValue(values)
	go SendValue(values)

	value := <-values
	fmt.Println(value)
}
