package exception

import (
	"fmt"
)

type ExceptionType string

// addExceptionToCatalog Function
type AnExceptionsTypeCatalog map[ExceptionType]Exception

var ExceptionsTypeCatalog AnExceptionsTypeCatalog = AnExceptionsTypeCatalog{}

func addExceptionToCatalog(e *Exception) {
	eT := e.GetType()

	_, exceptionTypeAlreadyExists := ExceptionsTypeCatalog[eT]

	if exceptionTypeAlreadyExists {
		panic("exceptionType already exists")
	}

	ExceptionsTypeCatalog[eT] = *e
}

// addExceptionToCatalog Function

type Exception struct {
	output string

	getType ExceptionType

	description string

	message string
}

type IException interface {
	Output() string

	GetType() ExceptionType
	IsType(typee ExceptionType) bool

	Description(description string) IException

	Message(message string) func(args ...any) IException
}

func (e *Exception) Output() string {
	return e.output
}

func (e *Exception) GetType() ExceptionType {
	return e.getType
}
func (e *Exception) IsType(typee ExceptionType) bool {
	return typee == e.getType
}

func New(newExceptionType ExceptionType) IException {
	newExceptionStruct := &Exception{
		getType: newExceptionType,
	}

	addExceptionToCatalog(newExceptionStruct)

	return newExceptionStruct
}

func (e *Exception) Description(description string) IException {
	e.description = description
	return e
}

func (e *Exception) Message(message string) func(args ...any) IException {
	e.message = message
	e.output = message

	return func(args ...any) IException {
		e.withArgs(args...)
		return e
	}
}

func (e *Exception) withArgs(args ...any) IException {
	e.output = fmt.Sprintf(e.message, args...)
	return e
}
