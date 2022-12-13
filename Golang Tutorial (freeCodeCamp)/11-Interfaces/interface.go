package Interfaces

import (
	"fmt"
)

func Interface() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer2 interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
