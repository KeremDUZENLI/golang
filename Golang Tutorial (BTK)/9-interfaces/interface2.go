package interfaces

import "fmt"

type Mortgage struct {
	creditPayment float64
	adress        string
	rate          float64
}

type Car struct {
	creditPayment float64
	car           string
	rate          float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPayment * m.rate / 12
}

func (c Car) Calculate() float64 {
	return c.creditPayment * c.rate / 12
}

func MontlyPayment(credits []CreditCalculator) float64 {
	var paymentTotal float64

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}

	return paymentTotal
}

func Interface2() {
	m1 := Mortgage{creditPayment: 12000, adress: "abc", rate: 0.10}
	m2 := Mortgage{creditPayment: 12000, adress: "def", rate: 0.10}
	c1 := Car{creditPayment: 12000, car: "bmw", rate: .10}

	credits := []CreditCalculator{m1, m2, c1}
	total := MontlyPayment(credits)

	fmt.Printf("Total Credits: %v\n", total)
}
