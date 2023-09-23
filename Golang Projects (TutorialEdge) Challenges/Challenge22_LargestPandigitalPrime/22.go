package Challenge22_LargestPandigitalPrime

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Factorial(n int) int {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func RandomNumbers(s string) string {
	splt := strings.Split(s, "")
	strListe := make([]string, 0, len(s))
	seen := make(map[int]bool)

	for i := 0; i < len(s); {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(len(s))
		y := splt[x]
		randomNumber, _ := strconv.Atoi(y)

		if !seen[randomNumber] {
			seen[randomNumber] = true
			strListe = append(strListe, strconv.Itoa(randomNumber))
			i++
		}
	}

	return strings.Join(strListe, "")
}

func RandomNumbersList(s string) []string {
	maxPossible := Factorial(len(s))
	strListe := []string{}
	seen := make(map[string]bool)

	for i := 0; i < maxPossible; {
		randomNumber := RandomNumbers(s)

		if !seen[randomNumber] {
			seen[randomNumber] = true
			strListe = append(strListe, randomNumber)
			i++
		}
	}

	return strListe
}

func RandomNumbersListMax(s string) int {
	maxNum := 0

	for _, v := range RandomNumbersList(s) {
		val, _ := strconv.Atoi(v)
		if val > maxNum {
			maxNum = val
		}
	}

	return maxNum
}

func CheckNumberPrime(s string) bool {
	n, _ := strconv.Atoi(s)
	if n <= 1 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func CheckNumberPrimeLargest(s string) int {
	maxNum := RandomNumbersListMax(s)
	maxNumPrime := 0

	for i := maxNum; i >= 2; i-- {
		if CheckNumberPrime(strconv.Itoa(i)) {
			maxNumPrime = i
			break
		}
	}

	return maxNumPrime
}
