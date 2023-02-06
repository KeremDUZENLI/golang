package Challenge23_LargestNPrime

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

func RandomNumbers(n int) string {
	// intListe := make([]int, 0, r)
	strListe := make([]string, 0, n)
	seen := make(map[int]bool)

	for i := 0; i < n; {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(n)

		if !seen[randomNumber] {
			seen[randomNumber] = true
			// intListe = append(intListe, randomNumber)
			strListe = append(strListe, strconv.Itoa(randomNumber))
			i++
		}
	}

	// return intListe
	return strings.Join(strListe, "")
}

func RandomNumbersList(n int) []string {
	maxPossible := Factorial(n)
	strListe := []string{}
	seen := make(map[string]bool)

	for i := 0; i < maxPossible; {
		randomNumber := RandomNumbers(n)

		if !seen[randomNumber] {
			seen[randomNumber] = true
			strListe = append(strListe, randomNumber)
			i++
		}
	}

	return strListe
}

func RandomNumbersListMax(n int) int {
	maxNum := 0

	for _, v := range RandomNumbersList(n) {
		val, _ := strconv.Atoi(v)
		if val > maxNum {
			maxNum = val
		}
	}

	return maxNum
}

func CheckNumberPrime(n int) bool {
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

func CheckNumberPrimeLargest(n int) int {
	maxNum := RandomNumbersListMax(n)
	maxNumPrime := 0

	for i := maxNum; i >= 2; i-- {
		if CheckNumberPrime(i) {
			maxNumPrime = i
			break
		}
	}

	return maxNumPrime
}
