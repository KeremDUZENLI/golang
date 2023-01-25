package function

import (
	"fmt"
	"generics/variable"
)

func PrintCard[C variable.Interfac](card C) {
	fmt.Println("card name: ", card.NameT())
}
