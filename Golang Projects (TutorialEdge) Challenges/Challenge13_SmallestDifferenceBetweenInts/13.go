package Challenge13_SmallestDifferenceBetweenInts

import (
	"math"
	"sort"
)

func CalcSmallestDifference(arr1, arr2 []int) int {
	var smallest []int

	for _, value1 := range arr1 {
		for _, value2 := range arr2 {
			smallest = append(smallest, value1-value2)
		}
	}

	for i := range smallest {
		smallest[i] = int(math.Abs(float64(smallest[i])))
	}

	sort.Ints(smallest)
	return smallest[0]
}

func Test() int {
	arr1 := []int{9, 8, 7, 6}
	arr2 := []int{7, 3, 2, 5}

	return CalcSmallestDifference(arr1, arr2)
}
