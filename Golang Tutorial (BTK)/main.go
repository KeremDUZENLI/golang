package main

import (
	maps "golesson/6-maps"
)

func main() {
	/* 1-conditionals */
	// conditionals.Compare(1, 2, 3)

	/* 2-loops */
	// loops.Loop1()
	// loops.Loop2()
	// loops.FriendNumbers()

	/* 3-arrays */
	// arrays.Array1()
	// arrays.Array2()
	// arrays.Array3()
	// arrays.MaxValueUntilX()

	/* 4-slices */
	// slices.Slice1()

	/* 5-functions */
	// functions.Function1(1, 3)

	// result1, result2, result3, result4 := functions.Function2(1, 2)
	// fmt.Printf("+ : %v\n- : %v\n* : %v\n/ : %v\n", result1, result2, result3, result4)

	// functions.VariadicFunction(1, 2, 3)
	// functions.VariadicFunction(1, 2, 3, 4, 5)
	// numberArray := []int{2, 3, 4, 7, 9, 15}
	// functions.VariadicFunction(numberArray...)

	/* 6-maps */
	// maps.Map()
	// maps.Range()
	// maps.OddNumbers()
	maps.OddNumbersMap()
}
