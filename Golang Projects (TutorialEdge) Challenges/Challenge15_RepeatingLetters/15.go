package Challenge15_RepeatingLetters

import (
	"strings"
)

func DoubleChars(original string) string {
	splt := strings.Split(original, "")
	newList := []string{}

	for i := 0; i < len(splt); i++ {
		newList = append(newList, splt[i], splt[i])
	}

	returnString := strings.Join(newList, "")
	return returnString
}
