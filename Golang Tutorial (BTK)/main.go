package main

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
	// maps.OddNumbersMap()

	/* 7-pointers_structs */
	// x := 10
	// pointers_structs.Pointer1(&x)
	// fmt.Printf("Main(): %v\n", x)

	// numbers := []int{1, 2, 3}
	// pointers_structs.Pointer2(numbers)
	// fmt.Printf("Main(): %v\n", numbers)

	// pointers_structs.Struct1()
	// pointers_structs.Struct2()

	/* 8-goRoutine_goChannel */
	// go goRoutine_goChannel.EvenNumbers()
	// go goRoutine_goChannel.OddNumbers()
	// time.Sleep(1 * time.Second)
	// fmt.Printf("\nMAIN...Routines\n\n")

	// even := make(chan int)
	// odd := make(chan int)
	// go goRoutines_goChannels.EvenNumbers_(even)
	// go goRoutines_goChannels.OddNumbers_(odd)

	// evenTotal, oddTotal := <-even, <-odd
	// fmt.Printf("%v, %v", evenTotal, oddTotal)
	// fmt.Printf("\nMAIN...Channels\n\n")

	/* 9-interfaces */
	// interfaces.Interface1()
	// interfaces.Interface2()

	/* 10-defer_debug */
	// defer_debug.Defer11()
	// defer_debug.Defer12()
	// defer_debug.Defer13()
	// defer_debug.Defer2()
	// defer_debug.Defer3()
	// defer_debug.Error1()
	// defer_debug.Error2()
	// defer_debug.Error3()
	// fmt.Println(defer_debug.Error4(102))

	/* 11-functions_strings */
	// functions_strings.FunctionString()

	/* 12-Restful */
	// fmt.Printf("\nGET: \n")
	// restful.RestfulGet()

	// fmt.Printf("\nPOST: \n")
	// restful.RestfulPost()

	/* PROJECT */
	// product, _ := project.AddProduct()
	// fmt.Println(product)

	// products, _ := project.GetAllProducts()
	// fmt.Println(products)

	// for i := 0; i < len(products); i++ {
	// 	fmt.Println(products[i].ProductName)
	// }
}
