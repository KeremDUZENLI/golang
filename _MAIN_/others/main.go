package main

import (
	"fmt"
	"strings"
)

func main() {
	var Liste []any
	Pointer("pointer", &Liste)
	fmt.Println(Liste)

	x := Input1()
	fmt.Println(x)

	var y, z string
	Input2(&y)
	Input2(&z)
	fmt.Println(y, z)
}

func Pointer(str string, listeOut *[]any) {
	splt := strings.Split(str, "")
	for _, v := range splt {
		*listeOut = append(*listeOut, v)
	}
}

func Input1() string {
	var path string
	print("Folder Path:   ")
	fmt.Scanln(&path)

	return path
}

func Input2(x *string) {
	var path string
	print("Folder Path:   ")
	fmt.Scanln(&path)

	*x = path
}
