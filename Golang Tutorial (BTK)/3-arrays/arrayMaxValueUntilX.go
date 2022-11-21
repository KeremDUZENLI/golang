package arrays

import (
	"fmt"
	"strconv"
)

func MaxValueUntilX() {
	var looping = true
	var inputVal string
	var listt []int
	var maxValue int

	for looping {
		fmt.Scanln(&inputVal)

		if inputVal == "x" {
			looping = false
		} else {
			a, _ := strconv.Atoi(inputVal)
			listt = append(listt, a)
		}
	}
	fmt.Printf("\nList : %v", listt)

	for i := 0; i < len(listt); i++ {
		if maxValue < listt[i] {
			maxValue = listt[i]
		}
	}
	fmt.Printf("\nMax Value : %v", maxValue)
}
