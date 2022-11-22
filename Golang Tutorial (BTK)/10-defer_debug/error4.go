package defer_debug

import "fmt"

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func Error4(predict int) (string, error) {
	x := borderException{predict, "Invalid prediction"}

	if predict < 1 || predict > 100 {
		return "", &x
	}

	return "Good!", nil

	// b := &borderException{predict, "Invalid prediction"}
	// fmt.Println(b.Error())
}
