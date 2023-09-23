package Challenge12_ArmstrongNumbers

type MyInt int

func (n *MyInt) IsArmstrong() bool {
	number := *n
	var sum MyInt

	for *n > 0 {
		x := *n % 10
		sum += x * x * x
		*n /= 10
	}

	*n = number
	return sum == number
}

func Armstrong(inptNum int) bool {
	x := MyInt(inptNum)

	return x.IsArmstrong()
}
