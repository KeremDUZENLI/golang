package main

import (
	"fmt"
)

var myVar string
var myVarChan = make(chan string)

func main() {
	res := scan()
	println(res)

	go channel()
	println(myVar, <-myVarChan)
}

func scan() string {
	var value string
	print("Value :   ")
	fmt.Scanln(&value)

	return value
}

func channel() {
	x := scan()
	y := scan()

	myVar = x
	myVarChan <- y
}
