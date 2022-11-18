package exceptionInt

import (
	exception "test/_exception"
)

const (
	MASK_CPF_IS_EMPTY           exception.ExceptionType = "MASK_CPF_IS_EMPTY"
	MASK_CPF_TRIMMED_IS_EMPTY   exception.ExceptionType = "MASK_CPF_TRIMMED_IS_EMPTY"
	MASK_CPF_EXCEEDS_MAX_LENGTH exception.ExceptionType = "MASK_CPF_EXCEEDS_MAX_LENGTH"
	MASK_CPF_WITHOUT_NUMBERS    exception.ExceptionType = "MASK_CPF_WITHOUT_NUMBERS"
	MASK_CPF_REGEX_ERROR        exception.ExceptionType = "MASK_CPF_REGEX_ERROR"
	INVALID_CPF                 exception.ExceptionType = "INVALID_CPF"
)

var MaskAnEmptyCpfException = exception.New(MASK_CPF_IS_EMPTY).
	Description("trying to mask a CPF that is empty").
	Message("trying to mask a CPF that is empty")

var MaskTrimmedCpfEmptyException = exception.New(MASK_CPF_TRIMMED_IS_EMPTY).
	Description("trying to mask a trimmed CPF that is empty").
	Message("trying to mask a trimmed CPF that is empty")

var MaskAnCpfExceedsMaxLengthException = exception.New(MASK_CPF_EXCEEDS_MAX_LENGTH).
	Description("the CPF has more than 11 digits").
	Message("the CPF %s that has more than 11 digits")

var MaskCpfWithoutNumbersException = exception.New(MASK_CPF_WITHOUT_NUMBERS).
	Description("trying to mask a CPF that has no numbers").
	Message("trying to mask a CPF %s that has no numbers")

var MaskCpfRegExCompileException = exception.New(MASK_CPF_REGEX_ERROR).
	Description("trying to mask CPF with RegEx").
	Message("trying to mask CPF with Regex - %s")

var InvalidCpfException = exception.New(INVALID_CPF).
	Description("invalid CPF").
	Message("invalid CPF")
