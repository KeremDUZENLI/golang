package If_Switch

import (
	"fmt"
)

func Switch1() {
	switch 2 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("not 1 or 2")
	}
}

func Switch2() {
	i := 5
	switch {
	case i == 4:
		fmt.Println("4")
	case i == 5:
		fmt.Println("5")
	default:
		fmt.Println("not 4 or 5")
	}
}
