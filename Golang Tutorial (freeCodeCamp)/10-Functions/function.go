package Functions

import (
	"fmt"
)

func Function() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello: ", i)
	}

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("TOTAL: ", s)

	d, err := divide(5.0, 0.0)
	fmt.Println(d, err)

	// Anonymous Function
	func() {
		fmt.Println("Anonymous")
	}()
}

// Local Functions
func sayMessage(msg string, idx int) {
	fmt.Println(msg, idx)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}
