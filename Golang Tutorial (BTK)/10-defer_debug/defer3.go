package defer_debug

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Printf("Saved: %v\n", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Log")
}

func Defer3() {
	p := product{productName: "Laptop", unitPrice: 1000}
	defer p.save()

	p = product{productName: "Mouse", unitPrice: 50}
	fmt.Println("Done...")
}
