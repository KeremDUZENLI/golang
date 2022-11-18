package unmask

import (
	"fmt"
	exception "test/_exception"
	function "test/_function"
	alias "test/formatter/cpf/alias"
	constant "test/formatter/cpf/constant"
)

func Unmask(maskedCpf string) (string, exception.IException) {

	if alias.IsEmpty(maskedCpf) {
		return constant.UNMASKED_CPF_WITH_ZEROS, alias.NullException()
	}

	trimmedMaskedCpf := alias.Trim(maskedCpf)

	if alias.IsEmpty(trimmedMaskedCpf) {
		return constant.UNMASKED_CPF_WITH_ZEROS, alias.NullException()
	}

	cpfWithOnlyNumbers := alias.OnlyNumbers(trimmedMaskedCpf)

	if alias.IsEmpty(cpfWithOnlyNumbers) {
		return constant.MASKED_CPF_WITH_ZEROS, alias.NullException()
	}

	cpfConvertedToInt64, cpfConvertedToInt64Exception := function.ToInt64(cpfWithOnlyNumbers)

	if alias.IsNullNot(cpfConvertedToInt64Exception) {
		return constant.MASKED_CPF_WITH_ZEROS, cpfConvertedToInt64Exception
	}

	if cpfConvertedToInt64 == 0 {
		return constant.UNMASKED_CPF_WITH_ZEROS, alias.NullException()
	}

	unmaskedCpfWithLeftZeros := fmt.Sprintf("%011d", cpfConvertedToInt64)

	return unmaskedCpfWithLeftZeros, alias.NullException()
}
