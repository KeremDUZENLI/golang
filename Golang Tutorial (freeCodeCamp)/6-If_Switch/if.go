package If_Switch

import (
	"fmt"
)

func NumberGuess1() {
	number := 50
	guess := 30

	if guess < number {
		fmt.Println("Too low")
	} else if guess > number {
		fmt.Println("Too high")
	} else if guess == number {
		fmt.Println("Same")
	}

	fmt.Println(guess < number, guess > number, guess == number)
}

func NumberGuess2() {
	guess := 35

	if guess < 1 || guess > 100 {
		fmt.Println("Guess Must Be Between 1 - 100")
	}

	if guess > 30 && guess < 50 {
		fmt.Println("Between 30 and 50")
	}
}
