package functions

import "fmt"

func Function2(num1 int, num2 int) (int, int, int, float32) {
	addition := num1 + num2
	subtraction := num1 - num2
	multiplication := num1 * num2
	division := float32(num1 / num2)

	fmt.Println(addition, subtraction, multiplication, division)
	return addition, subtraction, multiplication, division
}
