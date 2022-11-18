package validate

import (
	validationJ "test/formatter/cnpj/validation"
	validation "test/formatter/cpf/validation"
)

const (
	cpf  = 1
	cnpj = 2
	err  = 0
)

func CpfOrCnpj(cpfCnpj string) int {
	cpfValue, _ := validation.IsValid(cpfCnpj)
	cnpjValue, _ := validationJ.IsValid(cpfCnpj)

	if cpfValue {
		return cpf
	} else if cnpjValue {
		return cnpj
	} else {
		return err
	}
}
