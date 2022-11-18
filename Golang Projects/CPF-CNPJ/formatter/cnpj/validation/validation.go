package validation

import (
	"fmt"
	calculation "test/_calculation"
	exception "test/_exception"
	function "test/_function"
	alias "test/formatter/cnpj/alias"
	constant "test/formatter/cnpj/constant"
	exceptionInt "test/formatter/cnpj/exception"
	unmask "test/formatter/cnpj/unmask"
)

func AllDigitsAreEqual(document string) bool {
	for _, digit := range document {
		if document[0] != byte(digit) {
			return false
		}
	}

	return true
}

func validateFirstCnpjDigits(cnpj string) string {
	first12Digits := getFirst12Digits(cnpj)
	resultingdigit := 0

	cnpjFirstDigitTable := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	digitsSum := calculation.WeightedSum(first12Digits, cnpjFirstDigitTable)

	divisionRest := digitsSum % 11

	if divisionRest >= 2 {
		resultingdigit = 11 - divisionRest
	}

	return concatenateCnpjWithDigit(first12Digits, resultingdigit)
}

func validateLastCnpjDigits(firstValidatedDigits string) string {
	cnpjSecondDigitTable := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	digitsSum := calculation.WeightedSum(firstValidatedDigits, cnpjSecondDigitTable)
	cnpjLastRest := 0

	cnpjSecondSum := digitsSum % 11

	if cnpjSecondSum >= 2 {
		cnpjLastRest = 11 - cnpjSecondSum
	}

	return concatenateCnpjWithDigit(firstValidatedDigits, cnpjLastRest)
}

func finalCheck(lastValidatedDigits, Cnpj string) exception.IException {
	if lastValidatedDigits == Cnpj {
		return nil
	}

	return exceptionInt.InvalidCnpjException()
}

func getFirst12Digits(cnpj string) string {
	return cnpj[:12]
}

func concatenateCnpjWithDigit(cnpj string, digit int) string {
	return fmt.Sprintf("%s%d", cnpj, digit)
}

func calculateIfCnpjIsValid(cnpj string) exception.IException {
	firstValidatedDigits := validateFirstCnpjDigits(cnpj)

	lastValidatedDigits := validateLastCnpjDigits(firstValidatedDigits)

	invalidCnpjException := finalCheck(lastValidatedDigits, cnpj)

	if alias.IsNullNot(invalidCnpjException) {
		return invalidCnpjException
	}

	return nil
}

func IsValid(cnpj string) (bool, exception.IException) {
	formattedCnpj, couldNotUnmaskException := unmask.Unmask(cnpj)

	if alias.IsNullNot(couldNotUnmaskException) {
		return false, couldNotUnmaskException
	}

	if function.Length(formattedCnpj) != constant.UNMASKED_CNPJ_MAX_LENGHT {
		return false, exceptionInt.InvalidCnpjException(cnpj)
	}

	if AllDigitsAreEqual(formattedCnpj) {
		return false, exceptionInt.InvalidCnpjException(cnpj)
	}

	invalidCnpjException := calculateIfCnpjIsValid(formattedCnpj)

	if alias.IsNullNot(invalidCnpjException) {
		return false, invalidCnpjException
	}

	return true, nil
}
