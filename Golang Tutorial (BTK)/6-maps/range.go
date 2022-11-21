package maps

import "fmt"

func Range() {
	cities := []string{"Sao Paulo", "Santos", "Rio"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for index, value := range cities {
		fmt.Printf("index: %v / value: %v\n", index, value)
	}
}
