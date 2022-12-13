package defer_debug

import (
	"fmt"
	"os"
)

func Error1() {
	file, err := os.Open("error1.txt")

	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Printf("File not found : %v\n", pathError.Path)
			return
		} else {
			fmt.Println("File not found!")
			return
		}
	} else {
		fmt.Println(file.Name())
	}
}
