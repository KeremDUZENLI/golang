package loops

import "fmt"

func Loop1() {
	var (
		number           int = 80
		numberEstimated  int
		amountEstimation int
		success          string
	)

	fmt.Println("Estimated Number : ")
	fmt.Scanln(&numberEstimated)

	for number != numberEstimated {
		if numberEstimated < number {
			amountEstimation++
			fmt.Printf("%v is Low\n", numberEstimated)
			fmt.Scanln(&numberEstimated)
		}

		if numberEstimated > number {
			amountEstimation++
			fmt.Printf("%v is High\n", numberEstimated)
			fmt.Scanln(&numberEstimated)
		}
	}

	if number == numberEstimated {
		amountEstimation++

		if amountEstimation <= 3 {
			success = "Good"
		} else if amountEstimation <= 6 {
			success = "Medium"
		} else {
			success = "Bad"
		}

		fmt.Printf("%v Found!\n", numberEstimated)
		fmt.Printf("%v Try %v \n ", amountEstimation, success)
	}
}
