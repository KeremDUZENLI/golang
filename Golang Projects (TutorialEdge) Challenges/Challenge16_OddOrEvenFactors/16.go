package Challenge16_OddOrEvenFactors

func CheckFactors(num int) []int {
	factorialNumbers := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factorialNumbers = append(factorialNumbers, i)
		}
	}

	return factorialNumbers
}

func OddEvenFactors(num int) string {
	resList := CheckFactors(num)
	resLen := len(resList)

	if resLen%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
