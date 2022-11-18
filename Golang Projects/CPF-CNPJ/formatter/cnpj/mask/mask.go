package mask

import (
	"regexp"

	exception "test/_exception"
	function "test/_function"
	alias "test/formatter/cnpj/alias"
	constant "test/formatter/cnpj/constant"
	exceptionInt "test/formatter/cnpj/exception"
)

func Mask(unmaskedCnpj string) (string, exception.IException) {
	if function.IsEmpty(unmaskedCnpj) {
		return constant.MASKED_CNPJ_WITH_ZEROS, exceptionInt.MaskAnEmptyCnpjException()
	}

	if alias.Length(unmaskedCnpj) > constant.UNMASKED_CNPJ_MAX_LENGHT {
		return constant.MASKED_CNPJ_WITH_ZEROS, exceptionInt.MaskAnCnpjExceedsMaxLengthException(unmaskedCnpj)
	}

	trimmedCnpj := alias.Trim(unmaskedCnpj)

	if function.IsEmpty(trimmedCnpj) {
		return constant.MASKED_CNPJ_WITH_ZEROS, exceptionInt.MaskTrimmedCnpjEmptyException()
	}

	cnpjWithOnlyNumbers := alias.OnlyNumbers(unmaskedCnpj)

	if function.IsEmpty(cnpjWithOnlyNumbers) {
		return constant.MASKED_CNPJ_WITH_ZEROS, exceptionInt.MaskCnpjWithoutNumbersException(unmaskedCnpj)
	}

	const cnpjFormatPattern string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`

	cnpjRegEx, cnpjRegExError := regexp.Compile(cnpjFormatPattern)

	if cnpjRegExError != nil {
		return constant.MASKED_CNPJ_WITH_ZEROS, exceptionInt.MaskCnpjRegExCompileException(cnpjRegExError)
	}

	cnpjMasked := cnpjRegEx.ReplaceAllString(cnpjWithOnlyNumbers, "$1.$2.$3-$4")

	return cnpjMasked, alias.NullException()
}
