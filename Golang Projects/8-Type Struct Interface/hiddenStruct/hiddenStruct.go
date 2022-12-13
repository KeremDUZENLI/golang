package hiddenStruct

type AType string

type AStruct struct {
	Name AType
}

type AInterface interface {
	AFunction() AType
}

func (aliasForStruct *AStruct) AFunction() AType {
	return aliasForStruct.Name
}

func Create() (AStruct, AInterface) {
	var NewStruct AStruct
	NewStruct.Name = "XXX"

	var NewInterface AInterface
	NewInterface = &NewStruct

	return NewStruct, NewInterface
}
