package main

import (
	"fmt"
)

var pl = fmt.Println

/*
type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

func main() {
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 main st"
	tS.bal = 234.56

	getCustInfo(tS)
	newCustAdd(&tS, "123 South st")
	pl("Address : ", tS.address)

	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name : ", sS.name)
}
*/

type rectangle struct {
	lenght, height float64
}

func (r rectangle) Area() float64 {
	return r.lenght * r.height
}

func main() {
	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area : ", rect1.Area())
}
