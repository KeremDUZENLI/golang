package pointers_structs

import "fmt"

type product struct {
	name  string
	price float64
	brand string
}

func Struct1() {
	struct1 := product{"Laptop", 1000, "HP"}
	struct2 := product{name: "Desktop", price: 2000}

	fmt.Println(struct1)
	fmt.Println(struct2)
}
