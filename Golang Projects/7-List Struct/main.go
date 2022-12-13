package main

import (
	"basictest/main_basic"
	"basictest/main_list"
	"basictest/main_struct"
	"fmt"
)

func main() {
	fmt.Println(main_basic.AddBasic())

	fmt.Println(main_list.AddList())

	fmt.Println(main_struct.AddStruct1())
	fmt.Println(main_struct.AddStruct2())
}
