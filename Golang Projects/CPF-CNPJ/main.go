package main

import (
	"fmt"

	calculation "test/_calculation"
	exception "test/_exception"
	function "test/_function"
	validate "test/_validate"

	alias "test/formatter/cpf/alias"
	constant "test/formatter/cpf/constant"
	exceptionInt "test/formatter/cpf/exception"
	mask "test/formatter/cpf/mask"
	unmask "test/formatter/cpf/unmask"
	validation "test/formatter/cpf/validation"

	aliasJ "test/formatter/cnpj/alias"
	constantJ "test/formatter/cnpj/constant"
	exceptionIntJ "test/formatter/cnpj/exception"
	maskJ "test/formatter/cnpj/mask"
	unmaskJ "test/formatter/cnpj/unmask"
	validationJ "test/formatter/cnpj/validation"
)

func main() {
	// _calculation
	fmt.Println("\n _calculation :")
	fmt.Println(calculation.ApplyWeights("123", []int{1, 2, 3}))
	fmt.Println(calculation.WeightedSum("123", []int{1, 2, 3}))
	fmt.Println(calculation.SumSliceElements([]int{1, 2, 3}))
	fmt.Println(calculation.ParseByteToSring(1))

	// _exception
	fmt.Println("\n _exception :")
	var newType exception.ExceptionType
	var newStruct exception.Exception

	fmt.Println(newStruct.Output())
	fmt.Println(newStruct.GetType())
	fmt.Println(newStruct.IsType(newType))
	fmt.Println(newStruct.Description("description"))
	fmt.Println(newStruct.Message("message"))

	// _function
	fmt.Println("\n _function :")
	fmt.Println(function.IsEmpty("asd"))
	fmt.Println(function.Trim(" a s d "))
	fmt.Println(function.Length("asd"))
	fmt.Println(function.OnlyNumbers("asd123asd"))
	fmt.Println(function.ToInt("123"))
	fmt.Println(function.ToInt32("asd123asd"))
	fmt.Println(function.ToInt64("asd123asd"))
	fmt.Println(function.GetCommonFunctions())

	// cpf/alias
	fmt.Println("\n alias :")
	fmt.Println(alias.EMPTY_STRING, alias.NullException)

	// cnpj/alias
	fmt.Println("\n aliasJ :")
	fmt.Println(aliasJ.EMPTY_STRING, alias.NullException)

	// cpf/constant
	fmt.Println("\n constant :")
	fmt.Println(constant.UNMASKED_CPF_MAX_LENGHT)

	// cnpj/constant
	fmt.Println("\n constantJ :")
	fmt.Println(constantJ.UNMASKED_CNPJ_MAX_LENGHT)

	// cpf/exception
	fmt.Println("\n exceptionInt :")
	fmt.Println(exceptionInt.MASK_CPF_IS_EMPTY)
	fmt.Println(exceptionInt.MaskAnEmptyCpfException)

	// cnpj/exception
	fmt.Println("\n exceptionIntJ :")
	fmt.Println(exceptionIntJ.MASK_CNPJ_IS_EMPTY)
	fmt.Println(exceptionIntJ.MaskAnEmptyCnpjException)

	// cpf/mask
	fmt.Println("\n mask :")
	fmt.Println(mask.Mask("asd123asd"))

	// cnpj/mask
	fmt.Println("\n maskJ :")
	fmt.Println(maskJ.Mask("asd123asd"))

	// cpf/unmask
	fmt.Println("\n unmask :")
	fmt.Println(unmask.Unmask("asd123asd"))

	// cnpj/unmask
	fmt.Println("\n unmaskJ :")
	fmt.Println(unmaskJ.Unmask("asd123asd"))

	// cpf/validation
	fmt.Println("\n validation :")
	fmt.Println(validation.AllDigitsAreEqual("asd123asd"))
	fmt.Println(validation.IsValid("111.111.111-11"))

	// cnpj/validation
	fmt.Println("\n validationJ :")
	fmt.Println(validationJ.AllDigitsAreEqual("asd123asd"))
	fmt.Println(validationJ.IsValid("11.111.111/1111-11"))

	// FINAL
	fmt.Println("\n FINAL :")
	fmt.Println(validate.CpfOrCnpj("cpf"))
}
