package defer_debug

import (
	"errors"
	"fmt"
)

func Guessing(guess int) (string, error) {
	if guess < 1 || guess > 100 {
		return "", errors.New("not in 1-100")
	}

	return "Found!", nil
}

func Error3() {
	fmt.Println(Guessing(50))
	fmt.Println(Guessing(101))
}
