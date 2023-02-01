package main

import (
	"TutorialEdge/Challenge01_TypeAssertionsInGo"
	"TutorialEdge/Challenge02_SatisfyingInterfacesInGo"
	"TutorialEdge/Challenge03_SortingFlightsByPrice"
	"TutorialEdge/Challenge04_CheckingForDuplicates"
	"TutorialEdge/Challenge05_ImplementingAStackInGo"
	"TutorialEdge/Challenge06_ImplementingAQueueInGo"
	"fmt"
)

func main() {
	fmt.Println(Challenge01_TypeAssertionsInGo.GetDeveloper("Kerem", 20))

	fmt.Println(Challenge02_SatisfyingInterfacesInGo.Programmer())

	fmt.Println(Challenge03_SortingFlightsByPrice.SortedList())

	fmt.Println(Challenge04_CheckingForDuplicates.UniquesList())

	fmt.Println(Challenge05_ImplementingAStackInGo.Test())
	fmt.Println(Challenge05_ImplementingAStackInGo.Test().Peek())

	fmt.Println(Challenge06_ImplementingAQueueInGo.Test())
	fmt.Println(Challenge06_ImplementingAQueueInGo.Test().Peek())
}
