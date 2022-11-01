package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	pl("Slice Size: ", len(sl1))

	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}

	sArr := [5]int{1, 2, 3, 4, 5}

	sl2 := sArr[0:2]
	pl(sl2)

	sl3 := sArr[:3]
	pl("First 3: ", sl3)

	sl4 := sArr[2:]
	pl("Last 3: ", sl4)

	sl4 = append(sl4, 100)
	pl("sl4: ", sl4)
	pl("New sArr: ", sArr)
}
