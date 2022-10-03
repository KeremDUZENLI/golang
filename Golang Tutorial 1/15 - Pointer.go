package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal(myPtr *int) {
	*myPtr = 12
}

func doubleArrValue(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func main() {
	f1 := 10
	pl("f1 before func: ", f1)

	changeVal(&f1)
	pl("f1 after func: ", f1)

	f2 := 100
	var f2Ptr *int = &f2
	pl("f2 Address: ", f2Ptr)
	pl("f2 Value: ", *f2Ptr)

	pl("f2 before func: ", f2)
	changeVal(&f2)
	pl("f2 after func: ", f2)

	pArr := [4]int{1, 2, 3, 4}
	doubleArrValue(&pArr)
	pl(pArr)

	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average: %.3f\n", getAverage(iSlice...))
}
