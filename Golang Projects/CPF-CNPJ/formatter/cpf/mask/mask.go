package mask

import (
	"regexp"

	exception "test/_exception"
	function "test/_function"
	alias "test/formatter/cpf/alias"
	constant "test/formatter/cpf/constant"
	exceptionInt "test/formatter/cpf/exception"
)

func Mask(unmaskedCpf string) (string, exception.IException) {
	if function.IsEmpty(unmaskedCpf) {
		return constant.MASKED_CPF_WITH_ZEROS, exceptionInt.MaskAnEmptyCpfException()
	}

	if alias.Length(unmaskedCpf) > constant.UNMASKED_CPF_MAX_LENGHT {
		return constant.MASKED_CPF_WITH_ZEROS, exceptionInt.MaskAnCpfExceedsMaxLengthException(unmaskedCpf)
	}

	trimmedCpf := alias.Trim(unmaskedCpf)

	if function.IsEmpty(trimmedCpf) {
		return constant.MASKED_CPF_WITH_ZEROS, exceptionInt.MaskTrimmedCpfEmptyException()
	}

	cpfWithOnlyNumbers := alias.OnlyNumbers(unmaskedCpf)

	if function.IsEmpty(cpfWithOnlyNumbers) {
		return constant.MASKED_CPF_WITH_ZEROS, exceptionInt.MaskCpfWithoutNumbersException(unmaskedCpf)
	}

	const cpfFormatPattern string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`

	cpfRegEx, cpfRegExError := regexp.Compile(cpfFormatPattern)

	if cpfRegExError != nil {
		return constant.MASKED_CPF_WITH_ZEROS, exceptionInt.MaskCpfRegExCompileException(cpfRegExError)
	}

	cpfMasked := cpfRegEx.ReplaceAllString(cpfWithOnlyNumbers, "$1.$2.$3-$4")

	return cpfMasked, alias.NullException()
}
