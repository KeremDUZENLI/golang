package maps

import "fmt"

func OddNumbers() {
	var number int
	var liste []int
	var total int

	fmt.Print("Until Which Number: ")
	fmt.Scanln(&number)

	for i := 1; i <= number; i++ {
		liste = append(liste, i)
	}

	for _, value := range liste {
		if value%2 == 1 {
			total += value
		}
	}

	fmt.Println(liste)
	fmt.Printf("Total of odd numbers: %v\n", total)
}

func OddNumbersMap() {
	var itemNumber int
	var key string
	var value string

	type dictionary struct {
		key_   string
		value_ string
	}
	dictionaryTemp := dictionary{}
	dictionaryFinal := []dictionary{}

	fmt.Print("How many items: ")
	fmt.Scanln(&itemNumber)

	for i := 0; i < itemNumber; i++ {
		fmt.Scanln(&key)
		fmt.Scanln(&value)

		dictionaryTemp.key_ = key
		dictionaryTemp.value_ = value

		dictionaryFinal = append(dictionaryFinal, dictionaryTemp)
	}

	fmt.Println(dictionaryFinal)
}
