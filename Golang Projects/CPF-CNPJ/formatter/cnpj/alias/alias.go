package alias

import (
	exception "test/_exception"
	function "test/_function"
)

type (
	ExceptionType      = exception.ExceptionType
	InterfaceException = exception.IException
)

var EMPTY_STRING, NullException, IsNull, IsNullNot, IsEmpty, IsEmptyNot, Trim, Length, OnlyNumbers = function.GetCommonFunctions()
