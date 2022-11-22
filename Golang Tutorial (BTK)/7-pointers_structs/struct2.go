package pointers_structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) add() {
	fmt.Println("Added: ", c)
}

func Struct2() {
	a := customer{firstName: "A", lastName: "B", age: 20}
	a.add()
}
