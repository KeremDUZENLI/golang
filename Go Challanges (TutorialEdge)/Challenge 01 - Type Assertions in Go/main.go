package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	fmt.Println("Implement Me")

	var dev Developer

	dev.Name = name.(string)
	dev.Age = age.(int)

	return dev
}

func main() {
	fmt.Println("Hello World")

	var name interface{} = "Kerem"
	var age interface{} = 25

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
}
