package validation

import (
	"fmt"
	calculation "test/_calculation"
	exception "test/_exception"
	function "test/_function"
	alias "test/formatter/cpf/alias"
	constant "test/formatter/cpf/constant"
	exceptionInt "test/formatter/cpf/exception"
	unmask "test/formatter/cpf/unmask"
)

func AllDigitsAreEqual(document string) bool {
	for _, digit := range document {
		if document[0] != byte(digit) {
			return false
		}
	}

	return true
}

func validateFirstCpfDigits(cpf string) string {
	first90Digits := getFirst90Digits(cpf)
	resultingdigit := 0

	cpfFirstDigitTable := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	digitsSum := calculation.WeightedSum(first90Digits, cpfFirstDigitTable)

	divisionRest := digitsSum % 11

	if divisionRest >= 2 {
		resultingdigit = 11 - divisionRest
	}

	return concatenateCpfWithDigit(first90Digits, resultingdigit)
}

func validateLastCpfDigits(firstValidatedDigits string) string {
	cpfSecondDigitTable := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	digitsSum := calculation.WeightedSum(firstValidatedDigits, cpfSecondDigitTable)
	cpfLastRest := 0

	cpfSecondSum := digitsSum % 11

	if cpfSecondSum >= 2 {
		cpfLastRest = 11 - cpfSecondSum
	}

	return concatenateCpfWithDigit(firstValidatedDigits, cpfLastRest)
}

func finalCheck(lastValidatedDigits, Cpf string) exception.IException {
	if lastValidatedDigits == Cpf {
		return nil
	}

	return exceptionInt.InvalidCpfException()
}

func getFirst90Digits(cpf string) string {
	return cpf[:9]
}

func concatenateCpfWithDigit(cpf string, digit int) string {
	return fmt.Sprintf("%s%d", cpf, digit)
}

func calculateIfCpfIsValid(cpf string) exception.IException {
	firstValidatedDigits := validateFirstCpfDigits(cpf)

	lastValidatedDigits := validateLastCpfDigits(firstValidatedDigits)

	invalidCpfException := finalCheck(lastValidatedDigits, cpf)

	if alias.IsNullNot(invalidCpfException) {
		return invalidCpfException
	}

	return nil
}

func IsValid(cpf string) (bool, exception.IException) {
	formattedCpf, couldNotUnmaskException := unmask.Unmask(cpf)

	if alias.IsNullNot(couldNotUnmaskException) {
		return false, couldNotUnmaskException
	}

	if function.Length(formattedCpf) != constant.UNMASKED_CPF_MAX_LENGHT {
		return false, exceptionInt.InvalidCpfException()
	}

	if AllDigitsAreEqual(formattedCpf) {
		return false, exceptionInt.InvalidCpfException()
	}

	InvalidCpfException := calculateIfCpfIsValid(formattedCpf)

	if alias.IsNullNot(InvalidCpfException) {
		return false, InvalidCpfException
	}

	return true, nil
}
