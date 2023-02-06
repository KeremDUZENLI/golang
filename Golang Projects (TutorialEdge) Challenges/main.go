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
	"TutorialEdge/Challenge11_SetsAndSubsets"
	"TutorialEdge/Challenge12_ArmstrongNumbers"
	"TutorialEdge/Challenge13_SmallestDifferenceBetweenInts"
	"TutorialEdge/Challenge14_LeapYear"
	"TutorialEdge/Challenge15_RepeatingLetters"
	"TutorialEdge/Challenge16_OddOrEvenFactors"
	"TutorialEdge/Challenge17_DecodeTheSecret"
	"TutorialEdge/Challenge18_MinRotations"
	"TutorialEdge/Challenge19_CalculatingTheDifferenceBetweenSquares"
	"TutorialEdge/Challenge20_FindingTheNthTriangularNumber"
	"TutorialEdge/Challenge21_JSONAndStockDividends"
	"TutorialEdge/Challenge22_LargestPandigitalPrime"
	"TutorialEdge/Challenge23_LargestNPrime"
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

	fmt.Println("\nChallenge11_SetsAndSubsets: ")
	fmt.Println(Challenge11_SetsAndSubsets.Test())

	fmt.Println("\nChallenge12_ArmstrongNumbers: ")
	fmt.Println(Challenge12_ArmstrongNumbers.Armstrong(371))

	fmt.Println("\nChallenge13_SmallestDifferenceBetweenInts: ")
	fmt.Println(Challenge13_SmallestDifferenceBetweenInts.Test())

	fmt.Println("\nChallenge14_LeapYear: ")
	fmt.Println(Challenge14_LeapYear.CheckLeapYear(2020))

	fmt.Println("\nChallenge15_RepeatingLetters: ")
	fmt.Println(Challenge15_RepeatingLetters.DoubleChars("gophers"))

	fmt.Println("\nChallenge16_OddOrEvenFactors: ")
	fmt.Println(Challenge16_OddOrEvenFactors.CheckFactors(8))
	fmt.Println(Challenge16_OddOrEvenFactors.OddEvenFactors(8))

	fmt.Println("\nChallenge17_DecodeTheSecret: ")
	fmt.Println(Challenge17_DecodeTheSecret.DecodeSecret("VEZEU0ZVVFVTSk9I"))

	fmt.Println("\nChallenge18_MinRotations: ")
	fmt.Println(Challenge18_MinRotations.MinRotations([]int{15, 18, 2, 3, 6, 12}))

	fmt.Println("\nChallenge19_CalculatingTheDifferenceBetweenSquares: ")
	fmt.Println(Challenge19_CalculatingTheDifferenceBetweenSquares.DiffSquares(1, 2))

	fmt.Println("\nChallenge20_FindingTheNthTriangularNumber: ")
	fmt.Println(Challenge20_FindingTheNthTriangularNumber.TriangularNumber(3))

	fmt.Println("\nChallenge21_JSONAndStockDividends: ")
	fmt.Println(Challenge21_JSONAndStockDividends.Test())

	fmt.Println("\nChallenge22_LargestPandigitalPrime: ")
	numberStr := "7654321"
	fmt.Println(Challenge22_LargestPandigitalPrime.Factorial(len(numberStr)))
	fmt.Println(Challenge22_LargestPandigitalPrime.RandomNumbers(numberStr))
	fmt.Println(Challenge22_LargestPandigitalPrime.RandomNumbersList(numberStr))
	fmt.Println(Challenge22_LargestPandigitalPrime.RandomNumbersListMax(numberStr))
	fmt.Println(Challenge22_LargestPandigitalPrime.CheckNumberPrimeLargest(numberStr))

	fmt.Println("\nChallenge23_LargestNPrime: ")
	number := 3
	fmt.Println(Challenge23_LargestNPrime.Factorial(number))
	fmt.Println(Challenge23_LargestNPrime.RandomNumbers(number))
	fmt.Println(Challenge23_LargestNPrime.RandomNumbersList(number))
	fmt.Println(Challenge23_LargestNPrime.RandomNumbersListMax(number))
	fmt.Println(Challenge23_LargestNPrime.CheckNumberPrimeLargest(number))
}
