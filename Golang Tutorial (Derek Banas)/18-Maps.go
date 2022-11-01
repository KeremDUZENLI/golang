package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// var myMap map [keyType]valueType
	var heroes map[string]string

	heroes = make(map[string]string)

	villians := make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"

	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Crypto",
		2: "Bat Hound"}

	fmt.Printf("Batman is %v\n", heroes["Batman"])
	pl("Chip: ", superPets[3])

	_, ok := superPets[3]
	pl("Is there a 3rd pet: ", ok)

	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}
}
