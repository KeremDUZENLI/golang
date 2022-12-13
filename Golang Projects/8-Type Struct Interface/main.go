package main

import (
	"example/hiddenStruct"
	"fmt"
)

func main() {
	// ### VERSION 1 ###
	/*
		var varStruct hiddenStruct.AStruct
		varStruct.Name = "Kerem"

		var varInterface hiddenStruct.AInterface
		varInterface = &varStruct
	*/

	// ### VERSION 2 ###
	varStruct := hiddenStruct.AStruct{
		Name: "Kerem",
	}

	varInterface := varStruct

	fmt.Println(varInterface.AFunction())
	fmt.Println(hiddenStruct.Create())
}
