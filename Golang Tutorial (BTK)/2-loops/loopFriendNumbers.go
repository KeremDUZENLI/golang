package loops

import "fmt"

type table struct {
	NumberVal     int
	DivisionTotal int
}

var (
	numberInput   int
	divides_total int
	liste         []table
)

func FriendNumbers() {
	fmt.Println("Until which number: ")
	fmt.Scanln(&numberInput)

	for start_number := 1; start_number < numberInput; start_number++ {
		divides_total = 0
		for i := 1; i < start_number; i++ {
			if start_number%i == 0 {
				divides_total += i
				// fmt.Printf("Number: %v Devides: %v \n", start_number, i)
			}
		}
		// fmt.Printf("Number: %v Devides_Total: %v \n\n", start_number, divides_total)
		var NewStruct table
		NewStruct.NumberVal = start_number
		NewStruct.DivisionTotal = divides_total
		// fmt.Println(NewStruct)

		liste = append(liste, NewStruct)
	}

	// fmt.Println(liste)
	// fmt.Println(liste[0].NumberVal)

	for a := 0; a < len(liste); a++ {
		for b := 0; b < len(liste); b++ {
			if liste[a].NumberVal == liste[b].DivisionTotal && liste[a].DivisionTotal == liste[b].NumberVal && liste[a].NumberVal != liste[b].NumberVal {
				fmt.Printf("%v and %v Are Friend Numbers \n", liste[a].NumberVal, liste[b].NumberVal)
				liste = append(liste[:a], liste[a+1:]...) // Just to not print again same numbers
			}
		}
	}
}
