package exceptionInt

import (
	exception "test/_exception"
)

const (
	MASK_CNPJ_WITHOUT_NUMBERS    exception.ExceptionType = "MASK_CNPJ_WITHOUT_NUMBERS"
	MASK_CNPJ_IS_EMPTY           exception.ExceptionType = "MASK_CNPJ_IS_EMPTY"
	MASK_CNPJ_EXCEEDS_MAX_LENGTH exception.ExceptionType = "MASK_CNPJ_EXCEEDS_MAX_LENGTH"
	MASK_CNPJ_TRIMMED_IS_EMPTY   exception.ExceptionType = "MASK_CNPJ_TRIMMED_IS_EMPTY"
	MASK_CNPJ_REGEX_ERROR        exception.ExceptionType = "MASK_CNPJ_REGEX_ERROR"
	UNMASK_CNPJ_WITHOUT_NUMBERS  exception.ExceptionType = "UNMASK_CNPJ_WITHOUT_NUMBERS"
	INVALID_CNPJ                 exception.ExceptionType = "INVALID_CNPJ"
)

var MaskAnEmptyCnpjException = exception.New(MASK_CNPJ_IS_EMPTY).
	Description("trying to mask a CNPJ that is empty").
	Message("trying to mask a CNPJ that is empty")

var MaskTrimmedCnpjEmptyException = exception.New(MASK_CNPJ_TRIMMED_IS_EMPTY).
	Description("trying to mask a trimmed CNPJ that is empty").
	Message("trying to mask a trimmed CNPJ that is empty")

var MaskAnCnpjExceedsMaxLengthException = exception.New(MASK_CNPJ_EXCEEDS_MAX_LENGTH).
	Description("the CNPJ has more than 11 digits").
	Message("the CNPJ %s that has more than 11 digits")

var MaskCnpjWithoutNumbersException = exception.New(MASK_CNPJ_WITHOUT_NUMBERS).
	Description("trying to mask a CNPJ that has no numbers").
	Message("trying to mask a CNPJ %s that has no numbers")

var MaskCnpjRegExCompileException = exception.New(MASK_CNPJ_REGEX_ERROR).
	Description("trying to mask CNPJ with RegEx").
	Message("trying to mask CNPJ with Regex - %s")

var InvalidCnpjException = exception.New(INVALID_CNPJ).
	Description("invalid CNPJ").
	Message("invalid CNPJ")
