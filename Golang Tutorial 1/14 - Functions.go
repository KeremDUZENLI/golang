package main

import (
	"fmt"
)

var pl = fmt.Println

func sayHello() {
	pl("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x / y, nil
	}
}

func getSumVaradic(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	// func funcName(parameters) returnType {BODY}
	sayHello()
	pl(getSum(5, 4))
	pl(getTwo(10))
	pl(getQuotient(5, 2))
	pl(getSumVaradic(1, 2, 3, 4, 5))

	vArr := []int{1, 2, 3, 4}
	pl("Array Sum: ", getArraySum(vArr))
}
