package main

import (
	"TutorialEdge/Challenge01_TypeAssertionsInGo"
	"TutorialEdge/Challenge02_SatisfyingInterfacesInGo"
	"TutorialEdge/Challenge03_SortingFlightsByPrice"
	"TutorialEdge/Challenge04_CheckingForDuplicates"
	"TutorialEdge/Challenge05_ImplementingAStackInGo"
	"TutorialEdge/Challenge06_ImplementingAQueueInGo"
	"TutorialEdge/Challenge07_MinimumsMaximumsAndErrors"
	"TutorialEdge/Challenge08_CheckPermutations"
	"TutorialEdge/Challenge09_SinglyLinkedLists"
	"TutorialEdge/Challenge10_WordFrequencies"
	"fmt"
)

func main() {
	fmt.Println("\nChallenge01_TypeAssertionsInGo: ")
	fmt.Println(Challenge01_TypeAssertionsInGo.GetDeveloper("Kerem", 20))

	fmt.Println("\nChallenge02_SatisfyingInterfacesInGo: ")
	fmt.Println(Challenge02_SatisfyingInterfacesInGo.Programmer())

	fmt.Println("\nChallenge03_SortingFlightsByPrice: ")
	fmt.Println(Challenge03_SortingFlightsByPrice.SortedList())

	fmt.Println("\nChallenge04_CheckingForDuplicates: ")
	fmt.Println(Challenge04_CheckingForDuplicates.UniquesList())

	fmt.Println("\nChallenge05_ImplementingAStackInGo: ")
	fmt.Println(Challenge05_ImplementingAStackInGo.Test())
	fmt.Println(Challenge05_ImplementingAStackInGo.Test().Peek())

	fmt.Println("\nChallenge06_ImplementingAQueueInGo: ")
	fmt.Println(Challenge06_ImplementingAQueueInGo.Test())
	fmt.Println(Challenge06_ImplementingAQueueInGo.Test().Peek())

	fmt.Println("\nChallenge07_MinimumsMaximumsAndErrors: ")
	fmt.Println(Challenge07_MinimumsMaximumsAndErrors.Test())

	fmt.Println("\nChallenge08_CheckPermutations: ")
	fmt.Println(Challenge08_CheckPermutations.Test())

	fmt.Println("\nChallenge09_SinglyLinkedLists: ")
	fmt.Println(Challenge09_SinglyLinkedLists.Test())

	fmt.Println("\nChallenge10_WordFrequencies: ")
	fmt.Println(Challenge10_WordFrequencies.Test())
}
