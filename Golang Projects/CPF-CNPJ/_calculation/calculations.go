package calculation

import (
	function "test/_function"
	"test/formatter/cpf/alias"
)

func WeightedSum(numberAsText string, listOfWeights []int) int {
	weightedElements := ApplyWeights(numberAsText, listOfWeights)

	summedDigits := SumSliceElements(weightedElements)

	return summedDigits
}

func ApplyWeights(numbersIntStringFormat string, listOfWeights []int) []int {
	weightedElements := []int{}

	if len(numbersIntStringFormat) != len(listOfWeights) {
		return weightedElements
	}

	for currentIndex, currentWeight := range listOfWeights {
		unweightedNumber := ParseByteToSring(numbersIntStringFormat[currentIndex])

		valueAsInt, convertToIntException := function.ToInt(unweightedNumber)

		if alias.IsNullNot(convertToIntException) {
			return weightedElements
		}

		appliedWeight := (valueAsInt * currentWeight)
		weightedElements = append(weightedElements, appliedWeight)
	}

	return weightedElements
}

func SumSliceElements(numbersList []int) int {
	numbersSum := 0

	for _, number := range numbersList {
		numbersSum += number
	}

	return numbersSum
}

func ParseByteToSring(integerValue byte) string {
	return string(integerValue)
}
