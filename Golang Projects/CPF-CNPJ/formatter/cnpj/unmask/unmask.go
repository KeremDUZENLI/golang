package unmask

import (
	"fmt"
	exception "test/_exception"
	function "test/_function"
	alias "test/formatter/cnpj/alias"
	constant "test/formatter/cnpj/constant"
)

func Unmask(maskedCnpj string) (string, exception.IException) {

	if alias.IsEmpty(maskedCnpj) {
		return constant.UNMASKED_CNPJ_WITH_ZEROS, alias.NullException()
	}

	trimmedMaskedCnpj := alias.Trim(maskedCnpj)

	if alias.IsEmpty(trimmedMaskedCnpj) {
		return constant.UNMASKED_CNPJ_WITH_ZEROS, alias.NullException()
	}

	cnpjWithOnlyNumbers := alias.OnlyNumbers(trimmedMaskedCnpj)

	if alias.IsEmpty(cnpjWithOnlyNumbers) {
		return constant.MASKED_CNPJ_WITH_ZEROS, alias.NullException()
	}

	cnpjConvertedToInt64, cnpjConvertedToInt64Exception := function.ToInt64(cnpjWithOnlyNumbers)

	if alias.IsNullNot(cnpjConvertedToInt64Exception) {
		return constant.MASKED_CNPJ_WITH_ZEROS, cnpjConvertedToInt64Exception
	}

	if cnpjConvertedToInt64 == 0 {
		return constant.UNMASKED_CNPJ_WITH_ZEROS, alias.NullException()
	}

	unmaskedCnpjWithLeftZeros := fmt.Sprintf("%011d", cnpjConvertedToInt64)

	return unmaskedCnpjWithLeftZeros, alias.NullException()
}
