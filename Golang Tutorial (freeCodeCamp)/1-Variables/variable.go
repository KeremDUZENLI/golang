package Variables

import (
	"fmt"
	"strconv"
)

func Variable1() {
	var x int
	x = 1

	var y int = 2

	z := 3

	fmt.Println(x, y, z)
	fmt.Printf("%v, %T\n", z, z)
}

var (
	actorName    string = "Elisabeth"
	companion    string = "Sarah"
	doctorNumber int    = 3
	season       int    = 11
)

func Variable2() {
	println(actorName, companion, doctorNumber, season)
}

func VariableConversion() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
