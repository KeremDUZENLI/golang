package Challenge08_CheckPermutations

import (
	"sort"
)

func CheckPermutations1(str1, str2 string) bool {
	var res bool
	var str1Len int = len(str1)
	var str2Len int = len(str2)

	var str1Liste []string
	var str2Liste []string

	if str1Len != str2Len {
		res = false
		return res
	} else {
		for _, char := range str1 {
			str1Liste = append(str1Liste, string(char))
		}

		for _, char := range str2 {
			str2Liste = append(str2Liste, string(char))
		}

		for i := 0; i < str1Len; i++ {
			// fmt.Println(string(str1Liste[i]))
			// fmt.Println(string(str2Liste[str2Len-i-1]))
			if str1Liste[i] == str2Liste[str2Len-i-1] {
				res = true
			} else {
				res = false
			}
		}
		return res
	}
}

func CheckPermutations2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)

	return sortedStr1 == sortedStr2
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func Test() bool {
	str1 := "abcde"
	str2 := "medac"

	return CheckPermutations1(str1, str2)
}
