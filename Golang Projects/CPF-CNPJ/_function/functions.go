package function

import (
	"strconv"
	"strings"

	exception "test/_exception"
)

const (
	EMPTY_STRING string = ""
	DEFAULT_CORS string = "*"
	BLANK_SPACE  string = " "
)

func getNullExceptionAsFunction() exception.IException {
	return nil
}

func getIsNullAsFunction(argToCheck any) bool {
	return argToCheck == nil
}

func getIsNullAsFunctionNot(argToCheck any) bool {
	return argToCheck != nil
}

// STR
func IsEmpty(text string) bool {
	return len(text) == 0
}

func IsEmptyNot(text string) bool {
	return len(text) > 0
}

func Trim(text string) string {
	textAsString := string(text)
	trimmedText := strings.TrimSpace(textAsString)

	return trimmedText
}

func Length(text string) int {
	return len(text)
}

func OnlyNumbers(text string) string {
	textWithonlyNumbers := EMPTY_STRING

	for _, characters := range text {
		letter := string(characters)

		if letter >= "0" && letter <= "9" {
			textWithonlyNumbers += letter
		}
	}

	return textWithonlyNumbers
}

func ToInt(valueAsText string) (int, exception.IException) {
	convertedToInt64, convertedToIntException := ToInt64(valueAsText)

	convertedToInt := int(convertedToInt64)

	return convertedToInt, convertedToIntException
}

func ToInt32(valueAsText string) (int32, exception.IException) {
	convertedToInt64, convertToIntException := ToInt64(valueAsText)

	convertedToInt32 := int32(convertedToInt64)

	return convertedToInt32, convertToIntException
}

func ToInt64(valueAsText string) (int64, exception.IException) {
	parsedToInt64, _ := strconv.ParseInt(valueAsText, 10, 64)

	return parsedToInt64, nil
}

// STR

func GetCommonFunctions() (
	string,

	func() exception.IException,
	func(checkIfIsNull any) bool,
	func(checkIfIsNotNull any) bool,
	func(str string) bool,
	func(str string) bool,
	func(text string) string,
	func(text string) int,
	func(text string) string) {

	return EMPTY_STRING,
		getNullExceptionAsFunction,
		getIsNullAsFunction,
		getIsNullAsFunctionNot,
		IsEmpty, IsEmptyNot, Trim, Length, OnlyNumbers
}
